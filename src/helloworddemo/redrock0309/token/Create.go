package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

//头
type Header struct {
	alg string `json:"alg"`
	typ string `json:"typ"`
}

func NewHeader() Header {
	return Header{
		alg: "HS256",
		typ: "JWT",
	}
}

type Payload struct {// token里面添加用户信息，验证token后可能会用到用户信息
	Iss      string `json:"iss"`
	Exp      string `json:"exp"`
	IssueAt      string `json:"iat"`
	Username string `json:"username"`
	Uid      int
}

func Create(username string, id int) string {
	header := NewHeader()
	payload := Payload{
		Iss:      "szs",                                                           //
		Exp:      strconv.FormatInt(time.Now().Add(10*time.Hour).Unix(), 10),//持续时间
		IssueAt:  strconv.FormatInt(time.Now().Unix(), 10),                  //签发时间
		Username: username,                                                        //用户名
		Uid:      id,                                                              //id
	}
	h, _ := json.Marshal(header)  //json初始化
	p, _ := json.Marshal(payload)
	headerBase64 := base64.StdEncoding.EncodeToString(h)
	payloadBase64 := base64.StdEncoding.EncodeToString(p)
	str1 := strings.Join([]string{headerBase64, payloadBase64}, ".")

	key := "szs"
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(str1))
	s := mac.Sum(nil)
	signature := base64.StdEncoding.EncodeToString(s)
	token := str1 + "." + signature
	return token
}
