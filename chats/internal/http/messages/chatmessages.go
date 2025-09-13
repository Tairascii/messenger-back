package messages

import (
	"net/http"
	"time"

	"messenger/chats/internal/domain"
	"messenger/shared/responsewriter"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h *Handler) ChatMessages(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	chatIDParam := chi.URLParam(r, "chat_id")
	chatID, err := uuid.Parse(chatIDParam)
	if err != nil {
		responsewriter.ErrorResponseWriter(w, err, http.StatusBadRequest)
		return
	}

	messages, err := h.chatMessagesUseCase.ChatMessages(ctx, chatID)
	if err != nil {
		responsewriter.ErrorResponseWriter(w, err, http.StatusInternalServerError)
		return
	}

	responsewriter.JSONResponseWriter(w, http.StatusOK, mapDomainToResponse(messages))
}

type chatMessagesResponse struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	ID        int64     `json:"id"`
	Text      string    `json:"text"`
	IsEdited  bool      `json:"isEdited"`
	CreatedAt time.Time `json:"createdAt"`
	SenderID  uuid.UUID `json:"senderID"`
	ChatID    uuid.UUID `json:"chatID"`
}

func mapDomainToResponse(resp []domain.Message) chatMessagesResponse {
	messages := make([]Message, len(resp))
	for i, msg := range resp {
		messages[i] = Message(msg)
	}

	return chatMessagesResponse{
		Messages: messages,
	}
}
