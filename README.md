# genqlient-todos

A POC GraphQL client for gqlgen-todos built with
[genqlient](https://github.com/Khan/genqlient)

## How To

Ensure gqlgen-todos server is running and then create a Todo.

```bash
export TODO="fix car"
make create
```

Get all the Todos.

```bash
make get
```

