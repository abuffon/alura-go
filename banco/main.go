package main

import (
	"fmt"

	"github.com/abuffon/alura-go/banco/clientes"
	"github.com/abuffon/alura-go/banco/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	cliente := clientes.Titular{
		Nome:      "Anderson",
		CPF:       "036",
		Profissao: "Analista"}

	conta := contas.ContaCorrente{
		Titular:       cliente,
		NumeroAgencia: 111,
		NumeroConta:   123}

	conta.Depositar(100)
	fmt.Println(conta)
	PagarBoleto(&conta, 60)
	fmt.Println(conta.ObterSaldo())

	contaP := contas.ContaPoupanca{
		Titular:       cliente,
		NumeroAgencia: 111,
		NumeroConta:   123,
		Operacao:      13}
	contaP.Depositar(200)

	fmt.Println(contaP)
	PagarBoleto(&contaP, 60)
	fmt.Println(contaP.ObterSaldo())
	/*conta := contas.ContaCorrente{Titular: "Anderson", NumeroAgencia: 1453, NumeroConta: 27979, Saldo: 120.5}
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
	fmt.Println(*conta3 == *conta4)*/
}
