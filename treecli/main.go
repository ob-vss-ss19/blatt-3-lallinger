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
		switch msg.Type {
		case messages.CREATE:
			fmt.Printf("Id: %d\n", msg.Key)
			fmt.Printf("Token: %s\n", msg.Value)
		case messages.FIND:
			fmt.Printf("Value: %s\n", msg.Value)
		case messages.SUCCESS:
			fmt.Printf("Success")
		}
		wg.Done()
	case *messages.Traverse:
		for i, pair := range msg.Values {
			fmt.Printf("{%d,%s}", pair.Key, pair.Value)
			if i+1 < len(msg.Values) {
				fmt.Printf(",")
			}
		}
		fmt.Printf("\n")
		wg.Done()
	case *messages.Error:
		fmt.Printf("%s\n", msg.Message)
		wg.Done()
	}

}

var (
	flagBind   = flag.String("bind", "localhost:8092", "Bind to address")
	flagRemote = flag.String("remote", "127.0.0.1:8093", "remote host:port")

	id          *int
	token       = flag.String("token", "", "tree token")
	forceDelete = flag.Bool("no-preserve-tree", false, "force deletion of tree")

	rootContext *actor.RootContext
	pid         *actor.PID
	remotePid   *actor.PID
	wg          sync.WaitGroup
)

func main() {

	id = flag.Int("id", -1, "tree id")

	flag.Parse()

	remote.Start(*flagBind)
	props := actor.PropsFromProducer(func() actor.Actor {
		wg.Add(1)
		return &CliActor{}
	})
	rootContext = actor.EmptyRootContext
	pid = rootContext.Spawn(props)

	time.Sleep(5 * time.Second)

	pidResp, err := remote.SpawnNamed(*flagRemote, "remote", "treeService", 5*time.Second)
	if err != nil {
		panic(err)
	}
	remotePid = pidResp.Pid

	fmt.Printf("token: %s id: %d args: ", *token, *id)
	for _, arg := range flag.Args() {
		fmt.Printf("%s ", arg)
	}
	fmt.Printf("\n")

	switch flag.Args()[0] {
	case "newtree":
		newTree()
	case "insert":
		insert()
	case "search":
		search()
	case "remove":
		remove()
	case "delete":
		deleteTree()
	case "traverse":
		traverse()
	default:
		printError()
		return
	}
	wg.Wait()
}

func newTree() {
	if len(flag.Args()) != 2 {
		printError()
		return
	}
	tmp, _ := strconv.Atoi(flag.Args()[1])
	rootContext.RequestWithCustomSender(remotePid, &messages.Request{Type: messages.CREATE, Id: int32(tmp)}, pid)
}

func insert() {
	if len(flag.Args()) != 3 || *id == -1 || *token == "" {
		printError()
		return
	}
	tmp, _ := strconv.Atoi(flag.Args()[1])
	rootContext.RequestWithCustomSender(remotePid, &messages.Request{Type: messages.ADD, Key: int32(tmp), Value: flag.Args()[2], Token: *token, Id: int32(*id)}, pid)
}

func search() {
	if len(flag.Args()) != 2 || *id == -1 || *token == "" {
		printError()
		return
	}
	tmp, _ := strconv.Atoi(flag.Args()[1])
	rootContext.RequestWithCustomSender(remotePid, &messages.Request{Type: messages.FIND, Key: int32(tmp), Token: *token, Id: int32(*id)}, pid)
}

func remove() {
	if len(flag.Args()) != 2 || *id == -1 || *token == "" {
		printError()
		return
	}
	tmp, _ := strconv.Atoi(flag.Args()[1])
	rootContext.RequestWithCustomSender(remotePid, &messages.Request{Type: messages.REMOVE, Key: int32(tmp), Token: *token, Id: int32(*id)}, pid)
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
	rootContext.RequestWithCustomSender(remotePid, &messages.Request{Type: messages.DELETE, Token: *token, Id: int32(*id)}, pid)
}

func traverse() {
	if len(flag.Args()) > 1 || *id == -1 || *token == "" {
		printError()
		return
	}
	rootContext.RequestWithCustomSender(remotePid, &messages.Request{Type: messages.TRAVERSE, Token: *token, Id: int32(*id)}, pid)
}

func printError() {
	fmt.Println("Invalid arguments")
	wg.Done()
}
