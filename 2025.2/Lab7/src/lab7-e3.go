// Etapa 3: Múltiplos Sensores (Produtores)
// ●​ Agora considere que existem dois sensores (duas goroutines produtoras).
// ●​ Cada produtor gera uma quantidade aleatória de valores.
// ●​ O consumidor central (único) recebe valores de ambos continua imprimindo apenas os que estão acima do limite.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(ch chan int, done chan bool, max int) {
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

func consumer(ch chan int, done chan bool, threshold int) {
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
	fmt.Println("ETAPA 3: MÚLTIPLOS PRODUTORES")

	ch := make(chan int)
	producersDone := make(chan bool)
	consumersDone := make(chan bool)

	threshold := 50
	max := 10
	fmt.Printf("Limite: %d.\n", threshold)

	go producer(ch, producersDone, rand.Intn(max+1))
	go producer(ch, producersDone, rand.Intn(max+1))
	go consumer(ch, consumersDone, threshold)

	<-producersDone
	<-producersDone
	fmt.Println("Todos os valores foram lidos.")
	close(ch)
	<-consumersDone
}
