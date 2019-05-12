package main

import (
	"flag"
	"fmt"
	"github.com/ob-vss-ss19/blatt-3-lallinger/messages"
	"github.com/ob-vss-ss19/blatt-3-lallinger/tree"
	"sync"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
)

type ServiceActor struct{}

var context = actor.EmptyRootContext
var trees = make(map[int32]map[string]*actor.PID)

func (state *ServiceActor) Receive(context actor.Context) {

	switch msg := context.Message().(type) {
	case *messages.Request:
		switch msg.Type {
		case messages.CREATE:
			id := nextId()
			token := newToken()

			props := actor.PropsFromProducer(func() actor.Actor {
				return &tree.NodeActor{LeafSize: int(msg.Id)}
			})
			pid := context.Spawn(props)
			trees[id] = make(map[string]*actor.PID)
			trees[id][token] = pid
			context.Respond(&messages.Response{Key: int32(id), Value: token, Type: messages.CREATE})

		case messages.ADD:
			pid := getPID(msg.Id, msg.Token)
			if pid == nil {
				invalidAcess(context.Sender())
				return
			}
			context.Send(pid, &tree.Add{Key: int(msg.Key), Value: msg.Value})
			context.Respond(&messages.Response{Type: messages.SUCCESS})
		case messages.FIND:
			pid := getPID(msg.Id, msg.Token)
			if pid == nil {
				invalidAcess(context.Sender())
				return
			}
			context.Send(pid, &tree.Find{Key: int(msg.Key), Caller: context.Sender()})
		case messages.REMOVE:
			pid := getPID(msg.Id, msg.Token)
			if pid == nil {
				invalidAcess(context.Sender())
				return
			}
			context.Send(pid, &tree.Find{Key: int(msg.Key), Remove: true})
			context.Respond(&messages.Response{Type: messages.SUCCESS})
		case messages.TRAVERSE:
			pid := getPID(msg.Id, msg.Token)
			if pid == nil {
				invalidAcess(context.Sender())
				return
			}
			context.Send(pid, &tree.Traverse{Caller: context.Sender(), Start: pid})
		case messages.DELETE:
			pid := getPID(msg.Id, msg.Token)
			if pid == nil {
				invalidAcess(context.Sender())
				return
			}
			context.Send(pid, &tree.Delete{CurrentNode: pid})
			delete(trees, msg.Id)
			context.Respond(&messages.Response{Type: messages.SUCCESS})
		}
	}

}

func getPID(id int32, token string) *actor.PID {
	tmp := trees[id][token]
	if tmp != nil {
		return tmp
	}
	return nil
}

func invalidAcess(pid *actor.PID) {
	context.Send(pid, &messages.Error{Message: "Id and token do not match or tree does not exist"})
}

func NewMyActor() actor.Actor {
	fmt.Println("Service-Actor is up and running")
	return &ServiceActor{}
}

// nolint:gochecknoglobals
var flagBind = flag.String("bind", "localhost:8093", "Bind to address")

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()

	flag.Parse()
	remote.Start(*flagBind)
	remote.Register("treeService", actor.PropsFromProducer(NewMyActor))

}

func nextId() int32 {
	return int32(len(trees))
}

func newToken() string {
	//b := make([]byte, 4)
	//rand.Read(b)
	return "a"
	//return fmt.Sprintf("%x", b)
}
