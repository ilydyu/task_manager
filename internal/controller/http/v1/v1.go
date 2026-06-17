package v1

import "github.com/ilydyu/task_manager.git/internal/service"

type Handlers struct {
	s *service.Service
}

func New(s *service.Service) *Handlers {
	return &Handlers{
		s: s,
	}
}
