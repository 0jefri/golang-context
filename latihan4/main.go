package main

import (
	"context"
	"fmt"
	"time"
	// "golang.org/x/net/context"
)

// buatkan 1 variable tipe int, inisialisasikan dengan jumlah tertentu. kemudian buat sebuah fungsi untukmengurangi nilai dari variable itu, fungsi itu dijalankan pakai goroutine, pengurangan dilakukan setiap 2 detik. kemudian berikan context timeout pada detik ke 4.

func main() {
	jumlah := 100

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go Reduce(ctx, jumlah)

	<-ctx.Done()
	fmt.Println("Program selesai...")
}

func Reduce(ctx context.Context, key int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop context")
		default:
			key--
			fmt.Println("jumlah: ", key)
			time.Sleep(2 * time.Second)
		}
	}
}
