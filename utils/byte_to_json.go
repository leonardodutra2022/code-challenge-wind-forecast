package utils

import "encoding/json"

func ByteToJson(bytes []byte, obj interface{}) error {
	return json.Unmarshal(bytes, &obj)
}
