// Etapa 1: Produtor e Consumidor Simples
// ●​ Implemente um programa com uma goroutine produtora que gera valores aleatórios (simulando leituras de um sensor) entre 0 e 100.
// ●​ Uma goroutine consumidora lê do canal e imprime apenas os valores acima de um limite pré-definido (exemplo: maior que 50).
// ●​ Tanto o produtor quanto o consumidor rodam em loop infinito.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(ch chan int) {
	producerTime := 100 * time.Millisecond
	for {
		value := rand.Intn(101)
		fmt.Printf("Valor lido: %d.\n", value)
		ch <- value
		time.Sleep(producerTime)
	}
}

func consumer(ch chan int, threshold int) {
	consumerTime := 250 * time.Millisecond
	for {
		value := <-ch
		if value > threshold {
			fmt.Printf("Valor acima do limite: %d.\n", value)
		}
		time.Sleep(consumerTime)
	}
}

func main() {
	fmt.Println("ETAPA 1: PRODUTOR-CONSUMIDOR SIMPLES")

	ch := make(chan int)

	threshold := 50
	fmt.Printf("Limite: %d.\n", threshold)

	go producer(ch)
	go consumer(ch, threshold)
	select {}
}
