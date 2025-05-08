# Intrudução

Em Go, o único laço de repetição que temos é o **for**. Então há maneiras de utilizá-lo para subistituir o **while** e **do-while**

Tradicionalmente podemos usar o for destas maneiras: 

## For Normal

Aqui seria a forma de uso normal de um sor, ou seja, usando ele para iterar sobre alguma coisa.

```
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

Saída:

0
1
2
3
4
```

## For While

Dessa forma aqui nós não incializamos o for, apenas definimos uma condição de parada e iniciamos a execução.

```
var := 0
for var < 5 {
    fmt.Println(var)
    var++
}
```

## For Infinito

Nesse caso aqui o For fica num Loop, até cair em uma condição especifica. Aqui faremos o uso de **continue** e **break**

```
k := 0
for {
    fmt.Println("Loop infinito", k)
    k++
    if k == 2 {
        break // Sai do loop
    }
}
```

Esse tipo de For é equivalente a uma **while(true)**. Quando a condição for atendida o break é acioado é o loop acaba.

## For com Range

Esse For nós usamos para iterar sobre coleções, númericas, strings, structs... A lista é imensa

```
numeros := []int{10, 20, 30}
for indice, valor := range numeros {
    fmt.Printf("Índice: %d, Valor: %d\n", indice, valor)
}

Saída:

Índice: 0, Valor: 10
Índice: 1, Valor: 20
Índice: 2, Valor: 30
```

Ainda podemos utilizar o undercore *_*, ou sublinhado, para ignorarmos o índice ou o valor. Isso evitar erros já que uma variável no Go se for declarada precisa ser usada. Nesse nosso exemplo ficaria assim.

```
numeros := []int{10, 20, 30}
for _, valor := range numeros {
    fmt.Printf("Valor: %d\n", valor)
}

Saída:

Valor: 10
Valor: 20
Valor: 30
```

## Iteração Unicode

Aqui podemos iterar letra a letra de uma palavra

```
for i, c := range "Go" {
    fmt.Printf("Posição: %d, Caractere: %c\n", i, c)
}

Saída:

Posição: 0, Caractere: G
Posição: 1, Caractere: o
```


# Break e Continue


Break e Continue são usados para o controle de fluxo, ou seja, vamos supor que uma condição é atendida, se queremos para a execução do **for** então passamos o break na condição para encerrar o programa. 

Com o continue a situação é semelhante, mas a execução do programa não é encerrada, vamos supor que a condição é encontrar um valor, e então não precisa mais verificar, então usamos um continue e a execução será pulada para a próxima iteração

Exemplos:

- Break
```
for i := 0; i < 5; i++ {
    if i == 3 {
        break
    }
    fmt.Println(i)
}

Saída:

0
1
2
```

- Continue
```
for i := 0; i < 5; i++ {
    if i == 2 {
        continue
    }
    fmt.Println(i)
}

Saída:

0
1
3
4
```