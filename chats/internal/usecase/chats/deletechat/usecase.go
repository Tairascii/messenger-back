package deletechat

type UseCase struct {
	chatsParticipantsRepo ChatsParticipantsRepo
}

func New(cfg *Config) *UseCase {
	return &UseCase{
		chatsParticipantsRepo: cfg.ChatsParticipantsRepo,
	}
}
