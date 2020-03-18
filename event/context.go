package event

import (
	"bytes"
	"fmt"
	"github.com/CHAEUNPARK/simulator/config"
	"github.com/CHAEUNPARK/simulator/types"
	"io"
	"net/http"
)

var client = &http.Client{}

func Process(user *types.User, method string, url string, reqValue []byte, respValue interface{}, auth bool) (ret types.JSONResponse, err error) {
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
	if auth {
		req.Header.Add("Authorization", config.Cfg.AccessTokenPrefix+user.AccessToken)
	}
	return ret, nil
}
