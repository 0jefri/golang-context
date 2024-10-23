package main

import (
	"context"
	"fmt"
	"time"
)

//buatkan fungsi untuk menambahkan nilai kedalam context, kemudian buat 1 fungsi untuk menampilkan key, valuenya setiap 2 detik. buatkan lagi 1 fungsi untuk mengubah nilai value sebelumnya.

func main() {

}

func addValue(ctx context.Context, key string, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

func showValue(ctx context.Context) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context berhenti:", ctx.Err())
			return
		case <-ticker.C:
			if value := ctx.Value("key"); value != nil {
				fmt.Printf("key: key, value: %v\n", value)
			} else {
				fmt.Println("key tidak ditemukan")
			}
		}
	}
}
