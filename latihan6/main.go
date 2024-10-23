package main

import (
	"context"
	"fmt"
	"time"
)

type User struct {
	Name string
	Role string
}

func checkUserData(user User, ctx context.Context) context.Context {
	if user.Role == "Perawat1" || user.Role == "Perawat2" {
		return context.WithValue(ctx, "selectedUser", user)
	}
	return ctx
}

func displayMenu(ctx context.Context) {
	selectedUser, ok := ctx.Value("selectedUser").(User)
	if !ok {
		fmt.Println("Tidak ada user yang dipilih atau timeout.")
		return
	}

	fmt.Printf("User yang dipilih: %s (%s)\n", selectedUser.Name, selectedUser.Role)
	fmt.Println("Menampilkan menu:")
	switch selectedUser.Role {
	case "Perawat1":
		fmt.Println("1. Tambah Pasien")
		fmt.Println("2. Edit Pasien")
		fmt.Println("3. Update Pasien")
	case "Perawat2":
		fmt.Println("1. Tambah Pasien")
		fmt.Println("2. Edit Pasien")
		fmt.Println("3. Lihat Pasien")
		fmt.Println("4. Hapus Pasien")
	default:
		fmt.Println("Tipe user tidak dikenal")
	}
}

func main() {
	adminUser := User{Name: "Jefri", Role: "Perawat1"}
	memberUser := User{Name: "Agistar", Role: "Perawat2"}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var pilihan string

	for {
		fmt.Println("Pilih user (ketik 'Perawat1' atau 'Perawat2'):")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "Perawat1":
			ctx = checkUserData(adminUser, ctx)
			displayMenu(ctx)
		case "Perawat2":
			ctx = checkUserData(memberUser, ctx)
			displayMenu(ctx)
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
			continue
		}

		break
	}

	time.Sleep(6 * time.Second)
	fmt.Println("\nMencoba menampilkan menu setelah timeout:")
	displayMenu(ctx)
}
