package providers

import (
	"encoding/json"
	"io"
)

func UnmarshalBody(Body io.ReadCloser) map[string]interface{} {
	var result map[string]interface{}

	respBody, err := io.ReadAll(Body)
	if err != nil {
		return result
	}

	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return result
	}

	return result
}
