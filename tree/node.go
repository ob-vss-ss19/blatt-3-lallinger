package tree

import (
	"fmt"

	console "github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
)

type find struct {
	key   int32
	token string
}
type add struct {
	key   int32
	value string
	token string
}
type remove struct {
	key   int32
	token string
}
type delete struct {
	token string
}
type traverse struct {
	token string
}

type NodeActor struct {
	Token      string
	LeafSize   int
	LeftNode   *NodeActor
	RightNode  *NodeActor
	LeftMaxKey int32
	Values     map[int32]string
}

func (state *NodeActor) Receive(context actor.Context) {
	if context.Message().token != state.Token {
		// send error
	}

	switch msg := context.Message().(type) {
	case *find:
		if len(state.Values) == 0 && state.LeftNode != nil && state.RightNode != nil {
			// not a leaf search next nodes
			if state.LeftMaxKey <= msg.key {
				// send left
			} else {
				// send right
			}
		} else if state.LeftNode == nil && state.RightNode == nil {
			// node is leaf search for key
			tmp := state.Values[msg.key]
			if tmp != "" {
				// return value
			} else {
				// return error
			}
		} else {
			// undefined send error
		}
	case *add:
		if state.LeafSize < len(state.Values) && state.LeftNode == nil && state.RightNode == nil {
			// add key to leaf
			state.Values[msg.key] = msg.value
		} else if len(state.Values) == 0 && state.LeftNode != nil && state.RightNode != nil {
			// not a leaf
		} else if {
			// leaf full create new leafs
		} else {
			// undefined send error
		}

	case *traverse:

	case *remove:
		fmt.Printf("ok cu %v\n", msg.until)
	case *delete:
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
