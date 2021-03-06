// typeAssertion.go
package main

import "fmt"

type Hello struct{}

func (h Hello) String() string {
	return "Hello"
}

type World struct{}

func (w *World) String() string {
	return "Pinas!"
}

type Message struct {
	X fmt.Stringer
	Y fmt.Stringer
}

func (v Message) IsGreeting() (ok bool) {
	if _, ok = v.X.(*Hello); !ok {
		_, ok = v.Y.(*Hello)
	}
	return
}

func main() {
	m := &Message{}
	fmt.Println(m.IsGreeting())
	m.X = new(Hello)
	fmt.Println(m.IsGreeting())
	m.Y = new(World)
	fmt.Println(m.IsGreeting())
	m.Y = m.X
	fmt.Println(m.IsGreeting())
	m = &Message{X: new(World), Y: new(Hello)}
	fmt.Println(m.IsGreeting())
	m.X, m.Y = m.Y, m.X
	fmt.Println(m.IsGreeting())
}
