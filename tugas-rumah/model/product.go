package model

import (
	"context"
	"fmt"
)

type Product struct {
	Name  string
	Price int
}

type Keranjang struct {
	Product []Product
}

var products = []Product{
	{Name: "Handphone", Price: 1500000},
	{Name: "Laptop", Price: 800000},
	{Name: "Televisi", Price: 1000000},
	{Name: "Powerbank", Price: 300000},
	{Name: "Radio", Price: 200000},
}

func ShowProduct() {
	fmt.Println("\nDaftar Produk:")
	for i, p := range products {
		fmt.Printf("%d. %s - %d\n", i+1, p.Name, p.Price)
	}
}

func TambahKeranjang(keranjang *Keranjang) {
	var pilih int
	ShowProduct()
	fmt.Println("\nPilih produk untuk dimasukkan ke keranjang:")
	fmt.Scanln(&pilih)
	if pilih > 0 && pilih <= len(products) {
		keranjang.Product = append(keranjang.Product, products[pilih-1])
		fmt.Printf("%s ditambahkan ke keranjang.\n", products[pilih-1].Name)
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func TampilkanKeranjang(keranjang Keranjang) {
	fmt.Println("\nKeranjang Anda:")
	if len(keranjang.Product) == 0 {
		fmt.Println("Keranjang kosong.")
		return
	}
	total := 0
	for i, item := range keranjang.Product {
		fmt.Printf("%d. %s - %d\n", i+1, item.Name, item.Price)
		total += item.Price
	}
	fmt.Printf("Total: %d\n", total)
}

func checkout(keranjang Keranjang) {
	if len(keranjang.Product) == 0 {
		fmt.Println("Keranjang kosong, tidak ada yang bisa di-checkout.")
		return
	}

	var amount int
	TampilkanKeranjang(keranjang)
	fmt.Println("Masukkan jumlah pembayaran:")
	fmt.Scanln(&amount)

	total := 0
	for _, item := range keranjang.Product {
		total += item.Price
	}

	if amount >= total {
		fmt.Println("Pembayaran berhasil, barang segera dikirim!")
	} else {
		fmt.Println("Pembayaran kurang, silakan coba lagi.")
	}
}

func MainMenu(ctx context.Context, keranjang *Keranjang) {
	for {
		fmt.Println("\nMenu Utama:")
		fmt.Println("1. Tampilkan Daftar Produk")
		fmt.Println("2. Masukkan Produk ke Keranjang")
		fmt.Println("3. Lihat Keranjang")
		fmt.Println("4. Checkout")
		fmt.Println("5. Keluar")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			ShowProduct()
		case 2:
			TambahKeranjang(keranjang)
		case 3:
			TampilkanKeranjang(*keranjang)
		case 4:
			checkout(*keranjang)
		case 5:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
