package v1

import (
	"net/http"

	"github.com/ilydyu/task_manager.git/pkg/render"
)

func (h *Handlers) GetTasks(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("team_id")
	assigneeID := r.URL.Query().Get("assignee_id")
	status := r.URL.Query().Get("status")
	cursor := r.URL.Query().Get("cursor")

	output, err := h.s.GetTasks(r.Context(), teamID, assigneeID, status, cursor)

	if err != nil {
		render.Error(w, err, http.StatusBadRequest, "request failed")

		return
	}

	render.JSON(w, output, http.StatusOK)
}
