# Portfolio Backend

個人ポートフォリオ用のバックエンドシステム

## 技術スタック

- **言語**: Go 1.25
- **API**: GraphQL (gqlgen)
- **開発環境**: Docker Compose
- **ローカルDB**: MySQL 8.0
- **本番環境**: AWS ECS + Amazon Aurora
- **オーケストレーション**: Makefile

## ディレクトリ構成

```
portfolio-backend/
├── cmd/
│   ├── api/
│   │   └── main.go          # APIサーバーのエントリーポイント
│   └── worker/
│       └── main.go          # ワーカーのエントリーポイント
├── internal/
│   ├── graph/
│   │   ├── generated.go     # gqlgen生成コード（自動生成）
│   │   ├── models_gen.go    # gqlgen生成モデル（自動生成）
│   │   ├── resolver.go      # リゾルバー依存注入
│   │   ├── schema.graphqls  # GraphQLスキーマ定義
│   │   └── schema.resolvers.go  # リゾルバー実装
│   ├── service/
│   │   └── user_service.go  # ユーザービジネスロジック（API・ワーカー共通）
│   ├── repository/
│   │   └── user_repository.go  # リポジトリインターフェース（API・ワーカー共通）
│   ├── model/
│   │   └── user.go          # ドメインモデル
│   ├── infrastructure/
│   │   └── mysql/
│   │       └── user_repository.go  # MySQL実装
│   └── common/
│       └── env.go           # 共通ヘルパー関数
├── migrations/
│   └── 001_create_users_table.sql
├── .env.example             # 環境変数テンプレート
├── .gitignore
├── Dockerfile               # APIサーバー用Dockerイメージ
├── docker-compose.yml       # ローカル開発環境構成
├── gqlgen.yml               # gqlgen設定
├── Makefile                 # ローカル開発用コマンド
└── README.md
```

### 層の役割

| 層 | ディレクトリ | 説明 |
|---|---|---|
| エントリーポイント | `cmd/api`, `cmd/worker` | API・ワーカー各サーバーの起動 |
| GraphQL | `internal/graph` | スキーマ・リゾルバー実装 |
| サービス | `internal/service` | ビジネスロジック（API・ワーカー共通） |
| リポジトリ | `internal/repository` | DBアクセスインターフェース（共通） |
| モデル | `internal/model` | ドメインモデル定義 |
| インフラ | `internal/infrastructure/mysql` | MySQL実装 |
| 共通 | `internal/common` | 汎用ヘルパー関数 |

## セットアップ

### 前提条件

- Docker & Docker Compose がインストールされていること
- Make がインストールされていること

### 初回セットアップ

```bash
cd portfolio-backend

# .env ファイルを作成（テンプレートをコピー）
cp .env.example .env

# コンテナを起動してデータベースをセットアップ
make up
```

`make up` は初回実行時に `.env` が存在しない場合、`.env.example` から自動的にコピーします。

これにより以下が自動実行されます：
1. MySQL と Go アプリケーションコンテナが起動
2. `common_db` データベースが作成
3. `users` テーブルが作成
4. アプリケーションが `http://localhost:8080` で起動

### WSL (Windows) でのセットアップ

Windows の WSL2 Ubuntu 環境でセットアップする場合：

```bash
# practice リポジトリをクローン
git clone https://github.com/kojima1128/practice.git
cd practice/portfolio-backend

# セットアップスクリプトを実行（WSL Ubuntu ターミナルで）
chmod +x setup-wsl.sh
./setup-wsl.sh
```

詳細は [WSL-SETUP.md](../WSL-SETUP.md) を参照してください。

## 環境変数

`.env.example` を参考に `.env` を作成してください（`.env` はGit管理外）。

| 変数名 | デフォルト値 | 説明 |
|---|---|---|
| `PORT` | `8080` | APIサーバーのポート |
| `DB_HOST` | `db` | DBホスト名 |
| `DB_PORT` | `3306` | DBポート |
| `DB_USER` | `user` | DBユーザー名 |
| `DB_PASSWORD` | `password` | DBパスワード |
| `DB_NAME` | `common_db` | DB名 |
| `WORKER_INTERVAL` | `60s` | ワーカー実行間隔 |

## コマンド

| コマンド | 説明 |
|---|---|
| `make up` | コンテナ起動・マイグレーション実行 |
| `make down` | コンテナ停止・ボリューム削除 |
| `make stop` | コンテナ停止（ボリューム保持） |
| `make logs` | コンテナログ表示 |
| `make ps` | コンテナ状態確認 |
| `make build` | APIバイナリをローカルビルド |
| `make generate` | GraphQLコード再生成 |

## GraphQL API

### アクセス方法

- **URL**: `http://localhost:8080`
- **GraphQL エンドポイント**: `http://localhost:8080/query`
- **Playground**: `http://localhost:8080/`

### クエリ例

#### ユーザー一覧取得

```graphql
query {
  users {
    id
    name
    email
    createdAt
    updatedAt
  }
}
```

#### ユーザー作成

```graphql
mutation {
  createUser(input: {
    name: "山田太郎"
    email: "yamada@example.com"
  }) {
    id
    name
    email
    createdAt
    updatedAt
  }
}
```

#### ユーザー取得

```graphql
query {
  user(id: "1") {
    id
    name
    email
    createdAt
    updatedAt
  }
}
```

## データベース

### ローカル開発環境（MySQL）

- **ホスト**: `localhost`
- **ポート**: `3306`
- **ユーザー**: `user`
- **パスワード**: `password`
- **データベース名**: `common_db`

**MySQL へのアクセス例**:

```bash
mysql -h localhost -u user -p -D common_db
# パスワード: password
```

## マイグレーション

マイグレーション SQL ファイルは `migrations/` ディレクトリに配置します。

- `001_create_users_table.sql` - Users テーブル初期化

新しいマイグレーションを追加する場合：
1. `migrations/NNN_description.sql` という命名規則でファイルを作成
2. `make down` を実行してボリュームを削除
3. `make up` を実行してマイグレーションを再実行

## GraphQL コード生成

スキーマを変更した場合は `make generate` でコードを再生成してください。

```bash
# スキーマ変更後
make generate
```

## 次のステップ

- [ ] MySQL リポジトリの実装 (`internal/infrastructure/mysql/`)
- [ ] ワーカージョブの実装 (`cmd/worker/`)
- [ ] ユーザー認証機能
- [ ] エラーハンドリング
- [ ] ロギング機構
- [ ] ユニットテスト
- [ ] 本番環境へのデプロイメント（ECS + Aurora）

## ライセンス

MIT
