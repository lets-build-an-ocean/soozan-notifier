package providers

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// TODO: Return a boolean value and the response of server

func Main(mobile string, scenario string, params []string) int {
	url := "https://api.sms.ir/v1/send/verify"
	var templateID int
	var parameters []map[string]string

	switch scenario {
	case "OTP":
		templateID = 156591
		parameters = []map[string]string{{"name": "OTP", "value": params[0]}}
	case "ORDER_READY":
		templateID = 131187
		parameters = []map[string]string{{"name": "ORDER_NUMBER", "value": params[0]}}
	}

	body := map[string]interface{}{
		"mobile":     mobile,
		"templateID": templateID,
		"parameters": parameters,
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return 0
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return 0
	}
	// TODO: Move token to env vars
	req.Header.Set("x-api-key", "API_KEY")
	req.Header.Set("Accept", "text/plain")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0
	}
	defer resp.Body.Close()

	return resp.StatusCode
}

func Fallback(mobile string, scenario string, params []string) int {
	// TODO: Move token to env vars
	url := "https://console.melipayamak.com/api/send/shared/API_KEY"
	var bodyId int

	switch scenario {
	case "OTP":
		bodyId = 432895
	case "ORDER_READY":
		bodyId = 000000
	}

	body := map[string]interface{}{
		"to":     mobile,
		"bodyId": bodyId,
		"args":   params,
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return 0
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return 0
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0
	}
	defer resp.Body.Close()

	result := UnmarshalBody(resp.Body)
	if result["recId"] == nil {
		return 0
	}
	return 200
}
