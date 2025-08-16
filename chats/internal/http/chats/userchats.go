package chats

import (
	"net/http"

	"messenger/chats/internal/domain"
	"messenger/shared/responsewriter"

	"github.com/google/uuid"
)

// UserChats
//
//	@Summary		User's chats
//	@Description	returns user's chats
//	@Tags			Chats
//	@Produce		json
//	@Success		200	"ok"		chatsResponse
//	@Failure		400	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/api/v1/chats [get]
func (h *Handler) UserChats(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	chats, err := h.userChatsUseCase.UserChats(ctx)
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
