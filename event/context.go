package event

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/chaechaep/simulator/config"
	"github.com/chaechaep/simulator/types"
	"io"
	"net/http"
)

func Process(method string, url string, reqValue []byte, respValue interface{}, auth string) error {
	client := http.Client{}
	errorRet := types.Error{}
	var body io.Reader = nil
	if reqValue != nil {
		body = bytes.NewBuffer(reqValue)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return fmt.Errorf("request create error : %s", err.Error())
	}

	if reqValue != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	if auth != "" {
		req.Header.Add("Authorization", config.Cfg.AccessTokenPrefix+auth)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("send request error : %s", err.Error())
	}

	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(respValue); err != nil {
			return fmt.Errorf("response decode error : %s", err.Error())
		}
	} else {
		if err := json.NewDecoder(resp.Body).Decode(&errorRet); err != nil {
			return fmt.Errorf("response decode error : %s", err.Error())
		} else {
			return fmt.Errorf("status code : %d, error code : %s, error msg : %s", resp.StatusCode, errorRet.Errcode, errorRet.Error)
		}
	}
	fmt.Println(respValue)

	return nil
}
