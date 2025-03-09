package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Year   int     `json:"year"`
	Genre  string  `json:"genre"`
	Rating float64 `json:"rating"`
	Author Author  `json:"author"`
}

var books []Book

// Handler untuk mendapatkan semua buku
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Fetching all books")
	json.NewEncoder(w).Encode(books) // BUG: Tidak menangani kasus ketika books kosong
	//cek apakah book kosong
	if len(books) == 0 {
		http.Error(w, `{"message": "No books found"}`, http.StatusNotFound)
		return
	}
}

// Handler untuk mendapatkan detail buku berdasarkan ID
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	log.Println("Fetching book with ID:", params["id"])

	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	// mengubah error ke 404
	log.Println("Book not found")
	http.Error(w, "Book not found", http.StatusNotFound)

}

// Handler untuk menambahkan buku baru
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	//check apakah ID sudah dipakai
	for _, existingBook := range books {
		if existingBook.ID == book.ID {
			log.Println("Duplicate ID:", book.ID)
			http.Error(w, "Book ID already exists", http.StatusConflict) // 409 Conflict
			return
		}
	}
	// BUG: Tidak ada validasi untuk memastikan ID unik
	log.Println("Adding new book:", book.Title)
	books = append(books, book)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// Handler untuk memperbarui buku berdasarkan ID
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	log.Println("Updating book with ID:", params["id"])

	for index, book := range books {
		if book.ID == params["id"] {
			var updatedBook Book
			err := json.NewDecoder(r.Body).Decode(&updatedBook)
			if err != nil {
				log.Println("Failed to decode request body:", err)
				http.Error(w, "Invalid input", http.StatusBadRequest)
				return
			}
			//memastikan id buku gtetap sama
			updatedBook.ID = book.ID
			log.Println("Updating book with ID:", params["id"])
			books[index] = updatedBook // BUG: ID lama bisa berubah
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

// Handler untuk menghapus buku berdasarkan ID
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	log.Println("Deleting book with ID:", params["id"])

	for index, book := range books {
		if book.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)

			//perbaikan menambahkan respons json
			log.Println("Penghapusan Berhasil")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted successfully"})
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

// Fungsi utama untuk menjalankan server
func main() {
	r := mux.NewRouter()

	// Data contoh awal
	books = append(books, Book{
		ID:     "1",
		Title:  "Golang for Beginners",
		Year:   2021,
		Genre:  "Programming",
		Rating: 4.5,
		Author: Author{FirstName: "John", LastName: "Doe"},
	})
	books = append(books, Book{
		ID:     "2",
		Title:  "Advanced Golang",
		Year:   2023,
		Genre:  "Programming",
		Rating: 4.8,
		Author: Author{FirstName: "Jane", LastName: "Smith"},
	})

	// Route API
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	// Jalankan server
	fmt.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
