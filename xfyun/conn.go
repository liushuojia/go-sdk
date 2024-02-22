package xfYun

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func New() *Conn {
	return &Conn{
		//imageUrl: "https://spark-api.cn-huabei-1.xf-yun.com/v2.1/tti",
		//chatUrl:  "wss://spark-api.xf-yun.com/v3.5/chat",
		imageUrl: "https://spark-api.cn-huabei-1.xf-yun.com/v2.1/tti",
		chatUrl:  "wss://spark-api.xf-yun.com/v3.5/chat",
	}
}

type Callback func(message string, flag bool)

type Conn struct {
	appId     string
	apiKey    string
	apiSecret string
	uid       string

	imageUrl string
	chatUrl  string
}

func (c *Conn) SetUid(uid string) *Conn {
	c.uid = uid
	return c
}
func (c *Conn) SetAppId(appId string) *Conn {
	c.appId = appId
	return c
}
func (c *Conn) SetApiKey(apiKey string) *Conn {
	c.apiKey = apiKey
	return c
}
func (c *Conn) SetApiSecret(apiSecret string) *Conn {
	c.apiSecret = apiSecret
	return c
}
func (c *Conn) SetImageUrl(imageUrl string) *Conn {
	c.imageUrl = imageUrl
	return c
}
func (c *Conn) SetChatUrl(chatUrl string) *Conn {
	c.chatUrl = chatUrl
	return c
}

