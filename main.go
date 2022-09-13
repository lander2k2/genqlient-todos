package main

//go:generate genqlient

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/Khan/genqlient/graphql"
	"github.com/lander2k2/genqlient-todos/pkg/todos"
)

func main() {
	todoText := flag.String("todo", "", "The todo text")
	flag.Parse()

	ctx := context.Background()
	client := graphql.NewClient("http://localhost:8080/query", http.DefaultClient)

	if *todoText != "" {
		resp, err := todos.CreateTodo(*todoText, ctx, client)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", resp)
	} else {
		resp, err := todos.GetTodos(ctx, client)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", resp)
	}
}
