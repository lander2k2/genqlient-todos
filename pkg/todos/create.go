package todos

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

func CreateTodo(
	todoText string,
	ctx context.Context,
	client graphql.Client,
) (*createTodoResponse, error) {
	resp, err := createTodo(ctx, client, todoText)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
