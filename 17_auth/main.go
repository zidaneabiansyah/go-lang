package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

var users []User
var jwtSecret []byte

type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func init() {
	secret := make([]byte, 32)
	rand.Read(secret)
	jwtSecret = secret
	fmt.Printf("JWT Secret: %s\n\n", hex.EncodeToString(secret))

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	users = append(users, User{ID: 1, Username: "admin", Password: string(hashedPassword), Role: "admin"})

	hashedPassword2, _ := bcrypt.GenerateFromPassword([]byte("user123"), bcrypt.DefaultCost)
	users = append(users, User{ID: 2, Username: "user", Password: string(hashedPassword2), Role: "user"})
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateToken(user User) (string, error) {
	claims := Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "belajar-golang",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func validateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("metode signing tidak valid: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("token tidak valid")
	}

	return claims, nil
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			writeJSON(w, http.StatusUnauthorized, APIResponse{
				Success: false, Message: "header Authorization diperlukan",
			})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			writeJSON(w, http.StatusUnauthorized, APIResponse{
				Success: false, Message: "format Authorization: Bearer <token>",
			})
			return
		}

		claims, err := validateToken(parts[1])
		if err != nil {
			writeJSON(w, http.StatusUnauthorized, APIResponse{
				Success: false, Message: fmt.Sprintf("token invalid: %v", err),
			})
			return
		}

		r.Header.Set("X-User-ID", fmt.Sprintf("%d", claims.UserID))
		r.Header.Set("X-User-Role", claims.Role)
		next(w, r)
	}
}

func adminOnly(next http.HandlerFunc) http.HandlerFunc {
	return authMiddleware(func(w http.ResponseWriter, r *http.Request) {
		role := r.Header.Get("X-User-Role")
		if role != "admin" {
			writeJSON(w, http.StatusForbidden, APIResponse{
				Success: false, Message: "khusus admin",
			})
			return
		}
		next(w, r)
	})
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, APIResponse{
			Success: false, Message: "format JSON salah",
		})
		return
	}

	if req.Username == "" || req.Password == "" {
		writeJSON(w, http.StatusBadRequest, APIResponse{
			Success: false, Message: "username dan password wajib diisi",
		})
		return
	}

	for _, u := range users {
		if u.Username == req.Username {
			writeJSON(w, http.StatusConflict, APIResponse{
				Success: false, Message: "username sudah dipakai",
			})
			return
		}
	}

	hashed, err := hashPassword(req.Password)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, APIResponse{
			Success: false, Message: "gagal hash password",
		})
		return
	}

	role := req.Role
	if role != "admin" && role != "user" {
		role = "user"
	}

	newUser := User{
		ID:       len(users) + 1,
		Username: req.Username,
		Password: hashed,
		Role:     role,
	}
	users = append(users, newUser)

	writeJSON(w, http.StatusCreated, APIResponse{
		Success: true,
		Message: "registrasi berhasil",
		Data: map[string]interface{}{
			"id":       newUser.ID,
			"username": newUser.Username,
			"role":     newUser.Role,
		},
	})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, APIResponse{
			Success: false, Message: "format JSON salah",
		})
		return
	}

	var found *User
	for _, u := range users {
		if u.Username == req.Username {
			found = &u
			break
		}
	}

	if found == nil || !checkPassword(req.Password, found.Password) {
		writeJSON(w, http.StatusUnauthorized, APIResponse{
			Success: false, Message: "username atau password salah",
		})
		return
	}

	token, err := generateToken(*found)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, APIResponse{
			Success: false, Message: "gagal generate token",
		})
		return
	}

	writeJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Message: "login berhasil",
		Data: map[string]interface{}{
			"token":    token,
			"username": found.Username,
			"role":     found.Role,
		},
	})
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")

	var found *User
	for _, u := range users {
		if fmt.Sprintf("%d", u.ID) == userID {
			found = &u
			break
		}
	}

	if found == nil {
		writeJSON(w, http.StatusNotFound, APIResponse{
			Success: false, Message: "user tidak ditemukan",
		})
		return
	}

	writeJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Data: map[string]interface{}{
			"id":       found.ID,
			"username": found.Username,
			"role":     found.Role,
		},
	})
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Message: "selamat datang admin!",
	})
}

func verifyHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.URL.Query().Get("token")
	if tokenString == "" {
		writeJSON(w, http.StatusBadRequest, APIResponse{
			Success: false, Message: "parameter token diperlukan",
		})
		return
	}

	claims, err := validateToken(tokenString)
	if err != nil {
		writeJSON(w, http.StatusOK, APIResponse{
			Success: false,
			Message: fmt.Sprintf("token tidak valid: %v", err),
		})
		return
	}

	writeJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Message: "token valid",
		Data: map[string]interface{}{
			"user_id":  claims.UserID,
			"username": claims.Username,
			"role":     claims.Role,
			"expires":  claims.ExpiresAt.Time,
		},
	})
}

func main() {
	fmt.Println("BCRYPT — HASH PASSWORD")
	fmt.Println()

	password := "rahasia123"
	hash, err := hashPassword(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Password: %s\n", password)
	fmt.Printf("Hash:     %s\n", hash)
	fmt.Printf("Cocok?    %v\n", checkPassword(password, hash))
	fmt.Printf("Salah?    %v\n", checkPassword("passwordSalah", hash))

	fmt.Println()
	fmt.Println("JWT — GENERATE & VALIDATE TOKEN")
	fmt.Println()

	tokenString, err := generateToken(users[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Token: %s...\n", tokenString[:50])

	claims, err := validateToken(tokenString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("UserID:   %d\n", claims.UserID)
	fmt.Printf("Username: %s\n", claims.Username)
	fmt.Printf("Role:     %s\n", claims.Role)
	fmt.Printf("Expires:  %s\n", claims.ExpiresAt.Time)

	tamperedToken := tokenString[:len(tokenString)-5] + "xxxxx"
	_, err = validateToken(tamperedToken)
	fmt.Printf("Token rusak: %v\n", err)

	fmt.Println()
	fmt.Println("REST API — AUTH")
	fmt.Println()

	mux := http.NewServeMux()
	mux.HandleFunc("/register", registerHandler)
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/verify", verifyHandler)
	mux.HandleFunc("/profile", authMiddleware(profileHandler))
	mux.HandleFunc("/admin", adminOnly(adminHandler))

	fmt.Println("Server jalan di http://localhost:8080")
	fmt.Println()
	fmt.Println("ENDPOINTS:")
	fmt.Println("  POST /register   -> {\"username\":\"...\",\"password\":\"...\",\"role\":\"admin|user\"}")
	fmt.Println("  POST /login      -> {\"username\":\"admin\",\"password\":\"admin123\"}")
	fmt.Println("  GET  /verify?token=...  -> cek token")
	fmt.Println("  GET  /profile    -> Authorization: Bearer <token>")
	fmt.Println("  GET  /admin      -> Authorization: Bearer <token> (admin only)")
	fmt.Println()

	log.Fatal(http.ListenAndServe(":8080", mux))
}
