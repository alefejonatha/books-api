package server

import (
	"log"

	"github.com/alefejonatha/Http/restapi/server/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine //É o que precisa para iniciar o server
}

func NewServer() Server { //"Construtor" do server
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Print("Server is running at port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}
