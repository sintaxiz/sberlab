package v1

import (
	"log"
	"os"
	answer "restfulSample/apiSample/v1/answer"

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

	answer.NewResource().Register(wsContainer)

	return nil
}

func corsFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	chain.ProcessFilter(req, resp)
}
