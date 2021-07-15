package apiSample

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"

	v1 "terraformAPI/apiSample/v1"

	restful "github.com/emicklei/go-restful"
	"github.com/gorilla/handlers"

	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("api-server")
var serverPort = "9090"

type Server struct {
	api *v1.SampleAPI
}

func NewServer(api *v1.SampleAPI) *Server {
	server := &Server{
		api: api,
	}
	return server
}

func (apiServer *Server) RegisterAndServe() error {
	wsContainer := restful.NewContainer()
	wsContainer.RecoverHandler(recoveryHandler)
	wsContainer.DoNotRecover(false)

	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"Content-Length", "Content-Disposition"},
		AllowedHeaders: []string{restful.HEADER_ContentType, restful.HEADER_Accept, "Authorization", "Cache-Control", "Pragma"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		MaxAge:         365 * 24 * 60 * 60,
		CookiesAllowed: true,
		Container:      wsContainer}
	wsContainer.Filter(cors.Filter)

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