func (c *Conn) BuildImage(content string, width int64, height int64) (base64Image string, err error) {
	authAddr := c.assembleAuthUrl("POST", c.imageUrl, c.apiKey, c.apiSecret)
	reqMsg := []message{
		{Content: content, Role: "user"}, // 只能有1个  不能传对话历史
	}

	req := map[string]interface{}{
		"header": map[string]interface{}{
			"app_id": c.appId,
			"uid":    c.uid,
		},
		"parameter": map[string]interface{}{
			"chat": map[string]interface{}{
				//"domain": "general",
				"width":  width,
				"height": height,
			},
		},
		"payload": map[string]interface{}{
			"message": map[string]interface{}{
				"text": reqMsg,
			},
		},
	}

	reqData, err := json.Marshal(req)
	if err != nil {
		return
	}

	// 发起http请求
	respData, err := c.HttpTool("POST", authAddr, reqData, 20000)
	if err != nil {
		return
	}

	var result Response
	err = json.Unmarshal(respData, &result)
	if err != nil {
		return
	}
	if len(result.Payload.Choices.Text) <= 0 {
		err = errors.New(result.Header.Message)
		return
	}
	base64Image = result.Payload.Choices.Text[0].Content

	// 取出图片 解析并保存
	//base64Image := result.Payload.Choices.Text[0].Content
	//image, err := base64.StdEncoding.DecodeString(base64Image)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//// 保存图片
	//ioutil.WriteFile("ttiDemo.jpg", image, 0777)
	return
}
func (c *Conn) AskQuestion(question string, cb Callback) (err error) {
	d := websocket.Dialer{
		HandshakeTimeout: 5 * time.Second,
	}

	//握手并建立websocket 连接
	conn, resp, err := d.Dial(c.assembleAuthUrl("GET", c.chatUrl, c.apiKey, c.apiSecret), nil)
	if err != nil {
		return
	}

	if resp.StatusCode != 101 {
		var b []byte
		b, err = io.ReadAll(resp.Body)
		if err != nil {
			return
		}
		err = errors.New(fmt.Sprintf("code=%d,body=%s", resp.StatusCode, string(b)))
		return
	}

	go func() {
		messages := []message{
			{Role: "user", Content: question},
		}

		data := map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
			"header": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
				"app_id": c.appId, // 根据实际情况修改返回的数据结构和字段名
				"uid":    c.uid,
			},
			"parameter": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
				"chat": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
					"domain":      "generalv3.5", // 根据实际情况修改返回的数据结构和字段名
					"temperature": float64(0.8),  // 根据实际情况修改返回的数据结构和字段名
					"top_k":       int64(6),      // 根据实际情况修改返回的数据结构和字段名
					"max_tokens":  int64(2048),   // 根据实际情况修改返回的数据结构和字段名
					"auditing":    "default",     // 根据实际情况修改返回的数据结构和字段名
				},
			},
			"payload": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
				"message": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
					"text": messages, // 根据实际情况修改返回的数据结构和字段名
				},
			},
		}
		conn.WriteJSON(data)
	}()
	go func() {
		defer conn.Close()
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			var data map[string]interface{}
			if err := json.Unmarshal(msg, &data); err != nil {
				return
			}

			//解析数据
			payload, ok := data["payload"].(map[string]interface{})
			if !ok {
				break
			}

			choices, ok := payload["choices"].(map[string]interface{})
			if !ok {
				break
			}

			header, ok := data["header"].(map[string]interface{})
			if !ok {
				break
			}

			code, ok := header["code"].(float64)
			if !ok {
				break
			}

			if code != 0 {
				break
			}

			status, ok := choices["status"].(float64)
			if !ok {
				break
			}

			text, ok := choices["text"].([]interface{})
			if !ok {
				break
			}

			content, ok := text[0].(map[string]interface{})["content"].(string)
			if !ok {
				break
			}

			if status != 2 {
				//answer += content
				cb(content, false)
				continue
			} else {
				fmt.Println("收到最终结果")
				//answer += content
				cb(content, true)
				usage := payload["usage"].(map[string]interface{})
				temp := usage["text"].(map[string]interface{})
				totalTokens := temp["total_tokens"].(float64)
				fmt.Println("total_tokens:", totalTokens)
				conn.Close()
				break
			}
		}
	}()
	return
}
func (c *Conn) AskQuestionAndAnswer(question string) (result string, err error) {
	d := websocket.Dialer{
		HandshakeTimeout: 5 * time.Second,
	}

	//握手并建立websocket 连接
	conn, resp, err := d.Dial(c.assembleAuthUrl("GET", c.chatUrl, c.apiKey, c.apiSecret), nil)
	if err != nil {
		return
	}
	defer conn.Close()

	if resp.StatusCode != 101 {
		var b []byte
		b, err = io.ReadAll(resp.Body)
		if err != nil {
			return
		}
		err = errors.New(fmt.Sprintf("code=%d,body=%s", resp.StatusCode, string(b)))
		return
	}

	go func() {
		messages := []message{
			{Role: "user", Content: question},
		}

		data := map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
			"header": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
				"app_id": c.appId, // 根据实际情况修改返回的数据结构和字段名
				"uid":    c.uid,
			},
			"parameter": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
				"chat": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
					"domain":      "generalv3.5", // 根据实际情况修改返回的数据结构和字段名
					"temperature": float64(0.8),  // 根据实际情况修改返回的数据结构和字段名
					"top_k":       int64(6),      // 根据实际情况修改返回的数据结构和字段名
					"max_tokens":  int64(2048),   // 根据实际情况修改返回的数据结构和字段名
					"auditing":    "default",     // 根据实际情况修改返回的数据结构和字段名
				},
			},
			"payload": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
				"message": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
					"text": messages, // 根据实际情况修改返回的数据结构和字段名
				},
			},
		}
		conn.WriteJSON(data)
	}()

	result = ""
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return result, err
		}
		var data map[string]interface{}
		if err := json.Unmarshal(msg, &data); err != nil {
			return result, err
		}

		//解析数据
		payload, ok := data["payload"].(map[string]interface{})
		if !ok {
			return result, errors.New("payload 返回数据格式错误")
		}

		choices, ok := payload["choices"].(map[string]interface{})
		if !ok {
			return result, errors.New("choices 返回数据格式错误")
		}

		header, ok := data["header"].(map[string]interface{})
		if !ok {
			return result, errors.New("header 返回数据格式错误")

		}

		code, ok := header["code"].(float64)
		if !ok {
			return result, errors.New("header code 返回数据格式错误")

		}

		if code != 0 {
			return result, errors.New("code值错误")

		}

		status, ok := choices["status"].(float64)
		if !ok {
			return result, errors.New("choices status 返回数据格式错误")

		}

		text, ok := choices["text"].([]interface{})
		if !ok {
			return result, errors.New("choices text 返回数据格式错误")
		}

		content, ok := text[0].(map[string]interface{})["content"].(string)
		if !ok {
			return result, errors.New("choices text content 返回数据格式错误")
		}

		if status != 2 {
			result += content
			continue
		} else {
			fmt.Println("收到最终结果")
			result += content
			usage := payload["usage"].(map[string]interface{})
			temp := usage["text"].(map[string]interface{})
			totalTokens := temp["total_tokens"].(float64)
			fmt.Println("total_tokens:", totalTokens)
			return result, nil
		}
	}
}

