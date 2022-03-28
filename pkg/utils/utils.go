package utils

import (
	"encoding/json"
	"fmt"
)

func ToJsonString(s interface{}, leading string) string {
	sJSON, err := json.MarshalIndent(s, "", leading)
	if err != nil {
		panic(fmt.Sprintf("%s", err))
	}
	return string(sJSON)
}
