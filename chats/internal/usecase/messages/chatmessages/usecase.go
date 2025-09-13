package chatmessages

type UseCase struct {
	messageRepo           MessagesRepo
	chatsParticipantsRepo ChatsParticipantsRepo
}

func New(cfg *Config) *UseCase {
	return &UseCase{
		messageRepo:           cfg.MessagesRepo,
		chatsParticipantsRepo: cfg.ChatsParticipantsRepo,
	}
}

