package event

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/CHAEUNPARK/simulator/config"
	"github.com/CHAEUNPARK/simulator/types"
	"io"
	"net/http"
)

var client = &http.Client{}
var errorRet = types.Error{}

func Process(method string, url string, reqValue []byte, respValue interface{}, auth string) (ret map[string]interface{}, err error) {
	var body io.Reader = nil
	if reqValue != nil {
		body = bytes.NewBuffer(reqValue)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return ret, fmt.Errorf("request create error : %s", err.Error())
	}

	if reqValue != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	if auth != "" {
		req.Header.Add("Authorization", config.Cfg.AccessTokenPrefix+auth)
	}

	resp, err := client.Do(req)
	if err != nil {
		return ret, fmt.Errorf("send request error : %s", err.Error())
	}

	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&respValue); err != nil {
			return ret, fmt.Errorf("response decode error : %s", err.Error())
		}
	} else {
		if err := json.NewDecoder(resp.Body).Decode(&errorRet); err != nil {
			return ret, fmt.Errorf("response decode error : %s", err.Error())
		} else {
			return ret, fmt.Errorf("status code : %d, error code : %s, error msg : %s", resp.StatusCode, errorRet.Errcode, errorRet.Error)
		}
	}

	ret = respValue.(map[string]interface{})
	return ret, nil
}
