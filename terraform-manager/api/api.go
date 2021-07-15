package api

import (
	"log"
	"os"
	"terraformManager/api/validate"

	restful "github.com/emicklei/go-restful"
	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("TfAPI")

type TfAPI struct {
}

func NewTfAPI() *TfAPI {
	logger.SetLogLevel(loggo.INFO)

	api := &TfAPI{}

	restful.DefaultRequestContentType(restful.MIME_JSON)
	restful.DefaultResponseContentType(restful.MIME_JSON)
	restful.SetLogger(log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds))

	return api
}

func (api *TfAPI) Register(wsContainer *restful.Container, insecure bool) error {
	validate.NewResource().Register(wsContainer)
	return nil
}
