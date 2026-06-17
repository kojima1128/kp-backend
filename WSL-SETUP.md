# WSL Setup Guide

このガイドは、Windows 上の WSL2 Ubuntu で Portfolio Backend を開発するための初期セットアップを説明します。

## 前提条件

- Windows 10/11 に WSL2 がインストールされていること
- Ubuntu がインストールされていること（WSL2 で利用可能）
- VSCode がインストールされていること

## クイックセットアップ

### 1. practice リポジトリをクローン

```bash
# Windows PowerShell or CMD で実行
git clone https://github.com/kojima1128/practice.git
cd practice
```

### 2. WSL Ubuntu ターミナルを開く

- VSCode を開く
- `Ctrl + Shift + P` で コマンドパレットを開く
- "WSL: New WSL Window" を選択
- または WSL Ubuntu のターミナルを直接開く

### 3. setup-wsl.sh を実行

```bash
# practice ディレクトリで実行
chmod +x setup-wsl.sh
./setup-wsl.sh
```

スクリプトが以下を自動的に行います：

- ✅ Docker をインストール
- ✅ Docker Compose をインストール
- ✅ Make をインストール
- ✅ Go をインストール
- ✅ MySQL Client をインストール
- ✅ Docker デーモンを起動
- ✅ Go 依存関係をダウンロード
- ✅ コンテナを起動してデータベースマイグレーションを実行

## セットアップ後

### VSCode で開く

```bash
code portfolio-backend
```

### GraphQL Playground にアクセス

ブラウザで以下を開きます：
```
http://localhost:8080
```

### よく使うコマンド

```bash
# コンテナ起動（マイグレーション自動実行）
make up

# コンテナ停止（ボリューム削除）
make down

# コンテナ停止（ボリューム保持）
make stop

# ログ表示
make logs

# コンテナ状態確認
make ps
```

## Docker デーモンが起動しない場合

セットアップ後、以下のエラーが出る場合があります：

```
Cannot connect to the Docker daemon
```

**解決方法：**

```bash
# Docker デーモンを手動で起動
sudo service docker start

# または systemctl を使用
sudo systemctl start docker
```

## sudo パスワードが不要にするには

セットアップ後、Docker コマンドで sudo が必要な場合：

```bash
# Docker グループにユーザーを追加
sudo usermod -aG docker $USER
newgrp docker

# 確認
docker ps
```

## Windows → WSL ファイルアクセス

Windows 上のファイルから WSL にアクセスする場合：

```bash
# Windows のパス例
C:\Users\YourName\practice

# WSL 内でのパス
/mnt/c/Users/YourName/practice
```

## VSCode Remote WSL 拡張（推奨）

### インストール

1. VSCode を開く
2. 拡張機能を開く（`Ctrl + Shift + X`）
3. "Remote - WSL" を検索してインストール

### 使用方法

```bash
# WSL ターミナルで
code .

# または VSCode のリモートボタン（左下）から "WSL にて再度開く" を選択
```

メリット：
- Windows の VSCode UI を使用
- WSL 内の環境で開発
- ネイティブなパフォーマンス

## トラブルシューティング

### setup-wsl.sh の実行権限がない

```bash
chmod +x setup-wsl.sh
```

### Docker グループ設定後、反映されない

```bash
# 新しいグループを適用
newgrp docker

# または ターミナルを再起動
```

### Git の改行コード設定（推奨）

Windows での Git チェックアウト時に改行コードの問題が起きないようにします：

```bash
git config --global core.autocrlf input
```

### ホストの Windows から WSL のファイルを編集する場合

Windows ターミナルで：

```bash
# practice をスペースなしのパスにクローン
git clone https://github.com/kojima1128/practice.git practice
cd practice

# WSL で開く
wsl

# 次に setup-wsl.sh を実行
```

## 次のステップ

セットアップ完了後：

1. `portfolio-backend/` でコード開発開始
2. GraphQL スキーマを拡張
3. データベースマイグレーションを追加
4. リゾルバーを実装
5. テストを追加
6. 本番環境へのデプロイメント準備

## その他のリソース

- [WSL 公式ドキュメント](https://learn.microsoft.com/ja-jp/windows/wsl/)
- [Docker Desktop WSL 統合](https://docs.docker.com/desktop/wsl/)
- [VSCode Remote Development](https://code.visualstudio.com/docs/remote/remote-overview)
