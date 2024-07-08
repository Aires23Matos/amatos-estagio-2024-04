### Calculo do IRT
## Instrução para executar o programa 

1. Para gerar o  executavel

```
  go build -o cirt 
```   

1.2 Para sistema windows:

```
  go build -o cirt.exe 
```

2. para executar a aplicação deve executar o seguinte comando:

```
 ./cirt-linux cirt -d [dias úteis de trabalho] -f [número de faltas] -a [subsídio de alimentção] -t [subsídio de transporte] -s [salário base]
```

```
  Ex.: ./cirt.exe cirt   30  1  1000  1000   30000
```
2.1 Ou

```
   ./cirt --diasuteis [dias úteis de trabalho] --falta [número de falta]--subalimentacao [subsídio de alimentção] --subtransporte [subsídio de transporte] --salariobase [salário base]
```

```
  Ex.: ./cirt.exe cirt --diasuteis 30 --falta 2 --subalimentacao 500 --subtransporte 500  --salariobase 30000
```