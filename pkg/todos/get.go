package todos

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

func GetTodos(
	ctx context.Context,
	client graphql.Client,
) (*findTodosResponse, error) {
	resp, err := findTodos(ctx, client)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
