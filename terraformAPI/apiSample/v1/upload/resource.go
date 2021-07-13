package upload

import (
	"io/ioutil"

	"github.com/emicklei/go-restful"
	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("upload")

type Resource struct {
}

func NewResource() *Resource {
	logger.SetLogLevel(loggo.INFO)
	return &Resource{}
}

func (r *Resource) Register(container *restful.Container) *Resource {
	ws := new(restful.WebService)

	ws.Route(ws.POST("/upload").To(r.UploadScript).Operation("uploadScript"))

	container.Add(ws)
	return r
}

func (r *Resource) UploadScript(request *restful.Request, response *restful.Response) {
	logger.Infof("Received new script")

	request.Request.ParseMultipartForm(10 << 20) // file can be not more than 10mb
	file, handler, err := request.Request.FormFile("myScript")
	if err != nil {
		logger.Errorf("Error retrieving file from form-data: %s", err)
		return
	}
	defer file.Close()

	logger.Infof("Uploaded script: %+v", handler.Filename)
	logger.Infof("Script size: %+v", handler.Size)
	logger.Infof("MIME Header: %+v\n", handler.Header)

	tmpScript, err := ioutil.TempFile("temp-scripts", "upload-*.txt")
	if err != nil {
		logger.Errorf("Can not create temp file for script: %s", err)
		return
	}
	defer tmpScript.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		logger.Errorf("Can not read file: %s", err)
		return
	}
	tmpScript.Write(fileBytes)
	logger.Infof("Successfully loading a script")

}
