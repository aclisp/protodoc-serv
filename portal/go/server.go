package portal

import (
	"net/http"

	iconf "github.com/aclisp/protodoc-serv/internal/config"
	"github.com/aclisp/protodoc-serv/portal/bundle"
	"github.com/gorilla/mux"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
)

const (
	ServiceName = "portal"
)

type Server struct {
	Address   string `json:"portal_serving_address"`
	Namespace string `json:"namespace"`
}

func NewServer() *Server {
	s := &Server{}
	if err := iconf.LoadInitialConfigFromFile("serverinit.yaml", s); err != nil {
		panic(err)
	}
	return s
}

func (s *Server) Name() string {
	return s.Namespace + "." + ServiceName
}

func (s *Server) Run() {
	log.Init(log.WithFields(map[string]interface{}{"service": ServiceName}))

	// init url router
	r := mux.NewRouter()

	// setup route of html pages
	r.PathPrefix("/").Handler(http.FileServer(bundle.Assets))

	// start micro web service
	service := web.NewService(
		web.Name(s.Name()),
		web.Address(s.Address),
		web.Handler(r),
	)

	// run micro web server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
