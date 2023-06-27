package main

import (
	"context"
)

func main() {
	ctx := context.Background()

	store := dbConnect(ctx)
	defer store.Close(ctx)

	runRPC(store)
}
