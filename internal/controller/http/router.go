package http

import (
	"github.com/go-chi/chi/v5"
	v1 "github.com/ilydyu/task_manager.git/internal/controller/http/v1"
	"github.com/ilydyu/task_manager.git/internal/logger"
	"github.com/ilydyu/task_manager.git/internal/middlware"
	"github.com/ilydyu/task_manager.git/internal/service"
)

func Router(r *chi.Mux, s *service.Service, secret string) {
	v1 := v1.New(s)
	m := middlware.NewMiddlware(secret)

	r.Route("/api", func(r chi.Router) {
		r.Use(logger.Middleware)

		r.Route("/v1", func(r chi.Router) {
			// Auth
			r.Post("/register", v1.CreateUser)
			r.Post("/login", v1.Login)

			// Protected
			r.Group(func(r chi.Router) {
				r.Use(m.AuthMiddleware())

				// Teams
				r.Get("/teams", v1.GetUserTeams)
				r.Post("/teams", v1.CreateTeam)
				r.Post("/teams/{id}/invite", v1.Invite)

				// Tasks
				r.Get("/tasks/{id}/history", v1.GetTaskHistory)
				r.Get("/tasks", v1.GetTasks) // Предпологаю, что возвращать надо все задачи, клиент сам отбирает нужные по фильтрации
				r.Post("/tasks", v1.CreateTask)
				r.Put("/tasks/{id}", v1.UpdateTask)

				// Complex query
				r.Get("/teams/stats", v1.GetTeamStats)
				r.Get("/teams/top_creators", v1.GetTopCreators)
				r.Get("/tasks/invalid_assignments", v1.GetInvalidAssignments)
			})
		})
	})
}
