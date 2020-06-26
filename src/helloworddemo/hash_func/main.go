package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func hashStr(str string) string {
	byteStr := sha256.Sum256([]byte(str))
	return hex.EncodeToString(byteStr[:])
}

func main() {

	fmt.Println(hashStr("test1"))
}
