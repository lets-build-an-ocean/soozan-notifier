# 🐹 Soozan Notifier

> A fire-and-forget SMS notification microservice written in pure Go!

[![Go Version](https://img.shields.io/badge/Go-1.24.4-00ADD8?style=flat&logo=go)](https://golang.org)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker)](Dockerfile)

## 🚀 What is Soozan Notifier?

Soozan Notifier is a **fire-and-forget** SMS microservice designed for maximum simplicity and Go exploration.


## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   HTTP Client   │───▶│  Soozan Server  │───▶│  SMS Provider   │
│                 │    │                 │    │  (Simulated)    │
│  POST /new-task │    │  Async Goroutines│    │  20% Fail Rate  │
│  GET  /health   │    │  Retry Logic    │    │  200ms Delay    │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## 🚦 Quick Start

### Using Docker (Recommended)

```bash
# Build the image
docker build -t soozan-notifier .

# Run the container
docker run -p 8080:8080 soozan-notifier
```


