package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ctx = context.WithValue(ctx, "userID", 666)
	result := dbAccess(ctx)

	fmt.Fprintln(w, result)
}

func dbAccess(ctx context.Context) int {
	uid := ctx.Value("userID").(int)
	return uid
}
