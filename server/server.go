package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Addr         string
	Port         string
	ServerDir    string // the given path by the user
	CurrentDir   string // the current dir/path in the browser
	Ip           string
	IgnoreHidden bool
	router       *gin.Engine
}

func noCacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set Cache-Control headers
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")

		c.Next()
	}
}

// creates the gin engine server
func (s *Server) setupServer() {
	router := gin.Default()

	// serving static files
	router.LoadHTMLGlob("./static/*.html")
	router.Use(noCacheMiddleware())

	// defining the routes here
	// TODO: There must be something better than *, cons : / at the end, allows multiple ///
	router.GET("/browse/*dir", s.serveFiles)

	s.router = router
}

// start the gin server
func (s *Server) StartServer() {
	s.setupServer()
	s.router.Run(fmt.Sprintf("%s:%s", s.Addr, s.Port))
}

func NewServer(addr, port string) *Server {
	return &Server{
		Addr: addr,
		Port: port,
	}
}
