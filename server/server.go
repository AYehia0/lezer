package server

import (
	"fmt"

	"github.com/AYehia0/lezer/utils"
	"github.com/gin-gonic/gin"
)

// all the route and secret configs
type Config struct {
	BrowseRoute string
}

type Server struct {
	Addr             string
	Port             string
	ServerDir        string // the given path by the user
	CurrentDir       string // the current dir/path in the browser
	Ip               string
	DownloadLocation string
	IgnoreHidden     bool
	router           *gin.Engine
	Config           Config
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

	// Gin
	gin.SetMode(gin.ReleaseMode)

	// the router
	router := gin.New()

	// serving static files
	t, err := utils.LoadTemplate()
	if err != nil {
		panic(err)
	}
	router.SetHTMLTemplate(t)
	// router.LoadHTMLGlob("static/*.html")
	router.Use(noCacheMiddleware())

	// defining the routes here
	// TODO: There must be something better than *, cons : / at the end, allows multiple ///
	router.GET("/", s.handleMainIndex)
	router.GET(fmt.Sprintf("/%s/*dir", s.Config.BrowseRoute), s.serveFiles)
	router.POST("/share", s.uploadFiles)

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
