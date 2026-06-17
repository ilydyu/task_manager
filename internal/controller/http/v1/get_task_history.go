package v1

import (
	"net/http"
	"strconv"

	"github.com/ilydyu/task_manager.git/pkg/render"
)

func (h *Handlers) GetTaskHistory(w http.ResponseWriter, r *http.Request) {
	taskID, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		render.Error(w, err, http.StatusBadRequest, "invalid id in path")

		return
	}

	output, err := h.s.GetTaskHistory(r.Context(), int64(taskID))

	if err != nil {
		render.Error(w, err, http.StatusBadRequest, "request failed")

		return
	}

	render.JSON(w, output, http.StatusOK)
}
