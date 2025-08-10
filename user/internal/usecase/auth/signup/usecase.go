package signup

type UseCase struct {
	userRepo UserRepo
}

func New(cfg *Config) *UseCase {
	return &UseCase{
		userRepo: cfg.UserRepo,
	}
}
