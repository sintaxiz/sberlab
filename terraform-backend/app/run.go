package app

import (
	"os"
	apisample "terraformAPI/apiSample"

	v1 "terraformAPI/apiSample/v1"

	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("app")

func Run() {
	logger.SetLogLevel(loggo.INFO)

	sampleAPI := v1.NewSampleAPI()

	tfHostname, isHostSpecified := os.LookupEnv("TF_MANAGER_HOST")
	if !isHostSpecified {
		tfHostname = "http://localhost:9900"
		err := os.Setenv("TF_MANAGER_HOST", tfHostname)
		if err != nil {
			logger.Errorf("Can not set default host for tf manager...")
		}
	}
	logger.Infof("tf manager host: " + tfHostname)

	sampleServer := apisample.NewServer(sampleAPI)
	if err := sampleServer.RegisterAndServe(); err != nil {
		panic(err)
	}
}
