package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

// Struct untuk menerima input JSON
type ShortenRequest struct {
	URL string `json:"url"`
}

// Struct untuk respon JSON
type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func main() {
	// 1. Koneksi ke Redis (Perhatikan "redis:6379" adalah nama service di Docker)
	rdb = redis.NewClient(&redis.Options{
		Addr: "redis:6379", 
	})

	// Cek koneksi Redis
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Gagal konek ke Redis:", err)
	}

	// 2. Setup Fiber (Web Framework)
	app := fiber.New()

	// Endpoint 1: Memendekkan URL
	app.Post("/shorten", func(c *fiber.Ctx) error {
		body := new(ShortenRequest)
		if err := c.BodyParser(body); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		// Generate ID unik 6 karakter
		id := generateRandomString(6)

		// Simpan ke Redis (Key: ID, Value: URL Asli, Expire: 24 Jam)
		err := rdb.Set(ctx, id, body.URL, 24*time.Hour).Err()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan ke DB"})
		}

		// Balikin respon ke user
		return c.JSON(ShortenResponse{
			ShortURL: os.Getenv("APP_URL") + "/" + id,
		})
	})

	// Endpoint 2: Redirect (Akses URL Pendek)
	app.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		// Cari URL asli di Redis
		val, err := rdb.Get(ctx, id).Result()
		
		if err == redis.Nil {
			return c.Status(404).JSON(fiber.Map{"error": "Link tidak ditemukan"})
		} else if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Internal Error"})
		}

		// Redirect user ke URL asli
		return c.Redirect(val)
	})

	// Jalankan server di port 8080
	fmt.Println("Server running on port 8080")
	log.Fatal(app.Listen(":8080"))
}

// Fungsi helper untuk bikin random string
func generateRandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}