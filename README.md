# 🐹 Soozan Notifier

> A fire-and-forget SMS notification microservice written in pure Go!

[![Go Version](https://img.shields.io/badge/Go-1.24.4-00ADD8?style=flat&logo=go)](https://golang.org)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker)](Dockerfile)

## 🚀 What is Soozan Notifier?

Soozan Notifier is a **fire-and-forget** SMS microservice designed for maximum simplicity and Go exploration.


## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   HTTP Client   │───▶│    Notifier     │───▶│  SMS Provider   │
│                 │    │                 │    │                 │
│  POST /new-task │    │  Goroutine      │    │  4 Tries (max)  │
│  GET  /health   │    │  Retry Logic    │    │  2 Providers    │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## 🚦 Quick Start

### Using Docker (Recommended)

```bash
# Build the image
docker build -t soozan-notifier .

# Run the container
docker run -p 8080:8080 soozan-notifier

# Create a new task
curl -X POST http://localhost:8080/new-task \
  -H "Content-Type: application/json" \
  -d '{"number":"09148387871", "scenario":"OTP", "params":["000000"]}'
```


