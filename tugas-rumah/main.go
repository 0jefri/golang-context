package main

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-context/tugas-rumah/model"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	keranjang := model.Keranjang{}

	fmt.Println("Login")
	if login(ctx) {
		fmt.Println("Login berhasil.")
		model.MainMenu(ctx, &keranjang)
	} else {
		fmt.Println("Login gagal, silakan coba lagi.")
	}
}

// func mainMenu(ctx context.Context, keranjang *model.Keranjang) {
// 	panic("unimplemented")
// }

func login(ctx context.Context) bool {
	var username, password string
	fmt.Println("Masukkan username:")
	fmt.Scanln(&username)
	fmt.Println("Masukkan password:")
	fmt.Scanln(&password)

	if username == "jefri" && password == "12345" {
		return true
	}

	select {
	case <-ctx.Done():
		fmt.Println("Sesi login habis, silakan login kembali.")
		return false
	default:
		fmt.Println("Login gagal, silakan coba lagi.")
		return false
	}
}