// AssembleAuthUrl 鉴权
// @hosturl :  like  wss://ws-api.xfyun.cn/v2/iat
// @apikey : apiKey
// @apiSecret : apiSecret
// 注意：如果是ws请求，method应该是"GET";  如果是POST请求,method应该是"POST"
func (c *Conn) assembleAuthUrl(method, addr, apiKey, apiSecret string) string {
	if apiKey == "" && apiSecret == "" { // 不鉴权
		return addr
	}
	ul, err := url.Parse(addr) // 地址不对  也不鉴权
	if err != nil {
		return addr
	}
	//签名时间
	date := time.Now().UTC().Format(time.RFC1123)
	//参与签名的字段 host ,date, request-line
	signString := []string{"host: " + ul.Host, "date: " + date, method + " " + ul.Path + " HTTP/1.1"}
	//拼接签名字符串
	sgin := strings.Join(signString, "\n")
	//签名结果
	sha := c.hmacWithShaToBase64("hmac-sha256", sgin, apiSecret)
	//构建请求参数 此时不需要urlencoding
	authUrl := fmt.Sprintf("api_key=\"%s\", algorithm=\"%s\", headers=\"%s\", signature=\"%s\"", apiKey,
		"hmac-sha256", "host date request-line", sha)
	//将请求参数使用base64编码
	authorization := base64.StdEncoding.EncodeToString([]byte(authUrl))

	v := url.Values{}
	v.Add("host", ul.Host)
	v.Add("date", date)
	v.Add("authorization", authorization)
	//将编码后的字符串url encode后添加到url后面
	authAddr := addr + "?" + v.Encode()
	return authAddr
}
func (c *Conn) hmacWithShaToBase64(algorithm, data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	encodeData := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(encodeData)
}

func (c *Conn) HttpTool(method, authAddr string, data []byte, timeout int) ([]byte, error) {
	// 发起请求
	req, err := http.NewRequest(method, authAddr, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json;charset=UTF-8")
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Millisecond,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应
	respData, err := io.ReadAll(resp.Body)
	return respData, err
}

// 响应给用户的数据格式
type Response struct {
	Header struct {
		Code    int    `json:"code"`    // 0
		Message string `json:"message"` // Success
		Sid     string `json:"sid"`
		Status  int    `json:"status"` // 1表示会话的中间结果 2表示整个会话最后结果
	} `json:"header"`
	Payload struct {
		Choices *ChoiceText `json:"choices,omitempty"`
	}
}

type ChoiceText struct {
	Status int       `json:"status"` // 0表示第一个结果 1表示中间结果 2表示最后一个结果
	Seq    int       `json:"seq"`    // 结果编号 0，1，2，3...
	Text   []message `json:"text"`   // 结果文本
}

type message struct {
	Content string `json:"content"` // 用户的对话内容 或者 是base64的图片
	Role    string `json:"role"`    // system表示存放基本的prompt,人设信息位于顶层 user表示用户(问题或者图片)  assistant表示大模型
	// 多模数据的描述
	ContentType string `json:"content_type,omitempty"` // image表示图片  text是文本
	Index       int    `json:"index"`                  // 大模型的结果序号，在多候选中使用。  这个传给大模型引擎是多余的，但是不影响。  但是这个字段必须要传给用户，因此不能加omitempty属性
}
