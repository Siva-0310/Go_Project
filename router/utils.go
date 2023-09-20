package router

import (
	"encoding/json"
)

func ParseJson(data map[string]interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return jsonData, err
}
