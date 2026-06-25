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
├── .env-sample
├── .gitignore
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md
```

## セットアップ

### 環境変数

`.env-sample` をコピーして `.env` を作成してください。

```bash
cp .env-sample .env
```

その後、必要に応じて `.env` の値を編集してください。
