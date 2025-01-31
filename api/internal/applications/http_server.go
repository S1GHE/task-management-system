package applications

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"task-management-system-api/internal/applications/middleware"
	"task-management-system-api/internal/applications/projects"
	"task-management-system-api/internal/applications/users"
	userInfra "task-management-system-api/internal/infrastructure/users"
)

type HTTPServer struct {
	userRepo *userInfra.PostgresRepo
}

func NewHTTPServer(userRepo *userInfra.PostgresRepo) *HTTPServer {
	return &HTTPServer{
		userRepo: userRepo,
	}
}

func (s *HTTPServer) SetupHTTPServer() http.Handler {
	mux := http.NewServeMux()
	validate := validator.New()

	projectsHandler := projects.SetupHandler(validate)
	userHandler := users.SetupHandlers(validate, s.userRepo)

	mux.HandleFunc("POST /register/user", userHandler.RegisterUser)
	mux.HandleFunc("POST /login/user", userHandler.LoginUser)

	protectedMux := http.NewServeMux()
	protectedMux.HandleFunc("GET /api/projects", projectsHandler.GetProjects)
	protectedMux.HandleFunc("POST /api/projects", projectsHandler.CreateProject)
	protectedMux.HandleFunc("PATCH /api/projects/{project_id}/tasks", projectsHandler.CreateProjectTask)

	protectedRoutes := middleware.JWTMiddleware(protectedMux)

	finalMux := http.NewServeMux()
	finalMux.Handle("/", mux)
	finalMux.Handle("/api/", protectedRoutes)

	return finalMux
}
