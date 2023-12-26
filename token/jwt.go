package tokenConn

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/liushuojia/go-sdk/utils"
	"strings"
	"time"
)

const (
	//未设置密钥有效期， 默认设置
	JwtKey    = "1237894560abcqweyuirtphjlkasdfgzxcv,bnm.!@#$%^&*()_"
	JwtExpire = "5m"
	JwtIssuer = "liushuojia"
)

func New() *JWT {
	return &JWT{}
}

// 加密主体, 根据实际情况修改
type Claims struct {
	ID         int64  `json:"i"`            // wjt加密密钥 如果停用已颁发的wjt修改该值即可
	Key        string `json:"k"`            // wjt加密密钥 如果停用已颁发的wjt修改该值即可
	Data       any    `json:"d,omitempty"`  // 存储数据 map or struct
	UUID       string `json:"u,omitempty"`  // uuid
	ParentUUID string `json:"pu,omitempty"` // 父级uuid

	jwt.StandardClaims
}

// 处理结构
type JWT struct {
	Claims    Claims
	expire    string
	jwtSecret []byte
}

func (j *JWT) SetID(id int64) *JWT {
	j.Claims.ID = id
	return j
}
func (j *JWT) SetKey(key string) *JWT {
	j.Claims.Key = key
	return j
}
func (j *JWT) SetData(data any) *JWT {
	j.Claims.Data = data
	return j
}
func (j *JWT) SetExpire(Expire string) *JWT {
	j.expire = Expire
	return j
}
func (j *JWT) SetSecret(secret []byte) *JWT {
	j.jwtSecret = secret
	return j
}
func (j *JWT) SetIssuer(issuer string) *JWT {
	j.Claims.Issuer = issuer
	return j
}
func (j *JWT) SetUUID(uuid string) *JWT {
	j.Claims.UUID = uuid
	return j
}
func (j *JWT) SetParentUUID(parentUUID string) *JWT {
	j.Claims.ParentUUID = parentUUID
	return j
}
func (j *JWT) SetClaims(claims Claims) *JWT {
	j.Claims = claims
	return j
}

func (j *JWT) GetExpire() (ExpiresAt time.Time) {
	return time.Unix(j.Claims.StandardClaims.ExpiresAt, 0)
}
func (j *JWT) GetSecret() []byte {
	return j.jwtSecret
}
func (j *JWT) GetIssuer() string {
	return j.Claims.Issuer
}
func (j *JWT) GetUUID() string {
	return j.Claims.UUID
}
func (j *JWT) GetParentUUID() string {
	return j.Claims.ParentUUID
}
func (j *JWT) GetData() any {
	return j.Claims.Data
}
func (j *JWT) LoadData(o any) error {
	return utils.Copier(o, j.GetData())
}

// Generate 生成token
func (j *JWT) Generate() (token string, err error) {

	if len(j.jwtSecret) == 0 {
		j.SetSecret([]byte(JwtKey))
	}
	if j.expire == "" {
		j.SetExpire(JwtExpire)
	}
	if j.Claims.Issuer == "" {
		j.Claims.Issuer = JwtIssuer
	}

	m, _ := time.ParseDuration(j.expire)
	j.Claims.StandardClaims.ExpiresAt = time.Now().Add(m).Unix()
	j.Claims.IssuedAt = time.Now().Unix()

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, j.Claims)
	token, err = tokenClaims.SignedString(j.jwtSecret)
	return
}

// Parse 验证token
func (j *JWT) Parse(token string) (*Claims, error) {
	if token == "" {
		return nil, errors.New("token为空")
	}
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, errors.New("数据为空")
}

// Payload 获取token主体信息, 不验证加token是否有效
func (j *JWT) Payload(token string) (claims *Claims, err error) {
	if token == "" {
		err = errors.New("token为空")
		return
	}
	tmpArray := strings.Split(token, ".")
	if len(tmpArray) != 3 {
		err = errors.New("token参数传递错误")
		return
	}
	str, err := base64.RawURLEncoding.DecodeString(tmpArray[1])
	if err != nil {
		return
	}
	err = json.Unmarshal(str, &claims)
	return
}
