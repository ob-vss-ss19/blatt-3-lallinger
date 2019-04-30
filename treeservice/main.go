package main

import (
	console "github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/ob-vss-ss19/blatt-3-lallinger/tree"
)

func main() {
	context := actor.EmptyRootContext
	props := actor.PropsFromProducer(func() actor.Actor {
		return &NodeActor{Token: "a", Id: 1}
	})
	pid := context.Spawn(props)
	context.Send(pid, &Add{value: "hallo", Key: 4, id: 1, token: "a"})
	console.ReadLine() // nolint:errcheck
}
