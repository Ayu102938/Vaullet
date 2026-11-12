package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

// InitDB はデータベースを初期化して返します。
// ファイル(vaullet.db)はリポジトリにコミットせず、ローカルに自動生成されます。
func InitDB() *sql.DB {
	db, err := sql.Open("sqlite", "vaullet.db")
	if err != nil {
		log.Fatal("データベースの接続に失敗しました: ", err)
	}

	// 家計簿の1記録 = entries
	// amount は REAL(小数)。円単位で小数2桁までを想定しています。
	query := `
		CREATE TABLE IF NOT EXISTS entries (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			kind TEXT NOT NULL DEFAULT 'expense',
			category TEXT NOT NULL,
			subcategory TEXT,
			title TEXT NOT NULL,
			memo TEXT,
			amount REAL NOT NULL,
			account TEXT NOT NULL DEFAULT 'cash',
			date TEXT NOT NULL,
			receipt_path TEXT,
			created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
		CREATE INDEX IF NOT EXISTS idx_entries_date ON entries(date);
		CREATE INDEX IF NOT EXISTS idx_entries_category ON entries(category);
	`

	// SQL文をdbに送信して実行
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("テーブルの作成に失敗しました: ", err)
	}
	log.Println("データベースの初期化が完了しました")
	return db
}
