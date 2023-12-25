package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/farmerx/gorsa"
)

func New() *RSA {
	return &RSA{}
}

// RSA 一般来说rsa算法都是使用公钥加密，私钥解密，或者私钥签名，公钥验签。
type RSA struct {
	publicKey  []byte //公钥
	privateKey []byte //私钥
}

func (r *RSA) SetPublicKey(data []byte) *RSA {
	r.publicKey = data
	return r
}
func (r *RSA) SetPrivateKey(data []byte) *RSA {
	r.privateKey = data
	return r
}

func (r *RSA) GetPublicKey() []byte {
	return r.publicKey
}
func (r *RSA) GetPrivateKey() []byte {
	return r.privateKey
}

// GenerateKey 生成私钥
// bits 1024
func (r *RSA) GenerateKey(bits int) error {
	privateKeyTmp, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKeyTmp)

	publicKeyTmp := &privateKeyTmp.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKeyTmp)
	if err != nil {
		return err
	}

	r.SetPrivateKey(
		pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: derStream,
		}),
	).SetPublicKey(
		pem.EncodeToMemory(&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: derPkix,
		}),
	)
	return nil
}

// SignWithSha256 私钥签名
func (r *RSA) SignWithSha256(data []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	block, _ := pem.Decode(r.privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
}

// VerifySignWithSha256 公钥验证
func (r *RSA) VerifySignWithSha256(data, signData []byte) error {
	block, _ := pem.Decode(r.publicKey)
	if block == nil {
		return errors.New("public key error")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	hashed := sha256.Sum256(data)
	return rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signData)
}

// 一般服务器保存私钥， 客户保留公钥
// 公钥加密 => 私钥解密  或 私钥加密 => 公钥解密

// PublicDecrypt 公钥解密
func (r *RSA) PublicDecrypt(ciphertext []byte) ([]byte, error) {
	if err := gorsa.RSA.SetPublicKey(string(r.publicKey)); err != nil {
		return nil, err
	}
	return gorsa.RSA.PubKeyDECRYPT(ciphertext)
}

// PublicEncrypt 公钥加密
func (r *RSA) PublicEncrypt(data []byte) ([]byte, error) {
	if err := gorsa.RSA.SetPublicKey(string(r.publicKey)); err != nil {
		return nil, err
	}
	return gorsa.RSA.PubKeyENCTYPT(data)
}

// PrivateDecrypt 私钥解密
func (r *RSA) PrivateDecrypt(ciphertext []byte) ([]byte, error) {
	if err := gorsa.RSA.SetPrivateKey(string(r.privateKey)); err != nil {
		return nil, err
	}
	return gorsa.RSA.PriKeyDECRYPT(ciphertext)
}

// PrivateEncrypt 私钥加密
func (r *RSA) PrivateEncrypt(data []byte) ([]byte, error) {
	if err := gorsa.RSA.SetPrivateKey(string(r.privateKey)); err != nil {
		return nil, err
	}
	return gorsa.RSA.PriKeyENCTYPT(data)
}
