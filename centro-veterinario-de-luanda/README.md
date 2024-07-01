# Centro Veterinário de Luanda (C.V.L)

Projecto de aprendizado no estágio Zafir

## Autores
1. Aires de Matos
2. Manuel Afonso



## Descrição

O Centro Veterinário de Luanda (C.V.L) é uma instituição especializada nos cuidados de saúde animal, oferecendo diversos serviços e produtos que garantem o bem estar e a saúde dos nossos pacientes.

Neste momento enfrentamos muita dificuldade com o processo de internamento dos pacientes, ele é feito de forma manual, por intermédio de uma ficha, onde o técnico da entrada do paciente para a sala de internamento. Esta ficha é disponibilizada ao médico veterinário responsável do turno para averiguar o estado geral do paciente, e sobre ela, em cada turno, o médico responsável monitora o estado geral do paciente até que o internamento seja encerrado.

Este processo manual resulta em ineficiências na transferência de informações entre os técnicos e os médicos veterinários responsáveis, além de dificultar a monitorização contínua dos pacientes.

Solicitamos um software dedicado que facilita o processo de internamento e uniformizar as rondas diárias sobre o estado geral dos pacientes.


### Instalação

* Como baixar/clonar o programa
* Qualquer modificação precisa ser feito nos ficheiros/pastas

### Execução do Programa

* Como correr o programa
* Passo por passo

* Compilação
você pode usar o comando go build com variáveis de ambiente


Compilação para: 
- Windows
- Linux
- Mac OS

## Executável
1. Windows
```
go build -o vet-clinic-windows.exe
```
2. Linux
```
go build -o vet-clinic-linux
```
3. Mac OS
```
go build -o vet-clinic-macos
```
4. Mac OS arm64
```
go build -o vet-clinic-macos-arm64
```
5. Linux arm
```
go build -o vet-clinic-linux-arm
```

* Execução do CLI
Após a compilação, você pode executar o CLI no terminal

Exemplo: ./vet-clinic-macos  internar --pacienteID "1" --tutorID "1" --especie "Canino" --diagnostico "Pneumonia" --queixas "Tosse persistente" Fido 
```
./vet-clinic-macos  internar
```

Exemplo: ./vet-clinic-windows efetuar_ronda --id 1 --frequenciaCardiaca 75 --glicemia 90.5 --batimento 70 --temperatura 37.2 --pressaoArterial "120/80"

```
./vet-clinic-windows efetuar_ronda
```





