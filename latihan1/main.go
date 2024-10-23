package main

import (
	"context"
	"fmt"
)

func main() {
	key := "angka"
	ctx := context.WithValue(context.Background(), key, 5)

	ShowValue(ctx, key)
	CalculateMath(ctx, key)
}

func ShowValue(ctx context.Context, key string) {
	value := ctx.Value(key)

	if value != nil {
		fmt.Printf("nilai dari %s adalah %d\n", key, value)
	} else {
		fmt.Printf("tidak ada nilai untuk %s", key)
	}
}

func CalculateMath(ctx context.Context, key string) {
	if value, ok := ctx.Value(key).(int); ok {
		result := value * 2
		fmt.Printf("Hasil dari %d * 2 = %d\n", value, result)
	} else {
		fmt.Println("nilai tidak ditemukan....")
	}
}
