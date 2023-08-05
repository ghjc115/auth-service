package authService

import (
	"auth-service/structs"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type Handler struct {
	Service *Service
	Logger  *slog.Logger
}

func NewHandler(service *Service, logger *slog.Logger) *Handler {
	return &Handler{
		Service: service,
		Logger:  logger,
	}
}

func Decode(w http.ResponseWriter, r *http.Request, logger *slog.Logger) *structs.User {
	var req structs.User

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to decode request body. %s", err))

		render.JSON(w, r, structs.Response{Status: "ERROR", Error: "Check your JSON"})

		return nil
	}

	return &req
}

func Status(b bool) string {
	status := ""
	if b == true {
		status = "OK"
	} else {
		status = "ERROR"
	}

	return status
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	req := Decode(w, r, h.Logger)

	h.Logger.Info("User body decoded", slog.Any("request", req))

	status := Status(h.Service.AddUser(r.Context(), req))

	render.JSON(w, r, structs.Response{
		Status: status,
	})
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	nickname := chi.URLParam(r, "nickname")

	user := h.Service.GetUser(r.Context(), nickname)
	userResponse := structs.UserResponse{
		UUID:     user.UUID,
		Nickname: user.Nickname,
	}
	status := Status(user != nil)

	render.JSON(w, r, structs.Response{
		Status: status,
		Body:   &userResponse,
	})
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	nickname := chi.URLParam(r, "nickname")

	deleted := h.Service.DeleteUser(r.Context(), nickname)
	status := Status(deleted)

	render.JSON(w, r, structs.Response{
		Status: status,
	})
}
