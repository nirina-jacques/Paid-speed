package handler

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"io"
	"net/http"
	"payspeed/cmd/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type inputSigin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
}

func (cf *ServerConfig) Signin(w http.ResponseWriter, r *http.Request) {
	var input inputSigin
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
	}
	var user model.User
	if err := cf.Db.First(&user, "email = ?", input.Username).Error; err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	if CheckPasswordHash(input.Password, user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	token, err := getToken(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp := map[string]string{"token": token}
	jsonResponse, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func getToken(req *http.Request) (string, error) {
	body, _ := io.ReadAll(req.Body)
	var input inputSigin
	err := json.Unmarshal(body, &input)
	key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", err
	}
	expirationTime := time.Now().Add(120 * time.Minute)
	claims := &claims{
		Username: input.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signed, _ := token.SignedString(key)

	return signed, err
}

func (cf *ServerConfig) Register(w http.ResponseWriter, r *http.Request) {
	var input RegisterInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	hash, err := HashPassword(input.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user := model.User{
		ID:       uuid.New().String(),
		Email:    input.Email,
		Password: hash,
		FullName: input.FullName,
		Phone:    input.Phone,
	}
	if err := cf.Db.Create(&user).Error; err != nil {
		http.Error(w, "Email already used", http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusCreated)
	message := map[string]bool{"message": true}
	json.Marshal(message)
}
