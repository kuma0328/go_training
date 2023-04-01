package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kuma0328/go_training/usecase"
)

type TodoHandler struct {
	tu usecase.TodoUsecaseImpl
}

func NewTodoHanlder(tu usecase.TodoUsecaseImpl) *TodoHandler {
	return &TodoHandler{
		tu : tu,
	}
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var jsonBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&jsonBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	
}