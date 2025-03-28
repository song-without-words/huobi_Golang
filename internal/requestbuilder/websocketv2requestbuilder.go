package requestbuilder

import (
	"time"

	"github.com/song-without-words/huobi_Golang/internal/model"
	model2 "github.com/song-without-words/huobi_Golang/pkg/model"
)

type WebSocketV2RequestBuilder struct {
	akKey   string
	akValue string
	smKey   string
	smValue string
	svKey   string
	svValue string
	tKey    string
	tValue  string

	host string
	path string
	sign string

	signer SignerInterface
}

func (p *WebSocketV2RequestBuilder) Init(accessKey string, secretKey string, host string, path string, sign string) *WebSocketV2RequestBuilder {
	p.akKey = "accessKey"
	p.akValue = accessKey
	p.smKey = "signatureMethod"
	p.smValue = "HmacSHA256"
	p.svKey = "signatureVersion"
	p.svValue = "2.1"
	p.tKey = "timestamp"
	p.sign = sign
	p.host = host
	p.path = path

	if sign == "256" {
		p.signer = new(Signer).Init(secretKey) // Signer 实现了接口
	} else {
		// 使用 Ed25519 签名
		edSigner := new(Ed25519Signer)
		var err error
		edSigner, err = edSigner.Init(secretKey)
		if err != nil {
			// 处理错误
			return nil // 假设这是在一个返回 error 的函数中
		}
		p.signer = edSigner // Ed25519Signer 也实现了接口
	}

	return p
}

func (p *WebSocketV2RequestBuilder) Build() (string, error) {
	time := time.Now().UTC()
	return p.build(time)
}

func (p *WebSocketV2RequestBuilder) build(utcDate time.Time) (string, error) {
	time := utcDate.Format("2006-01-02T15:04:05")

	req := new(model2.GetRequest).Init()
	req.AddParam(p.akKey, p.akValue)
	req.AddParam(p.smKey, p.smValue)
	req.AddParam(p.svKey, p.svValue)
	req.AddParam(p.tKey, time)

	signature, err := p.signer.Sign("GET", p.host, p.path, req.BuildParams())

	auth := new(model.WebSocketV2AuthenticationRequest).Init()
	auth.Params.AccessKey = p.akValue
	auth.Params.Timestamp = time
	auth.Params.Signature = signature

	result, err := model2.ToJson(auth)
	return result, err
}
