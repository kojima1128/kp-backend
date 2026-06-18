# Portfolio Backend

個人ポートフォリオ用のバックエンドシステム

## 技術スタック

- **言語**: Go 1.21
- **API**: GraphQL
- **開発環境**: Docker Compose
- **ローカルDB**: MySQL 8.0
- **本番環境**: AWS ECS + Amazon Aurora
- **オーケストレーション**: Makefile

## プロジェクト構成

```
portfolio-backend/
├── main.go                    # アプリケーションエントリーポイント
├── go.mod                     # Go モジュール定義
├── Dockerfile                 # Docker イメージ定義
├── docker-compose.yml         # 開発環境構成
├── Makefile                   # ローカル開発用コマンド
├── migrations/                # データベースマイグレーション
│   └── 001_create_users_table.sql
└── graph/                     # GraphQL関連ファイル
    ├── schema.graphqls        # GraphQL スキーマ定義
    ├── models.go              # データモデル
    └── resolver.go            # GraphQL リゾルバー実装
```

## セットアップ

### 前提条件

- Docker & Docker Compose がインストールされていること
- Make がインストールされていること

### 初回セットアップ

```bash
cd portfolio-backend

# コンテナを起動してデータベースをセットアップ
make up
```

これにより以下が自動実行されます：
1. MySQL と Go アプリケーションコンテナが起動
2. `common_db` データベースが作成
3. `users` テーブルが作成
4. アプリケーションが `http://localhost:8080` で起動

## コマンド

### コンテナ起動（マイグレーション実行）

```bash
make up
```

### コンテナ停止（ボリューム削除）

```bash
make down
```

### コンテナ停止（ボリューム保持）

```bash
make stop
```

### ログ表示

```bash
make logs
```

### コンテナ状態確認

```bash
make ps
```

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

### 本番環境（Amazon Aurora MySQL）

本番環境では AWS RDS の Aurora MySQL を使用します。
接続情報は環境変数で設定します。

## マイグレーション

マイグレーション SQL ファイルは `migrations/` ディレクトリに配置します。

- `001_create_users_table.sql` - Users テーブル初期化

新しいマイグレーションを追加する場合：
1. `migrations/NNN_description.sql` という命名規則でファイルを作成
2. `make down` を実行してボリュームを削除
3. `make up` を実行してマイグレーションを再実行

## トラブルシューティング

### コンテナが起動しない場合

```bash
# ログ確認
make logs

# コンテナ強制削除
docker-compose down -v
make up
```

### データベース接続エラー

```bash
# MySQL が起動するまで待機時間を増やす
# docker-compose.yml の healthcheck の retries を調整
```

## 次のステップ

- [ ] 実際のデータベース接続実装
- [ ] GraphQL リゾルバーの実装
- [ ] ユーザー認証機能
- [ ] エラーハンドリング
- [ ] ロギング機構
- [ ] ユニットテスト
- [ ] インテグレーションテスト
- [ ] 本番環境へのデプロイメント（ECS + Aurora）

## ライセンス

MIT
