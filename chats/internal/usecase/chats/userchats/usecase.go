package userchats

type UseCase struct {
	chatsRepo ChatsRepo
}

func New(cfg *Config) *UseCase {
	return &UseCase{
		chatsRepo: cfg.ChatsRepo,
	}
}
