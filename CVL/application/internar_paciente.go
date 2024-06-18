package application

import "cvl/domain"

type ServiceCVL struct {
	repo *domain.FichaInternamentoRepository
}

func CVLService(repo *domain.FichaInternamentoRepository) *ServiceCVL {
	return &ServiceCVL{
		repo: repo,
	}
}

func (f *ServiceCVL) FichaDeInternamento() error {
	ficha := &domain.FichaDeInternamento{
		Id: "1",
	}
	return f.repo.Guardar(ficha)
}
