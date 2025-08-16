package chats

import (
	"net/http"

	"messenger/shared/responsewriter"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h *Handler) DeleteChat(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	chatIDParam := chi.URLParam(r, "chat_id")
	chatID, err := uuid.Parse(chatIDParam)
	if err != nil {
		responsewriter.ErrorResponseWriter(w, err, http.StatusBadRequest)
		return
	}

	err = h.deleteChatUseCase.DeleteByID(ctx, chatID)
	if err != nil {
		responsewriter.ErrorResponseWriter(w, err, http.StatusInternalServerError)
		return
	}

	responsewriter.JSONResponseWriter(w, http.StatusNoContent, nil)
}
