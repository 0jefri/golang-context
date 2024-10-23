package main

import (
	"context"
	"fmt"
)

func main() {
	// Menggunakan context.Background sebagai context root
	ctx := context.Background()

	fmt.Println(ctx)
}
