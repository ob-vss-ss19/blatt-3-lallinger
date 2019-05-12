package tree

import (
	"fmt"
	"sort"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/ob-vss-ss19/blatt-3-lallinger/messages"
)

type Find struct {
	Key    int
	Remove bool
	Caller *actor.PID
}
type Add struct {
	Key   int
	Value string
}
type Delete struct {
	CurrentNode *actor.PID
}

type KeyValuePair struct {
	Key   int
	Value string
}
type Traverse struct {
	Values         []KeyValuePair
	RemainingNodes []*actor.PID
	Caller         *actor.PID
	Start          *actor.PID
}

type NodeActor struct {
	LeafSize   int
	LeftNode   *actor.PID
	RightNode  *actor.PID
	LeftMaxKey int
	Values     map[int]string
}

func (state *NodeActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *Find:
		if len(state.Values) == 0 && state.LeftNode != nil && state.RightNode != nil {
			// not a leaf search next nodes
			if msg.Key <= state.LeftMaxKey {
				// send left
				context.Send(state.LeftNode, msg)
			} else {
				// send right
				context.Send(state.RightNode, msg)
			}
		} else if state.LeftNode == nil && state.RightNode == nil {
			// node is leaf search for key
			tmp := state.Values[msg.Key]
			if tmp != "" {
				if msg.Remove {
					delete(state.Values, msg.Key)
					fmt.Println("deleted key")
				} else {
					// return value
					context.Send(msg.Caller, &messages.Response{Value: tmp, Type: messages.FIND, Key: int32(msg.Key)})
					fmt.Println("found key")
				}
			} else {
				// return error key not found
				context.Send(msg.Caller, &messages.Error{Message: "Key not found"})
			}
		}
	case *Add:
		if (len(state.Values) < state.LeafSize || state.Values[msg.Key] != "") && state.LeftNode == nil && state.RightNode == nil {
			// add key to leaf
			if state.Values == nil {
				state.Values = make(map[int]string)
			}
			state.Values[msg.Key] = msg.Value
			fmt.Printf("added key: %d\n", msg.Key)

		} else if len(state.Values) == 0 && state.LeftNode != nil && state.RightNode != nil {
			// not a leaf
			if msg.Key <= state.LeftMaxKey {
				// add left
				context.Send(state.LeftNode, msg)
			} else {
				// add right
				context.Send(state.RightNode, msg)
			}
		} else if len(state.Values) == state.LeafSize && state.LeftNode == nil && state.RightNode == nil {
			// leaf full create new leafs
			fmt.Println("created new leafs")
			props := actor.PropsFromProducer(func() actor.Actor {
				return &NodeActor{LeafSize: int(state.LeafSize)}
			})
			state.LeftNode = context.Spawn(props)

			state.RightNode = context.Spawn(props)

			// send values to leafs
			state.Values[msg.Key] = msg.Value
			keys := sortKeys(state.Values)
			state.LeftMaxKey = keys[(len(keys)/2)-1]
			fmt.Printf("left max key %d\n", state.LeftMaxKey)
			for _, key := range keys {
				if key <= state.LeftMaxKey {
					// add half left
					fmt.Printf("send %d left\n", key)
					context.Send(state.LeftNode, &Add{Key: key, Value: state.Values[key]})
					delete(state.Values, key)
				} else {
					// add half right
					fmt.Printf("send %d right\n", key)
					context.Send(state.RightNode, &Add{Key: key, Value: state.Values[key]})
					delete(state.Values, key)
				}
			}
		}
	case *Traverse:
		if msg.Start != nil {
			// set root node as start node for traverse
			msg.Values = make([]KeyValuePair, 0)
			msg.RemainingNodes = make([]*actor.PID, 0)
			tmp := msg.Start
			msg.Start = nil
			if state.LeftNode == nil && state.RightNode == nil {
				// if root is leaf create slices and set start to nil
				fmt.Println("send to start")
				context.Send(tmp, msg)
				return
			}

			// if root is node create slices, set start to nil, add right node to remaining and forward
			msg.RemainingNodes = append(msg.RemainingNodes, state.RightNode)
			fmt.Println("send to left node from start")
			context.Send(state.LeftNode, msg)
			return
		} else if state.LeftNode != nil && state.RightNode != nil {
			// node is not leaf
			// while remaining nodes add right node to remaining and send to left node
			fmt.Printf("add right node send to left")
			msg.RemainingNodes = append(msg.RemainingNodes, state.RightNode)
			context.Send(state.LeftNode, msg)
		} else if len(msg.RemainingNodes) != 0 && state.LeftNode == nil && state.RightNode == nil {
			// leaf with remaining nodes to traverse
			for _, key := range sortKeys(state.Values) {
				fmt.Printf("appending %d\n", key)
				msg.Values = append(msg.Values, KeyValuePair{key, state.Values[key]})
			}
			next := msg.RemainingNodes[len(msg.RemainingNodes)-1]
			msg.RemainingNodes = msg.RemainingNodes[:len(msg.RemainingNodes)-1]
			fmt.Println("send to next node from leaf")
			context.Send(next, msg)
		} else if len(msg.RemainingNodes) == 0 && state.LeftNode == nil && state.RightNode == nil {
			// leaf with no remaining nodes to traverse
			for _, key := range sortKeys(state.Values) {
				fmt.Printf("appending %d in last leaf\n", key)
				msg.Values = append(msg.Values, KeyValuePair{key, state.Values[key]})
			}

			fmt.Println("send to caller")
			for _, pair := range msg.Values {
				fmt.Printf("sending key %d\n", pair.Key)
				context.Send(msg.Caller, &messages.Response{Value: pair.Value, Key: int32(pair.Key), Type: messages.TRAVERSE})
			}
			context.Send(msg.Caller, &messages.Response{Type: messages.SUCCESS})
		} else {
			fmt.Printf("error in traverse\n")
			context.Send(msg.Caller, &messages.Error{"Error while traversing"})
		}

	case *Delete:
		fmt.Println("stopping node")
		if state.LeftNode != nil {
			context.Send(state.LeftNode, &Delete{CurrentNode: state.LeftNode})
		}

		if state.RightNode != nil {
			context.Send(state.RightNode, &Delete{CurrentNode: state.RightNode})
		}

		context.Stop(msg.CurrentNode) // nolint
		fmt.Println("still running?")
	}
}

func sortKeys(Values map[int]string) []int {
	var keys []int
	for k := range Values {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	return keys
}
