package main

import (
	"context"
	"fmt"
	"time"
)

// Struktur data untuk menyimpan informasi
type Item struct {
	Name  string
	Value int
}

// Fungsi untuk memasukkan data struct ke dalam slice
func addItem(slice *[]Item, item Item) {
	*slice = append(*slice, item)
}

func main() {
	// Membuat slice untuk menyimpan Item
	var items []Item

	// Mengatur deadline untuk program berhenti setelah 10 detik
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Pastikan cancel dipanggil untuk membebaskan sumber daya

	// Memasukkan data ke dalam slice dalam goroutine
	go func() {
		for i := 0; i < 5; i++ {
			item := Item{Name: fmt.Sprintf("Item %d", i+1), Value: (i + 1) * 10}
			addItem(&items, item)
			fmt.Printf("Ditambahkan: %s dengan Value %d\n", item.Name, item.Value)
			time.Sleep(2 * time.Second) // Tidur selama 2 detik sebelum menambahkan item berikutnya
		}
	}()

	// Menunggu sampai deadline tercapai
	<-ctx.Done() // Menunggu hingga context selesai (baik karena timeout atau dibatalkan)

	// Menampilkan semua item yang telah ditambahkan
	fmt.Println("\nSemua item yang telah ditambahkan:")
	for _, item := range items {
		fmt.Printf("Name: %s, Value: %d\n", item.Name, item.Value)
	}

	fmt.Println("Program selesai.")
}
