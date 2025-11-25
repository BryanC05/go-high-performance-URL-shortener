# Gunakan image Go versi terbaru
FROM golang:alpine

# Set folder kerja di dalam container
WORKDIR /app

# Copy file go.mod dulu (biar cache layer docker jalan)
COPY go.mod ./

# Karena kita belum punya go.sum, kita skip copy go.sum dulu
# Nanti kita akan jalankan 'go mod tidy' di dalam container

# Copy sisa source code
COPY . .

# Download library yang dibutuhkan
RUN go get github.com/gofiber/fiber/v2
RUN go get github.com/redis/go-redis/v9
RUN go mod tidy

# Build aplikasi jadi file binary bernama 'main'
RUN go build -o main .

# Jalankan aplikasinya
CMD ["./main"]