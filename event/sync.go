package event

import (
	"fmt"
	"github.com/chaechaep/simulator/config"
	"github.com/chaechaep/simulator/types"
)

func GetSync(accessToken string, since string) (ret types.SyncResp, err error) {
	url := config.Cfg.BaseUrl + "/sync?full_state=false&timeout=30000"
	if since != "" {
		url += "&since=" + since
	}
	err = Process("GET", url, nil, &ret, accessToken)
	if err != nil {
		return ret, fmt.Errorf("get sync failed : %s", err)
	}
	fmt.Println(ret)
	return ret, nil
}
