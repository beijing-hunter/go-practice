package token

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)


func CheckToken(token string) (uid int, username string, err error) {
	fmt.Println("token:", token)
	arr := strings.Split(token, ".")
	if len(arr) != 3 {
		err = errors.New("token error1")//创建err为"oken error1"
		return
	}
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])

	_, err = base64.StdEncoding.DecodeString(arr[0])
	if err != nil {
		err = errors.New("token error2")
		return
	}
	pay, err := base64.StdEncoding.DecodeString(arr[1])
	if err != nil {
		err = errors.New("token error3")
		return
	}
	sign, err := base64.StdEncoding.DecodeString(arr[2])
	if err != nil {
		err = errors.New("token error4")
		return
	}

	str1 := arr[0] + "." + arr[1]

	key := []byte("szs")
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(str1))
	s := mac.Sum(nil)
	fmt.Println(sign)
	fmt.Println(s)
	if res := bytes.Compare(sign, s); res != 0 {//Compare 比较俩者是否相等
		fmt.Println("test")
		err = errors.New("token error5")
		return
	}

	var payload Payload
	json.Unmarshal(pay,&payload)
	uid = payload.Uid
	username = payload.Username
	return
}
