package applications

import (
	"net/http"
	"task-management-system-api/internal/applications/projects"
	"task-management-system-api/internal/applications/user"
)

type HTTPServer struct {
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{}
}

func (s *HTTPServer) SetupHTTPServer() http.Handler {
	mux := http.NewServeMux()

	projectsHandler := projects.SetupHandler()
	userHandler := user.SetupHandlers()

	mux.HandleFunc("GET /api/projects", projectsHandler.GetProjects)
	mux.HandleFunc("POST /api/projects", projectsHandler.CreateProject)

	mux.HandleFunc("PATCH /api/projects/{project_id}/tasks", projectsHandler.CreateProjectTask)

	mux.HandleFunc("POST /api/register/user", userHandler.RegisterUser)
	mux.HandleFunc("POST /api/login/user", userHandler.LoginUser)

	return mux
}
