package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var mySigningKey = []byte("AllYourBase")

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}

	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := client.Disconnect(ctx); err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
	}()

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/login", loginHandler)

	handler := newShortenHandler(client)
	mux.HandleFunc("/shorten", handler.ServeHTTP)

	redirectHandler := &redirectHandler{client: client}
	mux.HandleFunc("/{shorturl}", redirectHandler.ServeHTTP)

	srv := &http.Server{
		Handler: mux,
		Addr:    ":" + os.Getenv("PORT"),
	}

	log.Printf("api serving on %s...\n", os.Getenv("PORT"))
	srv.ListenAndServe()
}

type URLBody struct {
	URL string `json:"url" binding:"required"`
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type shortenHandler struct {
	client *mongo.Client
}

func newShortenHandler(client *mongo.Client) *shortenHandler {
	return &shortenHandler{client: client}
}

func (handler *shortenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var payload URLBody

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	const size = 6
	var shartenKey [size]byte
	for i := range size {
		shartenKey[i] = charset[rand.Intn(len(charset))]
	}

	collection := handler.client.Database("urls").Collection("shorten")
	_, err := collection.InsertOne(context.TODO(), bson.M{"key": string(shartenKey[:]), "original_url": payload.URL})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"short_url": fmt.Sprintf("http://localhost:8080/%s", string(shartenKey[:])),
	})
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

type redirectHandler struct {
	client *mongo.Client
}

func (handler *redirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("shorturl")

	collection := handler.client.Database("urls").Collection("shorten")

	var originalURL string

	var result struct {
		OriginalURL string `bson:"original_url"`
	}
	filter := bson.M{"key": key}

	if err := collection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}
