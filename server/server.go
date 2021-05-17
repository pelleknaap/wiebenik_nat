
package server

import (
	"github.com/gorilla/mux"
	"github.com/pelleknaap/wiebeniknat_backend/models"
	"github.com/rs/cors"
	"gorm.io/gorm"
	"net/http"
)

type Server struct {
	DB      *gorm.DB
	Handler http.Handler
	Router  *mux.Router
	Config  models.Config
}

func(s *Server) routes() {
	s.Router = mux.NewRouter()

}

func (s *Server) setupHandler() {
	s.Handler = cors.Default().Handler(s.Router)
}

//func (s *Server) openDB(){
//	dsn := s.Config.Yaml.DbDSN
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic(fmt.Sprintf("mysql open err: %s", err))
//	}
//
//	s.DB = db
//	err = db.AutoMigrate(&models.Customer{})
//	if err != nil {
//		panic(fmt.Sprintf("auto migration err: %s", err))
//	}
//}

func NewServer() (*Server, error) {
	s := &Server{}
	s.routes()
	err := s.Config.Parse()
	s.setupHandler()
	return s, err
}