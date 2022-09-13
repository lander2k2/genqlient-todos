// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package todos

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// __createTodoInput is used internally by genqlient
type __createTodoInput struct {
	TodoText string `json:"todoText"`
}

// GetTodoText returns __createTodoInput.TodoText, and is useful for accessing the field via an interface.
func (v *__createTodoInput) GetTodoText() string { return v.TodoText }

// createTodoCreateTodo includes the requested fields of the GraphQL type Todo.
type createTodoCreateTodo struct {
	User createTodoCreateTodoUser `json:"user"`
	Text string                   `json:"text"`
	Done bool                     `json:"done"`
}

// GetUser returns createTodoCreateTodo.User, and is useful for accessing the field via an interface.
func (v *createTodoCreateTodo) GetUser() createTodoCreateTodoUser { return v.User }

// GetText returns createTodoCreateTodo.Text, and is useful for accessing the field via an interface.
func (v *createTodoCreateTodo) GetText() string { return v.Text }

// GetDone returns createTodoCreateTodo.Done, and is useful for accessing the field via an interface.
func (v *createTodoCreateTodo) GetDone() bool { return v.Done }

// createTodoCreateTodoUser includes the requested fields of the GraphQL type User.
type createTodoCreateTodoUser struct {
	Id string `json:"id"`
}

// GetId returns createTodoCreateTodoUser.Id, and is useful for accessing the field via an interface.
func (v *createTodoCreateTodoUser) GetId() string { return v.Id }

// createTodoResponse is returned by createTodo on success.
type createTodoResponse struct {
	CreateTodo createTodoCreateTodo `json:"createTodo"`
}

// GetCreateTodo returns createTodoResponse.CreateTodo, and is useful for accessing the field via an interface.
func (v *createTodoResponse) GetCreateTodo() createTodoCreateTodo { return v.CreateTodo }

// findTodosResponse is returned by findTodos on success.
type findTodosResponse struct {
	Todos []findTodosTodosTodo `json:"todos"`
}

// GetTodos returns findTodosResponse.Todos, and is useful for accessing the field via an interface.
func (v *findTodosResponse) GetTodos() []findTodosTodosTodo { return v.Todos }

// findTodosTodosTodo includes the requested fields of the GraphQL type Todo.
type findTodosTodosTodo struct {
	Text string                 `json:"text"`
	Done bool                   `json:"done"`
	User findTodosTodosTodoUser `json:"user"`
}

// GetText returns findTodosTodosTodo.Text, and is useful for accessing the field via an interface.
func (v *findTodosTodosTodo) GetText() string { return v.Text }

// GetDone returns findTodosTodosTodo.Done, and is useful for accessing the field via an interface.
func (v *findTodosTodosTodo) GetDone() bool { return v.Done }

// GetUser returns findTodosTodosTodo.User, and is useful for accessing the field via an interface.
func (v *findTodosTodosTodo) GetUser() findTodosTodosTodoUser { return v.User }

// findTodosTodosTodoUser includes the requested fields of the GraphQL type User.
type findTodosTodosTodoUser struct {
	Name string `json:"name"`
}

// GetName returns findTodosTodosTodoUser.Name, and is useful for accessing the field via an interface.
func (v *findTodosTodosTodoUser) GetName() string { return v.Name }

func createTodo(
	ctx context.Context,
	client graphql.Client,
	todoText string,
) (*createTodoResponse, error) {
	req := &graphql.Request{
		OpName: "createTodo",
		Query: `
mutation createTodo ($todoText: String!) {
	createTodo(input: {text:$todoText,userId:"1"}) {
		user {
			id
		}
		text
		done
	}
}
`,
		Variables: &__createTodoInput{
			TodoText: todoText,
		},
	}
	var err error

	var data createTodoResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func findTodos(
	ctx context.Context,
	client graphql.Client,
) (*findTodosResponse, error) {
	req := &graphql.Request{
		OpName: "findTodos",
		Query: `
query findTodos {
	todos {
		text
		done
		user {
			name
		}
	}
}
`,
	}
	var err error

	var data findTodosResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
