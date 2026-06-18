# Portfolio Backend

Go 1.24 + GraphQL + MySQL のバックエンドプロジェクトです。

## ディレクトリ構成

```text
portfolio-backend/
├── cmd/
│   ├── api/
│   │   └── main.go
│   └── worker/
│       └── main.go
├── internal/
│   ├── graph/
│   ├── service/
│   ├── repository/
│   ├── model/
│   ├── infrastructure/
│   │   └── mysql/
│   └── common/
├── migrations/
├── .env.example
├── .gitignore
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md
```
