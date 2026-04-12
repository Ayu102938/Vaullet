# Vaullet

作成者(Author)
Ayu

## 📖 概要 (Overview)

Vaullet（ボレット）は、大学の部活・サークルにおける部費の立替申請や残高管理をスムーズに行うために開発されたデスクトップアプリケーションです。
Wailsフレームワークを採用し、軽量かつ高速なGo言語のバックエンドと、モダンでリッチなNext.js（React）のフロントエンドを組み合わせています。複雑な会計処理を、シンプルで直感的なUIで管理することを目指しています。

## 🛠 技術スタック (Tech Stack)

### Backend

- **言語**: Go
- **データベース**: SQLite (`modernc.org/sqlite`)
- **フレームワーク**: Wails

### Frontend

- **フレームワーク**: Next.js (App Router) / React
- **スタイリング**: Tailwind CSS
- **UIコンポーネント**: shadcn/ui (Radix UI)
- **アイコン**: lucide-react

## 📁 ディレクトリ構成 (Directory Structure)

```text
Vaullet/
├── app.go            # バックエンドの司令塔（フロントエンドとGoを繋ぐAPI）
├── database.go       # データベース(SQLite)の初期化・接続・テーブル定義
├── models.go         # データ構造（Transaction等）の定義
├── main.go           # アプリケーションのエントリーポイント
├── vaullet.db        # SQLiteデータベースファイル（自動生成）
└── frontend/         # フロントエンド（Next.js）の領土
    ├── app/
    │   ├── page.tsx  # メイン画面（ダッシュボード＆経費申請）
    │   └── globals.css
    ├── components/
    │   └── ui/       # shadcn/uiで追加したデザインコンポーネント群（button等）
    ├── lib/          # shadcn/ui等のユーティリティ関数
    ├── package.json  # フロントエンドのパッケージ管理
    └── wailsjs/      # Wailsが自動生成するGo関数のバインディング（型定義）
```

## 開発環境の立ち上げ (Getting Started)

**必須環境 (Prerequisites)**

- Go (1.20+)
- Node.js (18+)
- npm
- Wails CLI

## 起動コマンド

プロジェクトのルートディレクトリ（Vaullet/直下）で以下のコマンドを実行します。

```bash
wails dev
```

## 開発時の注意点

フロントエンドのライブラリ（npmパッケージやshadcn/uiのコンポーネント）を追加する場合は、必ず frontend ディレクトリに移動してからコマンドを実行してください。

```bash
cd frontend
npm install <package-name>
# または
npx shadcn@latest add <component-name>
```

## About

Wails template for Nextjs v15 with tailwindcss v4.

You can configure the project by editing `wails.json`. More information about the project settings can be found
here: https://wails.io/docs/reference/project-config

## New Project

You can create a new wails project using:

```
wails init -n "Your Project Name" -t https://github.com/kairo913/wails-nextjs-tailwind-template
```

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.
