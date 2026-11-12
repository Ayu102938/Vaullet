# Vaullet

作成者(Author): Ayu / Ayuki

個人用の家計簿管理アプリケーション。Wails (Go) + Next.js で作られたデスクトップアプリです。

## 📖 概要 (Overview)

Vaullet は、個人の収支をローカルの SQLite に記録して管理する家計簿アプリです。
Go 言語の軽量なバックエンドと、Next.js (React) のモダンなフロントエンドを Wails フレームワークで組み合わせ、単一のデスクトップアプリとして提供します。

データはすべてローカル（`vaullet.db`）に保存され、外部サービスやネットワークに送信しません。

## 🛠 技術スタック (Tech Stack)

### Backend

- **言語**: Go
- **データベース**: SQLite (`modernc.org/sqlite`, 純Goドライバ)
- **フレームワーク**: Wails v2

### Frontend

- **フレームワーク**: Next.js (App Router) / React
- **スタイリング**: Tailwind CSS v4
- **UIコンポーネント**: shadcn/ui (Radix UI)
- **アイコン**: lucide-react

## 📁 ディレクトリ構成 (Directory Structure)

```text
Vaullet/
├── app.go            # バックエンドのAPI（フロントエンドとGoを繋ぐBinding）
├── database.go       # SQLiteの初期化・テーブル定義
├── models.go         # データ構造（Entry）の定義
├── main.go           # アプリケーションのエントリーポイント
├── wails.json        # Wailsプロジェクト設定
├── scripts/          # ビルド・Wails CLI インストール用の補助スクリプト
└── frontend/         # フロントエンド（Next.js）
    ├── src/app/      # ページ（page.tsx / layout.tsx / globals.css）
    ├── src/components/ui/  # shadcn/ui コンポーネント
    ├── src/lib/      # ユーティリティ関数
    ├── wailsjs/      # Wailsが自動生成するGo関数のバインディング（触らない）
    └── package.json
```

`vaullet.db` は `.gitignore` 対象です。初回の起動時に `InitDB()` が自動生成します。

## 💾 データモデル

家計簿の1記録は `entries` テーブルに保存されます。

| カラム | 型 | 説明 |
| --- | --- | --- |
| `id` | INTEGER | 自動採番のID |
| `kind` | TEXT | `income`（収入） / `expense`（支出） |
| `category` | TEXT | 大分類（食費 / 住居 / 交通 など） |
| `subcategory` | TEXT | 細分類（任意） |
| `title` | TEXT | 項目名 |
| `memo` | TEXT | 備考（任意） |
| `amount` | REAL | **金額（円・小数可）** |
| `account` | TEXT | 資金源（`cash` / `card` / `savings` など） |
| `date` | TEXT | 日付（`YYYY-MM-DD`） |
| `receipt_path` | TEXT | 領収書の保存パス（任意） |
| `created_at` / `updated_at` | TEXT | 作成・更新時刻 |

金額は小数（`REAL` / Go の `float64`）として保存します。表示は原則として小数2桁まで丸めて扱います。

## 🔌 Backend API（Wails Binding）

| 関数 | 説明 |
| --- | --- |
| `AddEntry(e Entry) (int64, error)` | 記録を1件追加し、IDを返す |
| `GetEntries() ([]Entry, error)` | 記録を日付の新しい順に取得する |

編集・削除、月別集計などの API は今後の実装予定です。

## 🚀 開発環境の立ち上げ (Getting Started)

**必須環境 (Prerequisites)**

- Go 1.25 以上
- Node.js 18 以上 / npm
- Wails CLI (`go install github.com/wailsapp/wails/v2/cmd/wails@latest` でインストール)

**起動**

プロジェクトルート（`Vaullet/` 直下）で実行します。

```bash
wails dev
```

開発サーバーは `http://localhost:3000`（Next.js）を使用します。ブラウザで開発したい場合は Wails の dev サーバー（`http://localhost:34115`）から Go の関数を直接呼び出せます。

**ビルド**

```bash
wails build
```

`main.go` は `frontend/out` を embed するため、`go build` だけで通すには先に `cd frontend && npm install && npm run build` が必要です（`wails build` は自動的に実行します）。

## ⚠️ 開発時の注意点

npm パッケージや shadcn/ui のコンポーネントを追加する場合は、必ず `frontend` ディレクトリで実行します。

```bash
cd frontend
npm install <package-name>
# または
npx shadcn@latest add <component-name>
```

## 🗺 Roadmap

- [ ] フロントエンドの実装（ダッシュボード / 記録一覧 / 追加ダイアログ）
- [ ] 記録の編集・削除 API
- [ ] 月別・カテゴリ別の集計
- [ ] カテゴリ・アカウントの管理
- [ ] データ保存先の移動（`~/.local/share/vaullet/` など XDG基準）
- [ ] ライセンス (LICENSE) と CI の整備

## 🧞 Wails テンプレートについて

本プロジェクトは Next.js v15 + Tailwind CSS v4 の Wails テンプレートから始まりました。
プロジェクト設定は `wails.json` で変更できます（参考: https://wails.io/docs/reference/project-config ）。

同テンプレートは次のコマンドで新規作成できます。

```bash
wails init -n "Your Project Name" -t https://github.com/kairo913/wails-nextjs-tailwind-template
```
