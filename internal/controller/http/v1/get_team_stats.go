package v1

import (
	"net/http"

	"github.com/ilydyu/task_manager.git/pkg/render"
)

func (h *Handlers) GetTeamStats(w http.ResponseWriter, r *http.Request) {
	output, err := h.s.GetTeamStats(r.Context())

	if err != nil {
		render.Error(w, err, http.StatusBadRequest, "request failed")

		return
	}

	render.JSON(w, output, http.StatusOK)
}
