package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum([]byte(nil)))
}

func WriteErrorLog(s string) {
	file, _ := os.OpenFile("logs/error.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0664)
	_, e := file.WriteString(s)
	if e != nil {
		fmt.Println(e)
	}
}

func IfContainStr(s string, substrs []string) bool {
	r := false
	for _, substr := range substrs {
		if strings.Contains(s, substr) {
			r = true
			break
		}
	}
	return r
}
