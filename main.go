package main

import (
	"fmt"
	"runtime"
)

/*
func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("GOROUTINE\t", runtime.NumGoroutine())
}

// hands-on exercise #1

var wg sync.WaitGroup

func main() {
	fmt.Println("First Goroutine")
	fmt.Println("GOROUTINE\t", runtime.NumGoroutine())
	wg.Add(2)
	go Add(4, 5)
	go Incrementor()
	fmt.Println("GOROUTINE\t", runtime.NumGoroutine())
	wg.Wait()
}

func Add(a int, b int) {
	fmt.Println(a + b)
	wg.Done()

}

func Incrementor() {
	c := 0
	for i := 0; i < 10; i++ {
		c++
		fmt.Println(c)
	}
	wg.Done()

}

// hands-on exercise #2

//type struct
type Person struct {
	First string
	Last  string
}

// method set
func (P *Person) Speak() string {
	return "Hello, I am " + P.First + " " + P.Last
}

//interface
type human interface {
	Speak() string
}

// func
func SaySomething(h human) {
	fmt.Println(h.Speak())
}

// func main
func main() {
	Person1 := Person{First: "John", Last: "Doe"}

	//pointer reciever
	SaySomething(&Person1)

	//value receiver, will not work
	//SaySomething(Person1)
}

// hands-on exercise #3

func main() {
	var wg sync.WaitGroup

	incremeter := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := incremeter
			runtime.Gosched()
			v++
			incremeter = v
			fmt.Println(incremeter)
			wg.Done()
			}()
		}
		wg.Wait()
		fmt.Println("End value:", incremeter)
	}

	// hands-on exercise #4
	// mutex

	func main() {
		var wg sync.WaitGroup
		var mu sync.Mutex

		incremeter := 0
		gs := 100
		wg.Add(gs)

		for i := 0; i < gs; i++ {
			go func() {
				mu.Lock()
				v := incremeter
				v++
				incremeter = v
				fmt.Println(incremeter)
			mu.Unlock()
			wg.Done()
			}()
		}
		wg.Wait()
		fmt.Println("End value:", incremeter)
	}

	// hands-on exercise #4
	// Atomic

	func main() {
		var wg sync.WaitGroup
		var incremeter int64

		gs := 100
		wg.Add(gs)

		for i := 0; i < gs; i++ {
			go func() {
				atomic.AddInt64(&incremeter, 1)
				fmt.Println(atomic.LoadInt64(&incremeter))
				wg.Done()
				}()
			}
			wg.Wait()
			fmt.Println("End value:", incremeter)
		}
*/

// hands-on exercise #5

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
}
