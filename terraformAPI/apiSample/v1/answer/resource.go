package answer

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/emicklei/go-restful"
	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("answer")

type Resource struct {
}

func NewResource() *Resource {
	logger.SetLogLevel(loggo.INFO)
	return &Resource{}
}

func (r *Resource) Register(container *restful.Container) *Resource {
	ws := new(restful.WebService)
	ws.Path("/v1").
		Consumes(restful.MIME_JSON, "application/json").
		Produces(restful.MIME_JSON, "application/json")

	ws.Route(ws.GET("").To(r.GetV1).
		Doc("Returns v1 endpoint").
		Operation("getV1"))

	ws.Route(ws.GET("/validate").To(r.Validate).Operation("validateConfig"))

	container.Add(ws)
	return r
}

func (c *Resource) GetV1(request *restful.Request, response *restful.Response) {
	logger.Infof("GetV1")
	response.WriteHeaderAndEntity(http.StatusOK, "hi! how are you doing?")
}

func (c *Resource) Validate(request *restful.Request, response *restful.Response) {
	r := strings.NewReader("{\"configtext\": \"\"}")
	resp, err := http.Post("http://localhost:9900/validate", "application/json", r)
	if err != nil {
		logger.Errorf("Error while validating. Can not receive response from terraformManager. Reason: " + err.Error())
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	logger.Infof("Recieved from tfManager: %s", string(body))
	response.WriteHeaderAndEntity(http.StatusOK, string(body))
}
