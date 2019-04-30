package main

import (
	"fmt"

	console "github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
)

func main() {
	context := actor.EmptyRootContext
	props := actor.PropsFromProducer(func() actor.Actor {
		return &NodeActor{Token: "a", Id: 1}
	})
	pid := context.Spawn(props)
	context.Send(pid, &add{value: "hallo", Key: 4, id: 1, token: "a"})
	console.ReadLine() // nolint:errcheck
}
