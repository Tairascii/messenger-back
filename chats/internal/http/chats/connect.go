package chats

import (
	"encoding/json"
	"net/http"
	"time"

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

	userID, err := h.canJoinUseCase.CanJoin(ctx, chatID)
	if err != nil {
		responsewriter.ErrorResponseWriter(w, err, http.StatusForbidden)
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

	h.lockUser(chatID, userID, conn)
	defer h.unlockUser(chatID, userID)

	for {
		_, rawMsg, err := conn.ReadMessage()
		if err != nil {
			logger.Log.Errorf("conn.ReadMessage: %s", err.Error())
			return
		}

		var msg Message
		if err = json.Unmarshal(rawMsg, &msg); err != nil {
			logger.Log.Errorf("json.Unmarshal: %s", err.Error())
			return
		}

		logger.Log.Infof("read message: %s\n", msg)
		addedMsg, err := h.addMessageUseCase.AddMessage(ctx, domain.Message{
			Text:      msg.Text,
			IsEdited:  false,
			ChatID:    chatID,
			CreatedAt: time.Now(),
		})
		if err != nil {
			logger.Log.Errorf("addMessageUseCase.AddMessage: %s", err.Error())
			return
		}

		resp, err := json.Marshal(Message(addedMsg))
		if err != nil {
			logger.Log.Errorf("json.Marshal: %s", err.Error())
		}

		h.sendEveryone(chatID, resp)
	}
}

func (h *Handler) lockUser(chatID, userID uuid.UUID, conn *websocket.Conn) {
	h.mu.Lock()
	if _, ok := h.chatConnections[chatID]; !ok {
		h.chatConnections[chatID] = make(map[uuid.UUID]*websocket.Conn)
	}
	h.chatConnections[chatID][userID] = conn
	h.mu.Unlock()
}

func (h *Handler) unlockUser(chatID, userID uuid.UUID) {
	h.mu.Lock()
	delete(h.chatConnections[chatID], userID)
	if len(h.chatConnections[chatID]) == 0 {
		delete(h.chatConnections, chatID)
	}
	h.mu.Unlock()
}

func (h *Handler) sendEveryone(chatID uuid.UUID, data []byte) {
	for userID, conn := range h.chatConnections[chatID] {
		err := conn.WriteMessage(1, data)
		if err != nil {
			logger.Log.Errorf("conn.WriteMessage: chatID=%s, userID=%s", chatID, userID)
		}
	}
}

type Message struct {
	ID        int64     `json:"id"`
	Text      string    `json:"text"`
	IsEdited  bool      `json:"isEdited"`
	CreatedAt time.Time `json:"createdAt"`
	SenderID  uuid.UUID `json:"senderID"`
	ChatID    uuid.UUID `json:"chatID"`
}
