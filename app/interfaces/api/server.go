package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/justinas/alice"
	"github.com/rs/cors"

	"guild-hack-api/app/db"
	"guild-hack-api/app/interfaces/api/handler"
	"guild-hack-api/app/interfaces/api/middleware"
	"guild-hack-api/app/usecase"
)

type Server struct {
	router *mux.Router
	db     *sqlx.DB
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init() error {
	cs := db.NewMySQL("guild-hack:guild-hack@tcp(db:3306)/guild_hack?parseTime=true")
	dbcon, err := cs.Open()
	if err != nil {
		return fmt.Errorf("failed db init. %s", err)
	}
	s.db = dbcon
	s.router = s.Route()
	return nil
}

func (s *Server) Run(port int) {
	log.Printf("Listening on port %d", port)
	err := http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		handlers.CombinedLoggingHandler(os.Stdout, s.router),
	)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Route() *mux.Router {
	authMiddleware := middleware.NewAuth(s.db)
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Authorization", "Accept-Language", "Content-Type", "Content-Language", "Origin"},
		AllowedMethods: []string{
			http.MethodOptions,
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
	})

	commonChain := alice.New(
		middleware.RecoverMiddleware,
		corsMiddleware.Handler,
	)

	authChain := commonChain.Append(
		authMiddleware.Handler,
	)

	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/ping").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ping ok !!"))
	})

	userUseCase := usecase.NewUserUseCase(s.db)
	userHandler := handler.NewUserHandler(userUseCase)
	recruitUseCase := usecase.NewRecruitUseCase(s.db)
	recruitHandler := handler.NewRecruitHandler(recruitUseCase)

	v1r := r.PathPrefix("/v1").Subrouter()
	v1r.Methods(http.MethodPost, http.MethodOptions).Path("/users").Handler(commonChain.Then(AppHandler{userHandler.Create}))
	v1r.Methods(http.MethodGet, http.MethodOptions).Path("/users").Handler(commonChain.Then(AppHandler{userHandler.Index}))
	v1r.Methods(http.MethodGet, http.MethodOptions).Path("/users/{user_id}").Handler(commonChain.Then(AppHandler{userHandler.Show}))

	v1r.Methods(http.MethodPost, http.MethodOptions).Path("/recruits").Handler(authChain.Then(AppHandler{recruitHandler.Create}))

	return r
}
