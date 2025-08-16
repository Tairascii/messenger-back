package chats

import (
	"net/http"

	"messenger/shared/responsewriter"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// DeleteChat
//
//	@Summary		Delete chat
//	@Description	deletes user's chat but not completely
//	@Tags			Chats
//	@Produce		json
//	@Param			chat_id	path	uuid	true	"Chat ID"
//	@Success		200		"ok"
//	@Failure		400		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/api/v1/chats/{chat_id} [delete]
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
