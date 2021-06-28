package protodoc

import (
	iconf "github.com/aclisp/protodoc-serv/internal/config"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

const (
	ServiceName = "protodoc"
)

type Server struct {
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
	return s.Namespace + ".service." + ServiceName
}

func (s *Server) Run() {
	log.Init(log.WithFields(map[string]interface{}{"service": ServiceName}))

	// initialize the micro service
	var srvOpts []micro.Option
	srvOpts = append(srvOpts, micro.Name(s.Name()))
	service := micro.NewService(srvOpts...)

	RegisterProtoDocHandler(service.Server(), new(ProtoDoc))

	// Run micro server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
