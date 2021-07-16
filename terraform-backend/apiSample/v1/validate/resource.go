package validate

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/emicklei/go-restful"
	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("validate")

type Resource struct {
	tfHostname string
}

func NewResource(hostname string) *Resource {
	logger.SetLogLevel(loggo.INFO)
	logger.Infof("Creating new resource for validate with hostname=%s....", hostname)
	return &Resource{hostname}
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
	resp, err := http.Post(c.tfHostname+"/validate", "application/json", r)
	if err != nil {
		logger.Errorf("Error while validating. Can not receive response from terraformManager. Reason: " + err.Error())
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	logger.Infof("Recieved from tfManager: %s", string(body))
	response.WriteHeaderAndEntity(http.StatusOK, "terraform-manager answer is "+string(body))
}
