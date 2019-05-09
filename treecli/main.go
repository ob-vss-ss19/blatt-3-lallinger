package main

import (
	"flag"
	"fmt"
	"github.com/ob-vss-ss19/blatt-3-lallinger/messages"
	"strconv"
	"sync"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
)

type CliActor struct {
}

func (state *CliActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *messages.Response:
		if msg.Type == messages.Usage_CREATE {
			fmt.Println("Token: %s", msg.Value)
			fmt.Println("Id: %d", msg.Key)
		}

	case *messages.Traverse:
	case *messages.Error:
	}
}

var flagBind = flag.String("bind", "localhost:8092", "Bind to address")

var flagRemote = flag.String("remote", "127.0.0.1:8091", "remote host:port")

var id = flag.Int("id", -1, "tree id")
var token = flag.String("token", "", "tree token")
var forceDelete = flag.Bool("no-preserve-tree", false, "force deletion of tree")

var rootContext *actor.RootContext
var pid *actor.PID
var remotePid *actor.PID
var wg sync.WaitGroup

func main() {

	flag.Parse()

	remote.Start(*flagBind)

	props := actor.PropsFromProducer(func() actor.Actor {
		wg.Add(1)
		return &CliActor{}
	})
	rootContext = actor.EmptyRootContext
	pid = rootContext.Spawn(props)

	message := &messages.Error{}

	time.Sleep(5 * time.Second)

	pidResp, err := remote.SpawnNamed(*flagRemote, "remote", "treeService", 5*time.Second)
	if err != nil {
		panic(err)
	}
	remotePid = pidResp.Pid

	rootContext.RequestWithCustomSender(remotePid, message, pid)

	wg.Wait()

	switch flag.Args()[0] {
	case "newtree":
		newTree()
		return
	case "insert":
		insert()
		return
	case "search":
		search()
		return
	case "remove":
		remove()
		return
	case "delete":
		deleteTree()
		return
	case "traverse":
		traverse()
		return
	default:
		printError()
	}
}

func newTree() {
	if len(flag.Args()) > 1 {
		printError()
		return
	}

	message := &messages.Error{}

	rootContext.RequestWithCustomSender(remotePid, message, pid)
	//rootContext.RequestWithCustomSender(remotePid, &messages.Request{Type: messages.Usage_CREATE}, pid)
	wg.Wait()
}

func insert() {
	if len(flag.Args()) != 3 || *id == -1 || *token == "" {
		printError()
		return
	}
	tmp, _ := strconv.Atoi(flag.Args()[1])
	rootContext.RequestWithCustomSender(remotePid, &messages.Request{Type: messages.Usage_ADD, Key: int32(tmp), Value: flag.Args()[2], Token: *token, Id: int32(*id)}, pid)
}

func search() {
	if len(flag.Args()) != 2 || *id == -1 || *token == "" {
		printError()
		return
	}
	tmp, _ := strconv.Atoi(flag.Args()[1])
	rootContext.RequestWithCustomSender(remotePid, &messages.Request{Type: messages.Usage_FIND, Key: int32(tmp), Token: *token, Id: int32(*id)}, pid)
	wg.Wait()
}

func remove() {
	if len(flag.Args()) != 2 || *id == -1 || *token == "" {
		printError()
		return
	}
	tmp, _ := strconv.Atoi(flag.Args()[1])
	rootContext.RequestWithCustomSender(remotePid, &messages.Request{Type: messages.Usage_REMOVE, Key: int32(tmp), Token: *token, Id: int32(*id)}, pid)
}

func deleteTree() {
	if !*forceDelete {
		fmt.Println("Specify no-preserve-tree for tree deletion")
		return
	}

	if len(flag.Args()) > 1 || *id == -1 || *token == "" {
		printError()
		return
	}
	rootContext.RequestWithCustomSender(remotePid, &messages.Request{Type: messages.Usage_DELETE, Token: *token, Id: int32(*id)}, pid)
}

func traverse() {
	if len(flag.Args()) > 1 || *id == -1 || *token == "" {
		printError()
		return
	}
	rootContext.RequestWithCustomSender(remotePid, &messages.Request{Type: messages.Usage_TRAVERSE, Token: *token, Id: int32(*id)}, pid)
	wg.Wait()
}

func printError() {
	fmt.Println("Invalid arguments")
}
