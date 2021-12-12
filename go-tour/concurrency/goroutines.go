package concurrency

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func GoRoutineHello() {
	go say("world")
	say("hello")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func teste1(c chan int) {
	c <- 1
}

//Channels aren't like files; you don't usually need to close them. Closing is only necessary
//when the receiver must be told there are no more values coming, such as to terminate a range loop.
//Para fechar um channel basta usar close(c)
func teste2(c chan int) {
	time.Sleep(1000 * time.Millisecond)
	c <- 2
}

func GoRoutineFunction() {
	fmt.Println("Go routine...")
	s := []int{7, 2, 8, -9, 4, 0}

	//O segundo argumento é o buffer do channel
	c := make(chan int, 3)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	// x, ok <- c //permite validar através do "ok" se o channel possui mais valores a serem lidos

	fmt.Println(x, y, x+y)
	go teste2(c)
	go teste1(c)
	go teste1(c)
	//a ordem do channel acontece de acordo em que os eventos forem sendo finalizados
	y1 := <-c
	z1 := <-c
	x1 := <-c
	fmt.Println(x1, y1, z1)

	//É possível fazer um for com range em um channel
	//for i := range c {
	//	fmt.Println(i)
	//}
}

//https://go.dev/tour/concurrency/5
// The select statement lets a goroutine wait on multiple communication operations.
//A select blocks until one of its cases can run, then it executes that case.
//It chooses one at random if multiple are ready.
func SelectStatement() {
	c := make(chan int)
	quit := make(chan int)
	//fibonacci(c, quit)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		//irá executar este bloco até o count i do for da chamada ser finalizado
		//após ter sido finalizado, este case não poderá mais ser acessado
		case c <- x:
			x, y = y, x+y
		//e este passa a ser possível de ser acessado, o que ocasiona o quit
		case <-quit:
			fmt.Println("quit")
			return
			// É possível usar default no switch no qual irá ser executado quando nenhum
			// outro caso está pronto
		}
	}
}
