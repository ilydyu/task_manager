package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ilydyu/task_manager.git/internal/dto"
	"github.com/ilydyu/task_manager.git/internal/middlware"
	"github.com/ilydyu/task_manager.git/pkg/render"
)

func (h *Handlers) UpdateTask(w http.ResponseWriter, r *http.Request) {
	idx := r.Context().Value(middlware.UserIDContextKey).(string)
	userID, err := strconv.Atoi(idx)

	if err != nil {
		render.Error(w, err, http.StatusBadRequest, "id conversion error")

		return
	}

	taskID, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		render.Error(w, err, http.StatusBadRequest, "invalid id in path")

		return
	}

	var input dto.UpdateTaskInput

	err = json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		render.Error(w, err, http.StatusBadRequest, "json decode error")

		return
	}

	output, err := h.s.UpdateTask(r.Context(), input, int64(userID), int64(taskID))

	if err != nil {
		render.Error(w, err, http.StatusBadRequest, "request failed")

		return
	}

	render.JSON(w, output, http.StatusCreated)
}
