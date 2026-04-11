package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

// InitDB はデータベースを初期化して返します。
func InitDB() *sql.DB {
	db, err := sql.Open("sqlite", "vaullet.db")
	if err != nil {
		log.Fatal("データベースの接続に失敗しました: ", err)
	}
	query := `
		CREATE TABLE IF NOT EXISTS transactions (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			product TEXT,
			applicant_name TEXT,
			date TEXT,
			purpose TEXT,
			explanation TEXT,
			price INTEGER,
			expenditure TEXT,
			category TEXT,
			status TEXT,
			rejected_reason TEXT,
			receipt TEXT,
			is_advance_application TEXT
		);
	`

	// SQL文をdbに送信して実行
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("テーブルの作成に失敗しました: ", err)
	}
	log.Println("データベースの初期化が完了しました")
	return db
}
