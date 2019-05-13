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
			wg.Done()
		case messages.FIND:
			fmt.Printf("Value: %s\n", msg.Value)
			wg.Done()

		case messages.SUCCESS:
			// fmt.Printf("\nSuccess")
			wg.Done()
		case messages.TREES:
			fmt.Printf("Tree IDs: %s\n", msg.Value)
			wg.Done()
		}
	case *messages.Traverse:
		for i, v := range msg.Values {
			fmt.Printf("{%d,%s}", v.Key, v.Value)
			if i+1 < len(msg.Values) {
				fmt.Printf(", ")
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

	origUsage := flag.Usage

	flag.Usage = func() {
		fmt.Println("treecli [FLAGS] COMMAND [KEY/SIZE] [VALUE]")
		fmt.Println("FLAGS")
		origUsage()
		fmt.Println()
		fmt.Println("COMMAND")
		fmt.Println("  newtree SIZE")
		fmt.Println("	Creates new tree. SIZE parameter specifies leaf size (minimum 1). Returns id and token")
		fmt.Println("  insert KEY VALUE")
		fmt.Println("	Insert an integer KEY with given string VALUE into the tree. id and token flag must be specified")
		fmt.Println("  search KEY")
		fmt.Println("	Search the tree for KEY. Returns corresponding value if found. id and token flag must be specified")
		fmt.Println("  remove KEY")
		fmt.Println("	Removes the KEY from the tree. id and token flag must be specified")
		fmt.Println("  traverse")
		fmt.Println("	gets all keys and values in the tree sorted by keys. id and token flag must be specified")
		fmt.Println("  trees")
		fmt.Println("	Gets a list of all available tree ids")
		fmt.Println("  delete")
		fmt.Println("	Deletes the tree. id and token flag must be specified, also no-preserve-tree flag must be set to true")
		fmt.Println("")
		fmt.Println("Example:")
		fmt.Println("  newtree 2")
		fmt.Println("  --id=0 --token=d57a23df insert 42 'the answer'")
	}

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
	case "trees":
		trees()
	default:
		printError()
		return
	}
	wg.Wait()
}

func trees() {
	if len(flag.Args()) != 1 {
		printError()
		return
	}
	rootContext.RequestWithCustomSender(remotePid, &messages.Request{Type: messages.TREES}, pid)
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
	if len(flag.Args()) > 1 || *id == -1 || *token == "" {
		printError()
		return
	}

	if !*forceDelete {
		fmt.Println("Specify no-preserve-tree for tree deletion")
		wg.Done()
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
	flag.Usage()
	wg.Done()
}
