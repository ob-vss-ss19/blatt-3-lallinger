package main

import (
	"flag"
	"fmt"
)

var id *int
var token *string
var forceDelete *bool

func main() {

	id = flag.Int("id", -1, "tree id")
	token = flag.String("token", "", "tree token")
	forceDelete = flag.Bool("no-preserve-tree", false, "force deletion of tree")
	flag.Parse()

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
		delete()
		return
	case "traverse":
		traverse()
		return
	default:
		error()
	}
}

func newTree() {
	if len(flag.Args()) > 1 {
		error()
		return
	}
}

func insert() {
	if len(flag.Args()) != 3 || *id == -1 || *token == "" {
		error()
		return
	}
}

func search() {
	if len(flag.Args()) != 2 || *id == -1 || *token == "" {
		error()
		return
	}
}

func remove() {
	if len(flag.Args()) != 2 || *id == -1 || *token == "" {
		error()
		return
	}
}

func delete() {
	if !*forceDelete {
		fmt.Println("Specify no-preserve-tree for tree deletion")
		return
	}

	if len(flag.Args()) > 1 || *id == -1 || *token == "" {
		error()
		return
	}
}

func traverse() {
	if len(flag.Args()) > 1 || *id == -1 || *token == "" {
		error()
		return
	}
}

func error() {
	fmt.Println("Invalid arguments")
}
