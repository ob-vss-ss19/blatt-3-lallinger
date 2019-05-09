package tree

import (
	"sort"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/ob-vss-ss19/blatt-3-lallinger/messages"
)

type Find struct {
	Key    int
	Remove bool
	caller *actor.PID
}
type Add struct {
	Key   int
	Value string
}
type Delete struct {
	currentNode *actor.PID
}

type keyValuePair struct {
	Key   int
	Value string
}
type Traverse struct {
	values         []keyValuePair
	remainingNodes []*actor.PID
	caller         *actor.PID
	start          *actor.PID
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
				} else {
					// return value
					context.Send(msg.caller, &messages.Response{Value: tmp})
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
			if state.Values == nil {
				state.Values = make(map[int]string)
			}
			state.Values[msg.Key] = msg.Value

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
		if msg.start != nil {
			// set root node as start node for traverse
			msg.values = make([]keyValuePair, 0)
			msg.remainingNodes = make([]*actor.PID, 1)
			tmp := msg.start
			msg.start = nil
			if len(state.Values) != 0 && state.LeftNode == nil && state.RightNode == nil {
				// if root is leaf create slices and set start to nil
				context.Send(tmp, msg)
				return
			}

			// if root is node create slices, set start to nil, add right node to remaining and forward
			msg.remainingNodes = append(msg.remainingNodes, state.RightNode)
			context.Send(state.LeftNode, msg)
			return
		}

		if len(msg.remainingNodes) != 0 && len(state.Values) == 0 && state.LeftNode != nil && state.RightNode != nil {
			// node is not leaf
			// while remaining nodes add right node to remaining and sends to left node
			msg.remainingNodes = append(msg.remainingNodes, state.RightNode)
			context.Send(state.LeftNode, msg)
		}

		if len(msg.remainingNodes) != 0 && state.LeftNode == nil && state.RightNode == nil {
			// leaf with remaining nodes to traverse
			for key := range sortKeys(state.Values) {
				msg.values = append(msg.values, keyValuePair{key, state.Values[key]})
			}
			next := msg.remainingNodes[len(msg.remainingNodes)-1]
			msg.remainingNodes = msg.remainingNodes[:len(msg.remainingNodes)-2]
			context.Send(next, msg)
		}

		if len(msg.remainingNodes) == 0 && state.LeftNode == nil && state.RightNode == nil {
			// leaf with no remaining nodes to traverse
			for key := range sortKeys(state.Values) {
				msg.values = append(msg.values, keyValuePair{key, state.Values[key]})
			}
			context.Send(msg.caller, msg)
		}

	case *Delete:
		context.Send(state.LeftNode, &Delete{currentNode: state.LeftNode})
		context.Send(state.RightNode, &Delete{currentNode: state.RightNode})
		context.Stop(msg.currentNode)
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
