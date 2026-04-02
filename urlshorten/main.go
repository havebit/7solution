package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pallat/urlshorten/ginrouter"
	"github.com/pallat/urlshorten/shorten"
	"github.com/pallat/urlshorten/sqlite"
)

var mySigningKey = []byte("AllYourBase")

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}

	db, err := sql.Open("sqlite3", "./urlshortener.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		db.Close()
	}()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// mux := http.NewServeMux()
	r := gin.Default()

	// mux.HandleFunc("/login", loginHandler)

	storage := sqlite.NewStorage(db)
	handler := shorten.NewHandler(storage)
	r.POST("/shorten", ginrouter.NewHandler(handler.Handler))

	redirectHandler := shorten.NewRedirectHandler(storage)
	r.GET("/:shorturl", ginrouter.NewHandler(redirectHandler.Handler))

	srv := &http.Server{
		Handler: r,
		Addr:    ":" + os.Getenv("PORT"),
	}

	log.Printf("api serving on %s...\n", os.Getenv("PORT"))
	srv.ListenAndServe()
}

type SQLStorage struct {
	db *sql.DB
}

func (s *SQLStorage) Save(shortenURL, originalURL string) error {
	_, err := s.db.Exec("INSERT INTO urls (key, original_url) VALUES (?, ?)", shortenURL, originalURL)
	return err
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": ss,
	})
}

func authenMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authen := r.Header.Get("Authentication")
		tokenString := authen[7:]

		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return mySigningKey, nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		handler(w, r)
	}
}
