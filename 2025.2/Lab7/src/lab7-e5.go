// Etapa 5: Vários Consumidores
// ●​ Evolua para ter múltiplas goroutines consumidoras, cada uma processando valores recebidos do canal.
// ●​ Cada consumidor imprime valores com uma identificação própria (ex.: Consumidor 1 recebeu 87).

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(ch chan<- int, done chan<- bool, max int, id int) {
	producerTime := 100 * time.Millisecond
	fmt.Printf("Prod.: %d | Quantidade máxima de valores: %d.\n", id, max)
	for i := 0; i < max; i++ {
		value := rand.Intn(101)
		fmt.Printf("Prod.: %d | Valor lido: %d.\n", id, value)
		ch <- value
		time.Sleep(producerTime)
	}
	done <- true
}

func consumer(ch <-chan int, done chan<- bool, threshold int, id int) {
	consumerTime := 250 * time.Millisecond
	for value := range ch {
		if value > threshold {
			fmt.Printf("Cons.: %d | Valor recebido acima do limite: %d.\n", id, value)
		} else {
			fmt.Printf("Cons.: %d | Valor recebido: %d.\n", id, value)
		}
		time.Sleep(consumerTime)
	}
	done <- true
}

func main() {
	fmt.Println("ETAPA 5: MÚLTIPLOS PRODUTOS E CONSUMIDORES")

	buffer := 3
	ch := make(chan int, buffer)
	producersDone := make(chan bool)
	consumersDone := make(chan bool)

	threshold := 50
	max := 20
	fmt.Printf("Limite: %d.\n", threshold)
	fmt.Printf("Quantidade máxima de valores: %d.\n", max)
	fmt.Printf("Tamanho do buffer: %d.\n", buffer)

	producers := 5
	consumers := 2

	for p := 1; p <= producers; p++ {
		go producer(ch, producersDone, max, p)
	}

	for c := 1; c <= consumers; c++ {
		go consumer(ch, consumersDone, threshold, c)
	}

	for p := 0; p < producers; p++ {
		<-producersDone
	}
	fmt.Println("Todos os valores foram lidos.")
	close(ch)

	for c := 0; c < consumers; c++ {
		<-consumersDone
	}
	fmt.Println("Todos os valores foram analisados.")
}
