package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

var (
	items   []Item
	nextID  int
	mu      sync.RWMutex
)

func init() {
	items = []Item{
		{ID: 1, Name: "Laptop", Price: 15000000},
		{ID: 2, Name: "Mouse", Price: 250000},
	}
	nextID = 3
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Printf("[%s] %s selesai dalam %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		writeJSON(w, http.StatusNotFound, APIResponse{
			Success: false, Message: "endpoint gak ditemukan",
		})
		return
	}
	writeJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Message: "selamat datang di REST API belajar Golang",
	})
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/items")
	idStr = strings.TrimPrefix(idStr, "/")

	if idStr == "" {
		switch r.Method {
		case http.MethodGet:
			getAllItems(w, r)
		case http.MethodPost:
			createItem(w, r)
		default:
			writeJSON(w, http.StatusMethodNotAllowed, APIResponse{
				Success: false, Message: "method gak diizinkan",
			})
		}
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, APIResponse{
			Success: false, Message: "ID harus angka",
		})
		return
	}

	switch r.Method {
	case http.MethodGet:
		getItemByID(w, r, id)
	case http.MethodPut:
		updateItem(w, r, id)
	case http.MethodDelete:
		deleteItem(w, r, id)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, APIResponse{
			Success: false, Message: "method gak diizinkan",
		})
	}
}

func getAllItems(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	writeJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Data:    items,
	})
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		writeJSON(w, http.StatusBadRequest, APIResponse{
			Success: false, Message: "format JSON salah",
		})
		return
	}

	if newItem.Name == "" {
		writeJSON(w, http.StatusBadRequest, APIResponse{
			Success: false, Message: "name wajib diisi",
		})
		return
	}
	if newItem.Price <= 0 {
		writeJSON(w, http.StatusBadRequest, APIResponse{
			Success: false, Message: "price harus lebih dari 0",
		})
		return
	}

	mu.Lock()
	newItem.ID = nextID
	nextID++
	items = append(items, newItem)
	mu.Unlock()

	writeJSON(w, http.StatusCreated, APIResponse{
		Success: true,
		Message: "item berhasil ditambahkan",
		Data:    newItem,
	})
}

func getItemByID(w http.ResponseWriter, r *http.Request, id int) {
	mu.RLock()
	defer mu.RUnlock()

	for _, item := range items {
		if item.ID == id {
			writeJSON(w, http.StatusOK, APIResponse{
				Success: true,
				Data:    item,
			})
			return
		}
	}

	writeJSON(w, http.StatusNotFound, APIResponse{
		Success: false, Message: fmt.Sprintf("item dengan ID %d gak ketemu", id),
	})
}

func updateItem(w http.ResponseWriter, r *http.Request, id int) {
	var updated Item
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		writeJSON(w, http.StatusBadRequest, APIResponse{
			Success: false, Message: "format JSON salah",
		})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, item := range items {
		if item.ID == id {
			if updated.Name != "" {
				items[i].Name = updated.Name
			}
			if updated.Price > 0 {
				items[i].Price = updated.Price
			}
			writeJSON(w, http.StatusOK, APIResponse{
				Success: true,
				Message: "item berhasil diupdate",
				Data:    items[i],
			})
			return
		}
	}

	writeJSON(w, http.StatusNotFound, APIResponse{
		Success: false, Message: fmt.Sprintf("item dengan ID %d gak ketemu", id),
	})
}

func deleteItem(w http.ResponseWriter, r *http.Request, id int) {
	mu.Lock()
	defer mu.Unlock()

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			writeJSON(w, http.StatusOK, APIResponse{
				Success: true,
				Message: fmt.Sprintf("item %d berhasil dihapus", id),
			})
			return
		}
	}

	writeJSON(w, http.StatusNotFound, APIResponse{
		Success: false, Message: fmt.Sprintf("item dengan ID %d gak ketemu", id),
	})
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	minPrice := r.URL.Query().Get("min_price")

	mu.RLock()
	defer mu.RUnlock()

	result := items
	if name != "" {
		var filtered []Item
		for _, item := range result {
			if strings.Contains(strings.ToLower(item.Name), strings.ToLower(name)) {
				filtered = append(filtered, item)
			}
		}
		result = filtered
	}
	if minPrice != "" {
		min, err := strconv.ParseFloat(minPrice, 64)
		if err == nil {
			var filtered []Item
			for _, item := range result {
				if item.Price >= min {
					filtered = append(filtered, item)
				}
			}
			result = filtered
		}
	}

	writeJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Data:    result,
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/items", itemsHandler)
	mux.HandleFunc("/items/", itemsHandler)
	mux.HandleFunc("/search", queryHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      loggingMiddleware(mux),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("REST API jalan di http://localhost:8080")
	fmt.Println()
	fmt.Println("GET    /")
	fmt.Println("GET    /items")
	fmt.Println("POST   /items")
	fmt.Println("GET    /items/:id")
	fmt.Println("PUT    /items/:id")
	fmt.Println("DELETE /items/:id")
	fmt.Println("GET    /search?name=&min_price=")
	fmt.Println()

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("server error:", err)
	}
}
