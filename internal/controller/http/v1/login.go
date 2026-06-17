package v1

import (
	"encoding/json"
	"net/http"

	"github.com/ilydyu/task_manager.git/internal/dto"
	"github.com/ilydyu/task_manager.git/pkg/render"
)

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	input := dto.LoginInput{}

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		render.Error(w, err, http.StatusBadRequest, "json decode error")

		return
	}

	output, err := h.s.Login(r.Context(), input)

	if err != nil {
		render.Error(w, err, http.StatusBadRequest, "request failed")

		return
	}

	render.JSON(w, output, http.StatusOK)
}
