package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice_tentativas := []int{}
	jogadas := []int{}
	for i := 1; i <= 100; i++ {
		jogadas = append(jogadas, i)
	}
	for true {
		fmt.Println("Bem vindo ao jogo da adivinhação")
		min := 1
		max := 100
		var tentativas int = 0
		x := rand.Intn((max - min) + min)
		var n int
		fmt.Println("--------------------------------")
		for n != x {
			fmt.Println("Insira um número entre 1 e 100")
			tentativas++
			fmt.Scan(&n)
			if n < x {
				fmt.Println("O número é maior que ", n)
			} else if n > x {
				fmt.Println("O número é menor que ", n)
			} else {
				fmt.Println("--------------------------------")
				fmt.Println("Parabéns, você acertou!")
				fmt.Println("Você utilizou ", tentativas, " tentativas")
				slice_tentativas = append(slice_tentativas, tentativas)
			}
		}
		var jogar_denovo string
		fmt.Println("Você deseja jogar novamente? (sim ou não)")
		fmt.Scan(&jogar_denovo)
		if jogar_denovo != "sim" && jogar_denovo != "s" {
			fmt.Println("--------------------------------")
			num := 0
			for num < len(slice_tentativas) {
				fmt.Printf("Você utilizou %d tentativas na %dª jogada\n", slice_tentativas[num], jogadas[num])
				num++
			}
			break
		}
	}
}
