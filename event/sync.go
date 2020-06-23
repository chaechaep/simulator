package event

import (
	"fmt"
	"github.com/chaechaep/simulator/config"
	"github.com/chaechaep/simulator/types"
	"strconv"
)

func GetSync(accessToken string, since string, syncDuration int) (ret types.SyncResp, err error) {
	url := config.Cfg.Simulator.BaseUrl + "/sync?full_state=false&timeout=" + strconv.Itoa(syncDuration*1000)
	if since != "" {
		url += "&since=" + since
	}
	err = Process("GET", url, nil, &ret, accessToken)
	if err != nil {
		return ret, fmt.Errorf("get sync failed : %s", err)
	}
	return ret, nil
}
