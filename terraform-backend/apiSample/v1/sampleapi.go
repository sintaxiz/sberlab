package v1

import (
	"log"
	"os"
	"terraformAPI/apiSample/v1/upload"
	validate "terraformAPI/apiSample/v1/validate"

	"github.com/emicklei/go-restful"
	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("SampleAPI")

type SampleAPI struct {
}

func NewSampleAPI() *SampleAPI {
	logger.SetLogLevel(loggo.INFO)

	api := &SampleAPI{}

	restful.DefaultRequestContentType(restful.MIME_JSON)
	restful.DefaultResponseContentType(restful.MIME_JSON)
	restful.SetLogger(log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds))

	return api
}

func (api *SampleAPI) Register(wsContainer *restful.Container, insecure bool) error {
	wsContainer.Filter(corsFilter)

	validate.NewResource(os.Getenv("TF_MANAGER_HOST")).Register(wsContainer)
	upload.NewResource().Register(wsContainer)

	return nil
}

func corsFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	chain.ProcessFilter(req, resp)
}
