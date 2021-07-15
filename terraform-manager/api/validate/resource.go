package validate

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("validate")

type Config struct {
	ConfigText string `json:"configText"`
}

type Resource struct {
}

func NewResource() *Resource {
	logger.SetLogLevel(loggo.INFO)
	return &Resource{}
}

func (r *Resource) Register(container *restful.Container) *Resource {
	ws := new(restful.WebService)

	ws.Route(ws.POST("validate").To(r.validateConfig).
		Doc("Validate config").
		Operation("validateConfig").
		Reads(Config{}))

	container.Add(ws)
	return r
}

func (c *Resource) validateConfig(request *restful.Request, response *restful.Response) {
	logger.Infof("Validating config...")
	response.WriteHeaderAndEntity(http.StatusOK, "feature in development...")
	// config := &Config{}
	// err := request.ReadEntity(config)

	// var terraform = terraform.Terraform{}
	// terraform.AddScript(config.ConfigText)

	// if err != nil {
	// 	logger.Errorf("Cannot read body %v", err)
	// 	response.WriteHeaderAndEntity(http.StatusBadRequest, "Can not read request body")
	// 	return
	// }

	// //validationAnswer, err := terraform.Validate()
	// if err != nil {
	// 	logger.Errorf("Terraform can not validate: ", err.Error())
	// 	//response.WriteHeaderAndEntity(http.StatusBadGateway, "")
	// 	return
	// }
	//logger.Infof("Send validation answer: %s", validationAnswer)
	//response.WriteAsJson(validationAnswer)
}
