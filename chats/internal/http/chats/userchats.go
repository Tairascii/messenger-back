package chats

import (
	"net/http"

	"messenger/chats/internal/domain"
	"messenger/shared/responsewriter"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h *Handler) UserChats(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userIDParam := chi.URLParam(r, "user_id")
	userID, err := uuid.Parse(userIDParam)
	if err != nil {
		responsewriter.ErrorResponseWriter(w, err, http.StatusBadRequest)
		return
	}

	chats, err := h.userChatsUseCase.UserChats(ctx, userID)
	if err != nil {
		responsewriter.ErrorResponseWriter(w, err, http.StatusInternalServerError)
		return
	}

	responsewriter.JSONResponseWriter(w, http.StatusCreated, mapDomainToResponse(chats))
}

type Chat struct {
	ID                uuid.UUID `json:"id"`
	LastReadMessageID int64     `json:"last_read_message_id"`
}

type chatsResponse struct {
	Chats []Chat `json:"chats"`
}

func mapDomainToResponse(resp []domain.Chat) chatsResponse {
	chats := make([]Chat, len(resp))
	for i, ch := range resp {
		chats[i] = Chat{
			ID:                ch.ID,
			LastReadMessageID: ch.LastReadMessageID,
		}
	}
	return chatsResponse{
		Chats: chats,
	}
}
