package main

import (
	"fmt"
	"strings"
)

// Struct untuk mendefinisikan objek Buku
type Book struct {
	Title  string
	Author string
	ISBN   string
}

// Interface perpustakaan
type Library interface {
	AddBook(*Book)                // Menambahkan buku
	RemoveBook(isbn string) error // Menghapus buku berdasarkan ISBN
	ShowBooks()                   // Menampilkan semua buku
}

// Struct untuk perpustakaan
type MyLibrary struct {
	Books []Book
}

// Implementasi metode untuk menambahkan buku
func (l *MyLibrary) AddBook(book *Book) {
	// Tambahkan buku ke slice Books
	l.Books = append(l.Books, *book)
	fmt.Println("Buku berhasil ditambahkan dengan ISBN:", book.ISBN)
	// Setelah berhasil menambahkan buku, tanyakan kepada pengguna
	l.postAddMenu()
}

// Implementasi metode untuk menghapus buku berdasarkan ISBN
func (l *MyLibrary) RemoveBook(isbn string) error {
	// Loop untuk menemukan dan menghapus buku berdasarkan ISBN
	for i, b := range l.Books {
		if b.ISBN == isbn {
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			fmt.Println("Buku berhasil dihapus!")
			return nil
		}
	}
	// Jika ISBN tidak ditemukan
	return fmt.Errorf("Buku dengan ISBN %s tidak ditemukan", isbn)
}

// Implementasi metode untuk menampilkan semua buku
func (l *MyLibrary) ShowBooks() {
	if len(l.Books) == 0 {
		fmt.Println("Tidak ada buku di perpustakaan.")
		return
	}
	fmt.Println("Daftar Buku di Perpustakaan:")
	for _, b := range l.Books {
		fmt.Printf("Judul: %s, Penulis: %s, ISBN: %s\n", b.Title, b.Author, b.ISBN)
	}
}

// Fungsi untuk mendapatkan input buku dari pengguna
func inputBookDetails() *Book {
	var title, author, isbn string
	fmt.Print("Masukkan Judul: ")
	fmt.Scanln(&title)
	fmt.Print("Masukkan Penulis: ")
	fmt.Scanln(&author)
	fmt.Print("Masukkan ISBN 'min. 5 kata': ")
	fmt.Scanln(&isbn)
	// Menghapus spasi tambahan di ISBN
	isbn = strings.TrimSpace(isbn)
	return &Book{Title: title, Author: author, ISBN: isbn}
}

// Fungsi untuk memvalidasi ISBN
func validateISBN(isbn string) bool {
	isbn = strings.TrimSpace(isbn) // Menghapus spasi di awal/akhir
	if len(isbn) == 0 {
		return false // ISBN tidak boleh kosong
	}
	// Validasi dasar, ISBN minimal harus memiliki panjang 5 karakter
	if len(isbn) < 5 {
		return false
	}
	return true
}

// Menampilkan pilihan setelah buku ditambahkan
func (l *MyLibrary) postAddMenu() {
	fmt.Println("\nBuku berhasil ditambahkan. Apa yang ingin Anda lakukan?")
	fmt.Println("1. Tampilkan Daftar Buku")
	fmt.Println("2. Kembali ke Menu Utama")
	fmt.Print("Pilih opsi (1-2): ")
	var postAddChoice int
	fmt.Scanln(&postAddChoice)

	switch postAddChoice {
	case 1:
		l.ShowBooks()
	case 2:
		// Kembali ke menu utama, tidak perlu tindakan karena loop utama akan melanjutkan
		fmt.Println("Kembali ke menu utama...")
	default:
		fmt.Println("Pilihan tidak valid. Kembali ke menu utama...")
	}
}

// Fungsi utama
func main() {
	var myLibrary MyLibrary
	for {
		// Menu
		fmt.Println("\n===== Menu Manajemen Perpustakaan =====")
		fmt.Println("1. Tambah Buku")
		fmt.Println("2. Hapus Buku")
		fmt.Println("3. Tampilkan Daftar Buku")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih opsi (1-4): ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// Tambah buku
			book := inputBookDetails()
			if !validateISBN(book.ISBN) {
				fmt.Println("ISBN tidak valid! Pastikan ISBN memiliki panjang minimal 10 karakter.")
			} else {
				myLibrary.AddBook(book)
			}
		case 2:
			// Hapus buku
			fmt.Print("Masukkan ISBN buku yang akan dihapus: ")
			var isbn string
			fmt.Scanln(&isbn)
			isbn = strings.TrimSpace(isbn) // Hapus spasi ekstra
			err := myLibrary.RemoveBook(isbn)
			if err != nil {
				fmt.Println(err.Error())
			}
		case 3:
			// Tampilkan buku
			myLibrary.ShowBooks()
		case 4:
			// Keluar dari program
			fmt.Println("Keluar dari program...")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih opsi 1-4.")
		}
	}
}
