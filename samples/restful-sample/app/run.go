package app

import (
	apisample "restfulSample/apiSample"

	v1 "restfulSample/apiSample/v1"

	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("app")

func Run() {
	logger.SetLogLevel(loggo.INFO)

	sampleAPI := v1.NewSampleAPI()

	sampleServer := apisample.NewServer(sampleAPI)
	if err := sampleServer.RegisterAndServe(); err != nil {
		panic(err)
	}
}
