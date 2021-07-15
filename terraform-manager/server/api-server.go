package server

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"

	api "terraformManager/api"

	restful "github.com/emicklei/go-restful"
	"github.com/gorilla/handlers"

	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("api-server")
var serverPort = "9900"

type Server struct {
	api *api.TfAPI
}

func NewServer(api *api.TfAPI) *Server {
	server := &Server{
		api: api,
	}
	return server
}

func (apiServer *Server) RegisterAndServe() error {
	wsContainer := restful.NewContainer()
	wsContainer.RecoverHandler(recoveryHandler)
	wsContainer.DoNotRecover(false)

	err := apiServer.api.Register(wsContainer, true)
	if err != nil {
		return err
	}

	return apiServer.serve(wsContainer)
}

func (apiServer *Server) serve(wsContainer *restful.Container) error {

	log.Printf("start listening on: %s", serverPort)

	server := &http.Server{
		Addr:    ":" + serverPort,
		Handler: handlers.LoggingHandler(os.Stdout, wsContainer),
	}

	return server.ListenAndServe()
}

func recoveryHandler(panicReason interface{}, httpWriter http.ResponseWriter) {
	logger.Errorf("[restful] recover from panic situation: - %v\r\n", panicReason)
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("stack trace:\n"))
	for i := 2; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		buffer.WriteString(fmt.Sprintf("    %s:%d\r\n", path.Base(file), line))
	}
	logger.Debugf(buffer.String())
	httpWriter.WriteHeader(http.StatusInternalServerError)
}
