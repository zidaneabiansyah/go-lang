package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
}

type Order struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"created_at"`
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func connectDB() (*sql.DB, error) {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	pass := getEnv("DB_PASS", "postgres")
	name := getEnv("DB_NAME", "belajar_go")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func migrate(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			price NUMERIC(12,2) NOT NULL DEFAULT 0,
			stock INTEGER NOT NULL DEFAULT 0,
			created_at TIMESTAMP DEFAULT NOW()
		)`,
		`CREATE TABLE IF NOT EXISTS orders (
			id SERIAL PRIMARY KEY,
			product_id INTEGER NOT NULL REFERENCES products(id),
			quantity INTEGER NOT NULL,
			total NUMERIC(12,2) NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
		)`,
	}

	for _, q := range queries {
		if _, err := db.Exec(q); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Gagal konek DB: %v\n\nPastikan PostgreSQL jalan dan database 'belajar_go' sudah dibuat:\n  createdb belajar_go\n  atau via psql: CREATE DATABASE belajar_go;", err)
	}
	defer db.Close()

	fmt.Println("Koneksi database berhasil")
	fmt.Println()

	if err := migrate(db); err != nil {
		log.Fatal("Migrasi gagal:", err)
	}
	fmt.Println("Migrasi tabel selesai")
	fmt.Println()

	fmt.Println("CRUD — INSERT")
	now := time.Now()
	_, err = db.Exec(`
		INSERT INTO products (name, price, stock, created_at)
		VALUES ($1, $2, $3, $4)
	`, "Laptop Thinkpad", 15000000, 10, now)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Produk 1 berhasil ditambahkan")

	_, err = db.Exec(`
		INSERT INTO products (name, price, stock, created_at)
		VALUES ($1, $2, $3, $4)
	`, "Mouse Wireless", 250000, 50, now)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Produk 2 berhasil ditambahkan")

	items := []struct {
		name  string
		price float64
		stock int
	}{
		{"Keyboard Mechanical", 750000, 30},
		{"Monitor 24 inch", 3500000, 15},
		{"USB Hub 4 port", 150000, 100},
	}

	for _, item := range items {
		_, err := db.Exec(`
			INSERT INTO products (name, price, stock, created_at)
			VALUES ($1, $2, $3, $4)
		`, item.name, item.price, item.stock, now)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Produk 3-5 berhasil ditambahkan")

	fmt.Println()
	fmt.Println("CRUD — SELECT (all products)")
	rows, err := db.Query(`SELECT id, name, price, stock, created_at FROM products ORDER BY id`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt); err != nil {
			log.Fatal(err)
		}
		products = append(products, p)
	}
	for _, p := range products {
		fmt.Printf("  %d. %s — Rp%.0f (stok: %d)\n", p.ID, p.Name, p.Price, p.Stock)
	}

	fmt.Println()
	fmt.Println("CRUD — SELECT by ID")
	var product Product
	err = db.QueryRow(`SELECT id, name, price, stock, created_at FROM products WHERE id = $1`, 1).Scan(
		&product.ID, &product.Name, &product.Price, &product.Stock, &product.CreatedAt,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  Produk ID 1: %s — Rp%.0f\n", product.Name, product.Price)

	fmt.Println()
	fmt.Println("CRUD — UPDATE")
	result, err := db.Exec(`UPDATE products SET stock = $1 WHERE id = $2`, 8, 1)
	if err != nil {
		log.Fatal(err)
	}
	affected, _ := result.RowsAffected()
	fmt.Printf("  Baris terupdate: %d\n", affected)

	fmt.Println()
	fmt.Println("CRUD — DELETE")
	result, err = db.Exec(`DELETE FROM products WHERE id = $1`, 5)
	if err != nil {
		log.Fatal(err)
	}
	affected, _ = result.RowsAffected()
	fmt.Printf("  Baris terhapus: %d\n", affected)

	fmt.Println()
	fmt.Println("PREPARED STATEMENT")
	stmt, err := db.Prepare(`INSERT INTO products (name, price, stock, created_at) VALUES ($1, $2, $3, $4) RETURNING id`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var newID int
	err = stmt.QueryRow("Tablet", 5000000, 20, time.Now()).Scan(&newID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  Prepared insert berhasil, ID: %d\n", newID)

	fmt.Println()
	fmt.Println("TRANSACTION")
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	var productID int
	err = tx.QueryRow(`SELECT id FROM products WHERE name = $1`, "Laptop Thinkpad").Scan(&productID)
	if err != nil {
		log.Fatal(err)
	}

	tx.Exec(`INSERT INTO orders (product_id, quantity, total, created_at) VALUES ($1, $2, $3, $4)`,
		productID, 2, 30000000, time.Now())

	tx.Exec(`UPDATE products SET stock = stock - $1 WHERE id = $2`, 2, productID)

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("  Transaksi berhasil: order 2 laptop + stock berkurang")

	fmt.Println()
	fmt.Println("CLEANUP — hapus data testing")
	db.Exec(`DELETE FROM orders`)
	db.Exec(`DELETE FROM products`)
	fmt.Println("  Data testing dibersihkan")
}
