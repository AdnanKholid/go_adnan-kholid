package blogs

type blogsUsecase struct {
	blogsRepository Repository
}

func NewBlogsUsecase(br Repository) Usecase {
	return &blogsUsecase{
		blogsRepository: br,
	}
}

func (bu *blogsUsecase) GetAll() []Domain {
	return bu.blogsRepository.GetAll()
}
func (bu *blogsUsecase) GetByID(id string) Domain {
	return bu.blogsRepository.GetByID(id)
}
func (bu *blogsUsecase) Create(noteDomain *Domain) Domain {
	return bu.blogsRepository.Create(noteDomain)
}
func (bu *blogsUsecase) Update(id string, noteDomain *Domain) Domain {
	return bu.blogsRepository.Update(id, noteDomain)
}
func (bu *blogsUsecase) Delete(id string) bool {
	return bu.blogsRepository.Delete(id)
}

func (bu *blogsUsecase) GetByTitle(title string) Domain {
	return bu.blogsRepository.GetByTitle(title)
}

func (bu *blogsUsecase) GetByCategory(idCategory string) []Domain {
	return bu.blogsRepository.GetByCategory(idCategory)
}
