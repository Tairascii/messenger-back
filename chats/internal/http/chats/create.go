package chats

import (
	"encoding/json"
	"net/http"

	"messenger/shared/responsewriter"

	"github.com/google/uuid"
)

func (h *Handler) CreateChat(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var payload createPayload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		responsewriter.ErrorResponseWriter(w, err, http.StatusBadRequest)
		return
	}

	userID, err := uuid.Parse(payload.UserID)
	if err != nil {
		responsewriter.ErrorResponseWriter(w, err, http.StatusBadRequest)
		return
	}

	id, err := h.createChatUseCase.CreateChat(ctx, userID)
	if err != nil {
		responsewriter.ErrorResponseWriter(w, err, http.StatusInternalServerError)
		return
	}

	responsewriter.JSONResponseWriter(w, http.StatusOK, mapCreateDomainToResponse(id))
}

type createPayload struct {
	UserID string `json:"user_id"`
}

type createResponse struct {
	ID uuid.UUID `json:"id"`
}

func mapCreateDomainToResponse(resp uuid.UUID) createResponse {
	return createResponse{
		ID: resp,
	}
}
