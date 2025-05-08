# Introdução

Em Go, o conceito de "valor nulo" é tratado de forma diferente de outras linguagens. Antes de passarmos para isso vamos ver os valores "default" de cada tipo de variável em Go.

## Valores DEFAULT

| Tipo       | Valor Default                |
|------------|------------------------------|
| INT        | **0**                        |
| FLOAT      | **0.0**                      |
| STRING     | **""**                       |
| BOOL       | **false**                    |
| STRUCT     | **Todos os campos zerados**  |
| PONTEIROS  | **nil**                      |
| SLICES     | **nil**                      |
| FUNÇÕES    | **nil**                      |

## O que é o nil?

**Nil** em Go é o valor zero para ponteiros, slices, maps, channels e funções. Ele representa a ausência de uma valor válido. O **nil** é semelhante ao null do **Java** e do **JavaScripto**.

```
var ptr *int // ponteiro para int (não inicializado)
fmt.Println(ptr) // <nil>

if ptr == nil {
    fmt.Println("ptr é nil!") // Será executado
}
```

# Casos de Uso

Uma dos usos que precisei representar valores nulos em um struct, dessa forma usei ponteiros para os valores que poderiam ser nulos, e os valores que sempre seriam preenchidos com algum valor deixei normal.

- ### Struct com Campo Opcional

```
type Pessoa struct {
    Nome        string
    Idade       int
    Endereco    *string
}
```

Nesse caso, como o Endereco pode ser "null", então eu defino a variável como um ponteiro de string, com isso, se nenhum endereço de memória for atribuído a variável, o valor padrão de endereço será nil.