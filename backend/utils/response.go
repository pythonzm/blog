package utils

import (
	"fmt"
)

func GenResponse(code int, data interface{}, err error) map[string]interface{} {
	res := make(map[string]interface{})
	res["code"] = code
	res["data"] = data
	if err != nil {
		res["msg"] = fmt.Sprintf("%v: %v", ErrorText(code), err)
	} else {
		res["msg"] = ErrorText(code)
	}

	return res
}
