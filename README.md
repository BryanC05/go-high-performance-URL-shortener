# ⚡️ High-Performance URL Shortener (Go + Redis)

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=flat&logo=go&logoColor=white)
![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?style=flat&logo=redis&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=flat&logo=docker&logoColor=white)

A specialized microservice designed for **extreme low-latency** URL redirection. Built with Golang (Fiber) and Redis to handle thousands of concurrent requests with millisecond response times.

---

## 🚀 Performance Benchmark

I conducted a stress test using **k6** to compare this architecture against a traditional SQL-based backend.

| Metric | Laravel + MySQL | Go + Redis (This Project) | Improvement |
| :--- | :--- | :--- | :--- |
| **Concurrency** | 50 VUs | **100 VUs** | 2x Load |
| **Response Time** | ~200ms | **< 5ms** | **40x Faster** |
| **Throughput** | ~2,400 RPS | **~10,000+ RPS** | **Massive Scale** |

> **Conclusion:** By using an in-memory database (Redis) and a compiled language (Go), we achieved significant performance gains for read-heavy operations.

---

## 🛠 Tech Stack
* **Language:** Go 1.21 (Fiber Framework)
* **Database:** Redis (Alpine Image)
* **DevOps:** Docker Compose
* **Testing:** k6 (Load Testing)

## 📦 How to Run
```bash
# Start Services
docker-compose up -d --build

# Generate Short URL
curl -X POST http://localhost:8080/shorten \
-H "Content-Type: application/json" \
-d '{"url": "[https://google.com](https://google.com)"}'

# Run Load Test
k6 run load-test-go.js
