package domain

import "errors"

type FichaInternamentoRepository struct {
	fichas map[string]*FichaDeInternamento
}

func NewFichaInternementorepository() *FichaInternamentoRepository {
	return &FichaInternamentoRepository{
		fichas: make(map[string]*FichaDeInternamento),
	}
}

func (t *FichaInternamentoRepository)Tamanho() int{
	return len(t.fichas)
}

func (s *FichaInternamentoRepository) Guardar(ficha *FichaDeInternamento) error {
	if _, f := s.fichas[ficha.Id]; f {
		return errors.New("já existe ficha de internamento")
	}
	s.fichas[ficha.Id] = ficha
	return nil
}

func (r *FichaInternamentoRepository) Buscar(id string) (*FichaDeInternamento, error) {
    if ficha, b := r.fichas[id]; b {
        return ficha, nil
    }
    return nil, errors.New("ficha de internamento não encontrada")
}