package main

import (
	"fmt"

	"github.com/abuffon/alura-go/banco/contas"
)

func main() {
	conta := contas.ContaCorrente{Titular: "Anderson", NumeroAgencia: 1453, NumeroConta: 27979, Saldo: 120.5}
	//conta2 := contas.ContaCorrente{Titular: "Anderson", NumeroAgencia: 1453, NumeroConta: 27979, Saldo: 120.5}
	//conta := contas.ContaCorrente{"Anderson", 1453, 27979, 120.5}
	conta2 := contas.ContaCorrente{"Anderson", 1453, 27979, 120.5}
	fmt.Println(conta == conta2)
	fmt.Println(&conta == &conta2)

	conta.Sacar(10.0)
	conta.Transferir(10, &conta2)

	var conta3 *contas.ContaCorrente
	conta3 = new(contas.ContaCorrente)
	conta3.Titular = "José"
	fmt.Println(conta3)
	fmt.Println(*conta3)

	var conta4 *contas.ContaCorrente
	conta4 = new(contas.ContaCorrente)
	conta4.Titular = "José"
	fmt.Println(conta3 == conta4)
	fmt.Println(*conta3 == *conta4)
}
