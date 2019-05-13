package tree

import (
	cr "crypto/rand"
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/ob-vss-ss19/blatt-3-lallinger/messages"
	mr "math/rand"
	"sync"
	"testing"
	"time"
)

type TestActor struct {
	t       *testing.T
	wg      *sync.WaitGroup
	indices []int
}

var values = make([]KeyValuePair, 0) //[]KeyValuePair{{-6, "minus sechs"}, {-3, "minus 3"}, {0, "null"}, {1, "eins"}, {3, "drei"}, {6, "sechs"}}

func (state *TestActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *messages.Traverse:
		i := 0
		for _, k := range state.indices {
			if msg.Values[i].Value != values[k].Value || int(msg.Values[i].Key) != values[k].Key {
				fmt.Printf("should be: %d %s but is %d %s\n", msg.Values[i].Key, msg.Values[i].Value, values[k].Key, values[k].Value)
				state.t.Error()
			}
			i++
		}
		state.wg.Done()
	case *messages.Response:
		switch msg.Type {
		case messages.FIND:
			if int(msg.Key) != values[state.indices[0]].Key || msg.Value != values[state.indices[0]].Value {
				fmt.Printf("should be: %d %s but is %d %s\n", msg.Key, msg.Value, values[state.indices[0]].Key, values[state.indices[0]].Value)
				state.t.Error()
			}
			state.wg.Done()
		}

	}
}

func createRandValues() {

	values = make([]KeyValuePair, 0)

	pairs := make(map[int]string)

	for i := 0; i < 5; i++ {
		pairs[int(mr.Int31n(1000))] = newToken()
	}

	for _, v := range sortKeys(pairs) {
		values = append(values, KeyValuePair{Key: v, Value: pairs[v]})
	}

}

func TestAdd(t *testing.T) {
	createRandValues()
	context := actor.EmptyRootContext
	props := actor.PropsFromProducer(func() actor.Actor {
		return &NodeActor{LeafSize: 1}
	})
	tree := context.Spawn(props)
	var wg sync.WaitGroup

	indices := make(map[int]string)
	for range values {
		tmp := mr.Intn(len(values))
		indices[tmp] = ""
		context.Send(tree, &Add{Key: values[tmp].Key, Value: values[tmp].Value})
	}
	time.Sleep(1 * time.Second)

	props = actor.PropsFromProducer(func() actor.Actor {
		wg.Add(1)
		return &TestActor{t, &wg, sortKeys(indices)}
	})
	testAct := context.Spawn(props)
	context.Send(tree, &Traverse{Caller: testAct, Start: tree})
	time.Sleep(1 * time.Second)
	wg.Wait()
}

func TestDelete(t *testing.T) {
	createRandValues()

	context := actor.EmptyRootContext
	props := actor.PropsFromProducer(func() actor.Actor {
		return &NodeActor{LeafSize: 2}
	})
	tree := context.Spawn(props)
	var wg sync.WaitGroup

	indices := make(map[int]string)
	for range values {
		tmp := mr.Intn(len(values))
		indices[tmp] = ""
		context.Send(tree, &Add{Key: values[tmp].Key, Value: values[tmp].Value})
	}
	time.Sleep(2 * time.Second)

	for k := range indices {
		if mr.Int31n(100) < 50 {
			delete(indices, k)
			context.Send(tree, &Find{Key: values[k].Key, Remove: true, Caller: tree})
		}
	}

	time.Sleep(2 * time.Second)

	props = actor.PropsFromProducer(func() actor.Actor {
		wg.Add(1)
		return &TestActor{t, &wg, sortKeys(indices)}
	})
	testAct := context.Spawn(props)
	context.Send(tree, &Traverse{Caller: testAct, Start: tree})
	wg.Wait()
	time.Sleep(1 * time.Second)
}

func TestFind(t *testing.T) {
	createRandValues()

	context := actor.EmptyRootContext
	props := actor.PropsFromProducer(func() actor.Actor {
		return &NodeActor{LeafSize: 2}
	})
	tree := context.Spawn(props)
	var wg sync.WaitGroup

	indices := make(map[int]string)
	for range values {
		tmp := mr.Intn(len(values))
		indices[tmp] = ""
		context.Send(tree, &Add{Key: values[tmp].Key, Value: values[tmp].Value})
	}
	time.Sleep(2 * time.Second)

	for k := range indices {
		props = actor.PropsFromProducer(func() actor.Actor {
			wg.Add(1)
			return &TestActor{t, &wg, []int{k}}
		})
		testAct := context.Spawn(props)
		context.Send(tree, &Find{Caller: testAct, Key: values[k].Key})
	}
	wg.Wait()
	time.Sleep(1 * time.Second)
}

func newToken() string {
	b := make([]byte, 4)
	_, _ = cr.Read(b)
	return fmt.Sprintf("%x", b)
}
