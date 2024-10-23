package main

import (
	"context"
	"fmt"
	"time"
)

// 3 fungsi, fungsi pertama mencetak sebuah text setiap 2 detik , fungsi kedua mencetak text setiap 1 detik, fungsi ketiga mencetak tekx setiap 3 detik. buat context untuk menghentikan semua fungsi yg berjalan di detik kelima

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	go Text1(ctx)
	go Text2(ctx)
	go Text3(ctx)

	time.Sleep(6 * time.Second)
	fmt.Println("Program selesai...")
}

func Text1(ctx context.Context) {
	// value := ctx.Value(key)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context dibatalkan:", ctx.Err())
			return
		case <-ticker.C:
			fmt.Println("Ini Text untuk 2 detik")
		}
	}
}

func Text2(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context dibatalkan:", ctx.Err())
			return
		case <-ticker.C:
			fmt.Println("Ini Text untuk 1 detik")
		}
	}
}

func Text3(ctx context.Context) {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context dibatalkan:", ctx.Err())
			return
		case <-ticker.C:
			fmt.Println("Ini Text untuk 3 detik")
		}
	}
}
