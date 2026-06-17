package v1

import (
	"net/http"
	"strconv"

	"github.com/ilydyu/task_manager.git/internal/middlware"
	"github.com/ilydyu/task_manager.git/pkg/render"
)

func (h *Handlers) GetUserTeams(w http.ResponseWriter, r *http.Request) {
	idx := r.Context().Value(middlware.UserIDContextKey).(string)
	id, err := strconv.Atoi(idx)

	if err != nil {
		render.Error(w, err, http.StatusBadRequest, "id conversion error")

		return
	}

	output, err := h.s.GetUserTeams(r.Context(), id)

	if err != nil {
		render.Error(w, err, http.StatusBadRequest, "request failed")

		return
	}

	render.JSON(w, output, http.StatusOK)
}
