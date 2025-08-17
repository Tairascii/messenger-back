package chats

import (
	"net/http"

	"messenger/chats/internal/domain"
	"messenger/shared/logger"
	"messenger/shared/responsewriter"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func (h *Handler) ConnectToChat(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	chatIDParam := chi.URLParam(r, "chat_id")
	chatID, err := uuid.Parse(chatIDParam)
	if err != nil {
		responsewriter.ErrorResponseWriter(w, err, http.StatusBadRequest)
		return
	}

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		responsewriter.ErrorResponseWriter(w, err, http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			logger.Log.Errorf("conn.ReadMessage: %s", err.Error())
			return
		}

		logger.Log.Infof("read message: %s\n", msg)
		err = h.addMessageUseCase.AddMessage(ctx, domain.Message{
			Text:     "", // todo take from msg
			IsEdited: false,
			ChatID:   chatID,
		})
		if err != nil {
			logger.Log.Errorf("addMessageUseCase.AddMessage: %s", err.Error())
			return
		}
	}
}
