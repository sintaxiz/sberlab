package tfApp

import (
	"terraformManager/api"
	"terraformManager/server"

	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("app")

func Run() {
	logger.SetLogLevel(loggo.INFO)

	tfAPI := api.NewTfAPI()

	tfServer := server.NewServer(tfAPI)

	if err := tfServer.RegisterAndServe(); err != nil {
		panic(err)
	}
}
