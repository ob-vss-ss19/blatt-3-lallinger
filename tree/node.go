package tree

import (
	"fmt"
	"sort"

	console "github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
)

type Find struct {
	Key    int
	Remove bool
	Id     int
	Token  string
}
type Add struct {
	Key   int
	Value string
	Id    int
	Token string
}
type Delete struct {
	Id    int
	Token string
}
type Traverse struct {
	Pairs  map[int]string
	Caller *NodeActor
	Id     int
	Token  string
}

type NodeActor struct {
	Token      string
	Id         int
	LeafSize   int
	LeftNode   *NodeActor
	RightNode  *NodeActor
	LeftMaxKey int
	Values     map[int]string
}

func (state *NodeActor) Receive(context actor.Context) {
	if context.Message().Token != state.Token || context.Message().Id != state.Id {
		// send error
	}

	switch msg := context.Message().(type) {
	case *Find:
		if len(state.Values) == 0 && state.LeftNode != nil && state.RightNode != nil {
			// not a leaf search next nodes
			if msg.Key <= state.LeftMaxKey {
				// send left
			} else {
				// send right
			}
		} else if state.LeftNode == nil && state.RightNode == nil {
			// node is leaf search for key
			tmp := state.Values[msg.Key]
			if tmp != "" {
				if msg.Remove {
					delete(state.Values, msg.Key)
				} else {
					// return value
				}
			} else {
				// return error
			}
		} else {
			// undefined send error
		}
	case *Add:
		if len(state.Values) < state.LeafSize && state.LeftNode == nil && state.RightNode == nil {
			// add key to leaf
			state.Values[msg.Key] = msg.Value
			fmt.Printf("added %v\n", msg.Value)
		} else if len(state.Values) == 0 && state.LeftNode != nil && state.RightNode != nil {
			// not a leaf
			if msg.Key <= state.LeftMaxKey {
				// add left
			} else {
				// add right
			}
		} else if len(state.Values) == state.LeafSize && state.LeftNode == nil && state.RightNode == nil {
			// leaf full create new leafs

			// send values to leafs
			state.Values[msg.Key] = msg.Value
			keys := sortKeys(state.Values)
			state.LeftMaxKey = keys[(len(keys)/2)-1]

			for key := range keys {
				if key <= state.LeftMaxKey {
					// add half left
				} else {
					// add half right
				}
			}

			state.Values = nil
		} else {
			// undefined send error
		}

	case *Traverse:

	case *Delete:
		if state.LeftNode != nil && state.RightNode != nil {
			// not a leaf forward delete
		} else {
			// delete leaf
		}
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
