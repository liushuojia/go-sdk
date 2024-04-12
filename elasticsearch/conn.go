package elasticsearchConn

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

func New() *Conn {
	return &Conn{}
}

/*
可变更curl模式 索引
curl -X PUT 'localhost:9200/weather'
curl -X DELETE 'localhost:9200/weather'

查询一般安装一个中文分词， 凡是需要搜索的中文字段，都要单独设置一下
curl -X PUT 'localhost:9200/accounts' -d '
{
  "mappings": {
    "person": {
      "properties": {
        "user": {
          "type": "text",
          "analyzer": "ik_max_word",
          "search_analyzer": "ik_max_word"
        },
        "title": {
          "type": "text",
          "analyzer": "ik_max_word",
          "search_analyzer": "ik_max_word"
        }
      }
    }
  }
}'

curl -X PUT 'localhost:9200/accounts' -d '

curl -X GET localhost:9200/{index}/{type}/{id}

获取数据
curl -X GET localhost:9200/test001/_doc/0079

修改数据（创建 或 修改 全量更新）
curl -X PUT localhost:9200/test001/_doc/0079

删除数据
curl -X DELETE localhost:9200/test001/_doc/0079

局部更新
curl -X POST localhost:9200/test001/_doc/0079/_update
{
   "doc": {
      "my": 1
   }
}

创建数据 id随机生成
curl -X POST localhost:9200/test001/_doc

查询数据
curl -X GET localhost:9200/test001/_doc/_search
{
  "query" : { "match" : { "desc" : "软件" }}
  "from":0,
  "size":10
}'
*/

type Conn struct {
	Addresses []string
	Username  string
	Password  string
	client    *elasticsearch.Client
}

func (es *Conn) SetAddresses(addressesList ...string) *Conn {
	es.Addresses = addressesList
	return es
}
func (es *Conn) SetUsername(username string) *Conn {
	es.Username = username
	return es
}
func (es *Conn) SetPassword(password string) *Conn {
	es.Password = password
	return es
}

func (es *Conn) Version() string {
	return elasticsearch.Version
}
func (es *Conn) Connect() error {
	log.Println("[elasticsearch]", "connect elasticsearch", elasticsearch.Version)

	cfg := elasticsearch.Config{
		Addresses: es.Addresses,
		Username:  es.Username,
		Password:  es.Password,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			TLSClientConfig: &tls.Config{
				MinVersion:         tls.VersionTLS12,
				InsecureSkipVerify: true,
			},
		},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return err
	}
	es.client = client
	return nil
}
func (es *Conn) Info() (*Info, error) {
	res, err := es.CheckResponse(es.client.Info())
	if err != nil {
		return nil, err
	}

	var r *Info
	if err := es.Source(res, r); err != nil {
		return nil, err
	}
	return r, nil
}

/*
EsApiIndexRequest

	document := struct {
		Name string `json:"name"`
		Type int    `json:"type"`
	}{

		Name: "go-elasticsearch",
		Type: 2,
	}
	data, _ := json.Marshal(document)
	req := esapi.IndexRequest{
		Index:      "test",
		DocumentID: "10099", //为空时 创建 ID自动生成， 不为空时，创建或更新
		Body:       bytes.NewReader(data),
		Refresh:    "true",
	}

// =======================

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "test",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	req := esapi.IndexRequest{
		Index:      "test",
		DocumentID: "10099",
		Body:       bytes.NewReader(buf.Bytes()),
		Refresh:    "true",
	}
*/
func (es *Conn) EsApiIndexRequest(indexRequest esapi.IndexRequest) (*esapi.Response, error) {
	//es.Connect()
	return indexRequest.Do(context.Background(), es.client)
}
func (es *Conn) CheckResponse(res *esapi.Response, err error) (*Response, error) {
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.IsError() {
		return nil, errors.New(res.String())
	}

	var r *Response
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}
	return r, nil
}
func (es *Conn) Source(r *Response, data interface{}) error {
	if r == nil {
		return errors.New("data is empty")
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(r.Source); err != nil {
		return err
	}
	if err := json.NewDecoder(&buf).Decode(&data); err != nil {
		return err
	}
	return nil
}

func (es *Conn) Save(index, id string, data interface{}) (string, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		return "", err
	}
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: id,
		Body:       bytes.NewReader(buf.Bytes()),
		Refresh:    "true",
	}

	res, err := es.CheckResponse(es.EsApiIndexRequest(req))
	if err != nil {
		return "", err
	}

	return res.Id, nil
}
func (es *Conn) Get(index, id string, data interface{}) error {
	res, err := es.CheckResponse(es.client.Get(index, id))
	if err != nil {
		return err
	}

	if err := es.Source(res, &data); err != nil {
		return err
	}
	return nil
}
func (es *Conn) Delete(index, id string) error {
	_, err := es.CheckResponse(es.client.Delete(index, id))
	if err != nil {
		return err
	}
	return nil
}

// Update es.Update("test001", "0079", strings.NewReader(`{"doc": { "language": "Go" }}`)),
func (es *Conn) Update(index, id string, data io.Reader) error {
	res, err := es.client.Update(index, id, data)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

// Query
// query 查询内容
// size 数量 from 偏移
// "sort": [
// { "price": "asc" },
// { "date": "desc" }
// ]
// {"query": { "match": { "name" : "go" } }}
func (es *Conn) Query(index string, query io.Reader, data interface{}) (int64, error) {
	res, err := es.client.Search(
		es.client.Search.WithIndex(index),
		es.client.Search.WithBody(query),
		es.client.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	var r *QueryResponse
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return 0, err
	}

	var list []interface{}
	for _, item := range r.Hits.Hits {
		list = append(list, item.Source)
	}

	total := r.Hits.Total.Value
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(list); err != nil {
		return total, nil
	}
	if err := json.NewDecoder(&buf).Decode(&data); err != nil {
		return total, nil
	}

	return total, nil
}

func (es *Conn) IndexCreate(index string) error {
	_, err := es.CheckResponse(es.client.Indices.Create(index))
	return err
}
func (es *Conn) IndexDelete(index string) error {
	_, err := es.CheckResponse(es.client.Indices.Delete([]string{index}))
	return err
}
func (es *Conn) IndexExists(index string) error {
	_, err := es.CheckResponse(es.client.Indices.Exists([]string{index}))
	return err
}

//response, err := client.Indices.Exists([]string{"my_index2"})

//fmt.Println(client.Indices.Delete([]string{"test"}))

//fmt.Println(client.Indices.Create("my_index1"))
