// Etapa 2: Produtor Finito
// ●​ Modifique a versão anterior para que o produtor gere apenas 10.000 valores aleatórios.
// ●​ Quando terminar, feche o canal e encerre o programa.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(ch chan int, max int) {
	producerTime := 100 * time.Millisecond
	for i := 0; i < max; i++ {
		value := rand.Intn(101)
		fmt.Printf("Valor lido: %d.\n", value)
		ch <- value
		time.Sleep(producerTime)
	}
	fmt.Println("Todos os valores foram lidos.")
	close(ch)
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
	fmt.Println("ETAPA 2: PRODUTOR FINITO")

	ch := make(chan int)
	done := make(chan bool)

	threshold := 50
	max := 10
	fmt.Printf("Limite: %d. Quantidade máxima de valores: %d.\n", threshold, max)

	go producer(ch, max)
	go consumer(ch, done, threshold)

	<-done
}
