package main

import (
	"fmt"

	console "github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
)

type hello struct{ who string }
type goodbye struct{ until string }
type helloActor struct{}

func (state *helloActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *hello:
		fmt.Printf("Hello %v\n", msg.who)
	case *goodbye:
		fmt.Printf("ok cu %v\n", msg.until)
	}
}

func main() {
	context := actor.EmptyRootContext
	props := actor.PropsFromProducer(func() actor.Actor {
		return &helloActor{}
	})
	pid := context.Spawn(props)
	context.Send(pid, &hello{who: "Roger"})
	context.Send(pid, &goodbye{until: "Tomorrow"})
	console.ReadLine() // nolint:errcheck
}
