package users

type userUsecase struct {
	userRepository Repository
}

func NewUserUsecase(ur Repository) Usecase {
	return &userUsecase{
		userRepository: ur,
	}
}

func (uu *userUsecase) GetAll() []Domain {
	return uu.userRepository.GetAll()
}

func (uu *userUsecase) Create(noteDomain *Domain) Domain {
	return uu.userRepository.Create(noteDomain)
}
