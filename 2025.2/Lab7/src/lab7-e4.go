// Etapa 4: Canal Unidirecional e Bufferizado
// ●​ Refaça o programa anterior:
//     ○​ Defina o canal como unidirecional (chan<- para produtores e <-chan para consumidores).
//     ○​ Use um canal bufferizado (ex.: buffer de 100 elementos) para reduzir bloqueios entre produtores e consumidor.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(ch chan<- int, done chan<- bool, max int) {
	producerTime := 100 * time.Millisecond
	fmt.Printf("Quantidade máxima de valores: %d.\n", max)
	for i := 0; i < max; i++ {
		value := rand.Intn(101)
		fmt.Printf("Valor lido: %d.\n", value)
		ch <- value
		time.Sleep(producerTime)
	}
	done <- true
}

func consumer(ch <-chan int, done chan<- bool, threshold int) {
	consumerTime := 250 * time.Millisecond
	for value := range ch {
		if value > threshold {
			fmt.Printf("Valor acima do limite: %d.\n", value)
		}
		time.Sleep(consumerTime)
	}
	fmt.Println("Todos os valores foram analisados.")
	done <- true
}

func main() {
	fmt.Println("ETAPA 4: CANAIS UNIDIRECIONAIS E BUFFERIZADOS")

	buffer := 3
	ch := make(chan int, buffer)
	producersDone := make(chan bool)
	consumersDone := make(chan bool)

	threshold := 50
	max := 20
	fmt.Printf("Limite: %d. Tamanho do buffer: %d.\n", threshold, buffer)

	go producer(ch, producersDone, rand.Intn(max+1))
	go producer(ch, producersDone, rand.Intn(max+1))
	go consumer(ch, consumersDone, threshold)

	<-producersDone
	<-producersDone
	fmt.Println("Todos os valores foram lidos.")
	close(ch)
	<-consumersDone
}
