package ebarimt

import (
	"fmt"
	"github.com/lambda-platform/ebarimt/posapi"
	"github.com/lambda-platform/ebarimtgrpc/config"
)

func PosAPI(pos_id string) *posapi.PosAPI {
	var API *posapi.PosAPI

	api, err := posapi.NewPosAPI(config.SoFilesPath + pos_id + ".so")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		API = api
		defer API.Close()
	}

	return API
}
