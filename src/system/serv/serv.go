package serv

import (
	"log"
	"net/http"
	"os"
	"time"

	"groupup/src/models"
	"groupup/src/system/router"

	"github.com/go-xorm/xorm"
	"github.com/gorilla/handlers"
)

// Server stores port and database
type Server struct {
	port string
	Db   *xorm.Engine
}

// NewServer returns a new instance of Server
func NewServer() Server {
	return Server{}
}

// Init all vals
func (s *Server) Init(port string, db *xorm.Engine) {
	log.Println("Initializing server...")
	s.port = ":" + port
	s.Db = db
}

// Start the server
func (s *Server) Start() {
	log.Println("Starting server on port" + s.port)

	models.Init()

	r := router.NewRouter()

	r.Init(s.Db)

	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Cache-Control", "X-App-Token"}),
		handlers.ExposedHeaders([]string{""}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(r.Router))
	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)

	newServer := &http.Server{
		Handler:      handler,
		Addr:         "0.0.0.0" + s.port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(newServer.ListenAndServe())
}
