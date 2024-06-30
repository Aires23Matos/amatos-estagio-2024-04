package test_test

import (
	"testing"
	"time"
	"vet-clinic/application/usecase"
	"vet-clinic/repository"
)

func TestEfetuarRonda(t *testing.T){
	t.Run("Deve efetuar ronda no horário entre 8 e 22 horas", func(t *testing.T) {
		//Arrange
		repo := repository.NewRondaRepository()
    	service := usecase.NewServiceRonda(repo)
        horario := time.Date(2024, time.June, 30, 10, 0, 0, 0, time.UTC)
		frequenciaCardiaca := 75
        glicemia := 90.5
        batimento := 70
        temperatura := 37.2
        pressaoArterial := "120/80"
		//Act
        err := service.EfetuarRonda("1", horario, frequenciaCardiaca, glicemia, batimento, temperatura, pressaoArterial)
		//Assert
        if err != nil {
            t.Fatalf("Erro ao efetuar ronda no horário permitido: %v", err)
        }

        ronda, err := repo.BuscarID("1")
        if err != nil {
            t.Fatalf("Erro ao buscar ronda: %v", err)
        }

        if !ronda.Realizada {
            t.Errorf("Esperado que a ronda fosse realizada, mas não foi")
        }
    })
    t.Run("Não deve efetuar ronda fora do horário entre 8 e 22 horas", func(t *testing.T) {
		//Arrange
		repo := repository.NewRondaRepository()
    	service := usecase.NewServiceRonda(repo)
        horario := time.Date(2024, time.June, 30, 6, 0, 0, 0, time.UTC)
		frequenciaCardiaca := 75
        glicemia := 90.5
        batimento := 70
        temperatura := 37.2
        pressaoArterial := "120/80"
		//Act
        err := service.EfetuarRonda("2", horario,frequenciaCardiaca, glicemia, batimento, temperatura, pressaoArterial)
		//Assert
        if err == nil {
            t.Fatalf("Esperado erro ao efetuar ronda fora do horário permitido")
        }

        
        _, err = repo.BuscarID("2")
        if err == nil {
            t.Fatalf("Não era esperado encontrar ronda com ID '2'")
        }
    })

	t.Run("Deve efetuar duas rondas por dia no horário entre 8 e 22 horas", func(t *testing.T) {
		//Arrange
		repo := repository.NewRondaRepository()
    	service := usecase.NewServiceRonda(repo)
        horario1 := time.Date(2024, time.June, 30, 10, 0, 0, 0, time.UTC)
        horario2 := time.Date(2024, time.June, 30, 18, 0, 0, 0, time.UTC)
        frequenciaCardiaca := 75
        glicemia := 90.5
        batimento := 70
        temperatura := 37.2
        pressaoArterial := "120/80"
		//Act
        err := service.EfetuarDuasRondasPorDia(horario1, horario2,frequenciaCardiaca, glicemia, batimento, temperatura, pressaoArterial,frequenciaCardiaca, glicemia, batimento, temperatura, pressaoArterial)
		//Assert
        if err != nil {
            t.Fatalf("Erro ao efetuar duas rondas no horário permitido: %v", err)
        }

        ronda1, err := repo.BuscarID("1")
        if err != nil {
            t.Fatalf("Erro ao buscar a primeira ronda: %v", err)
        }
        if !ronda1.Realizada {
            t.Errorf("Esperado que a primeira ronda fosse realizada, mas não foi")
        }

        ronda2, err := repo.BuscarID("2")
        if err != nil {
            t.Fatalf("Erro ao buscar a segunda ronda: %v", err)
        }
        if !ronda2.Realizada {
            t.Errorf("Esperado que a segunda ronda fosse realizada, mas não foi")
        }
    })
	t.Run("Deve efetuar a ronda com os 5 parâmetros do exame estado geral", func(t *testing.T) {
		//Arragen
		repo := repository.NewRondaRepository()
    	service := usecase.NewServiceRonda(repo)

        horario := time.Date(2024, time.June, 30, 10, 0, 0, 0, time.UTC)
        frequenciaCardiaca := 75
        glicemia := 90.5
        batimento := 70
        temperatura := 37.2
        pressaoArterial := "120/80"
		//Act
        err := service.EfetuarRonda("1", horario, frequenciaCardiaca, glicemia, batimento, temperatura, pressaoArterial)
		//Assert
        if err != nil {
            t.Fatalf("Erro ao efetuar ronda: %v", err)
        }

        ronda, err := repo.BuscarID("1")
        if err != nil {
            t.Fatalf("Erro ao buscar a ronda: %v", err)
        }

        if ronda.FrequenciaCardiaca != frequenciaCardiaca {
            t.Errorf("Esperado frequência cardíaca %d, mas obteve %d", frequenciaCardiaca, ronda.FrequenciaCardiaca)
        }
        if ronda.Glicemia != glicemia {
            t.Errorf("Esperado glicemia %f, mas obteve %f", glicemia, ronda.Glicemia)
        }
        if ronda.Batimento != batimento {
            t.Errorf("Esperado batimento %d, mas obteve %d", batimento, ronda.Batimento)
        }
        if ronda.Temperatura != temperatura {
            t.Errorf("Esperado temperatura %f, mas obteve %f", temperatura, ronda.Temperatura)
        }
        if ronda.PressaoArterial != pressaoArterial {
            t.Errorf("Esperado pressão arterial %s, mas obteve %s", pressaoArterial, ronda.PressaoArterial)
        }
    })
}