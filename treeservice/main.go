package main

import (
	"flag"
	"github.com/ob-vss-ss19/blatt-3-lallinger/messages"
	"github.com/ob-vss-ss19/blatt-3-lallinger/tree"
	"sync"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/log"
	"github.com/AsynkronIT/protoactor-go/remote"
	_ "github.com/ob-vss-ss19/blatt-3-lallinger/tree"
)

type ServiceActor struct{}

var context = actor.EmptyRootContext
var trees = make(map[int]map[string]*actor.PID)

func (state *ServiceActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *messages.Request:

		switch msg.Type {
		case messages.Usage_CREATE:

			id := nextId()
			token := newToken()

			props := actor.PropsFromProducer(func() actor.Actor {
				return &tree.NodeActor{Id: id, LeafSize: int(msg.Id)}
			})
			pid := context.Spawn(props)
			trees[id][token] = pid
			context.Respond(&messages.Response{Key: int32(id), Value: token})

		case messages.Usage_ADD:
		case messages.Usage_FIND:
		case messages.Usage_REMOVE:
		case messages.Usage_TRAVERSE:
		case messages.Usage_DELETE:
		default:

		}

		context.Respond(&messages.Response{
			Key: 1,
		})
	default: // just for linter
	}
}

func NewMyActor() actor.Actor {
	log.Message("Hello-Actor is up and running")
	return &ServiceActor{}
}

// nolint:gochecknoglobals
var flagBind = flag.String("bind", "localhost:8091", "Bind to address")

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()

	flag.Parse()
	remote.Start(*flagBind)

	remote.Register("hello", actor.PropsFromProducer(NewMyActor))

	/*
		props := actor.PropsFromProducer(func() actor.Actor {
			return &tree.NodeActor{Id: 1, LeafSize: 2} // nolint:errcheck
		})
		pid := context.Spawn(props)
		context.Send(pid, &tree.Add{Value: "hallo", Key: 4})
		context.Send(pid, &tree.Find{Key: 4})
		console.ReadLine() // nolint:errcheck*/
}

func nextId() int {
	return 1
}

func newToken() string {
	return "a"
}
