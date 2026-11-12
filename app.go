package main

import (
	"context"
	"database/sql"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
	db  *sql.DB // dbの接続を入れるやつらしい
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
	a.db = InitDB() //dbを初期化して、結果をAppのa.dbに保存する
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	if a.db != nil {
		// データベースの接続を閉じる
		a.db.Close()
		fmt.Print("データベースを閉じました")
	}
}

// AddEntry は家計簿の記録を1件追加し、作成されたIDを返します。
func (a *App) AddEntry(e Entry) (int64, error) {
	// 安全なSQL文の準備
	query := `
	INSERT INTO entries (
		kind, category, subcategory, title, memo, amount, account, date, receipt_path
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	// SQL文を実行し、？の部分に構造体のデータを当てはめる
	result, err := a.db.Exec(query, e.Kind, e.Category, e.Subcategory, e.Title, e.Memo, e.Amount, e.Account, e.Date, e.ReceiptPath)

	// エラー情報をフロントエンドに返す
	if err != nil {
		return 0, err
	}

	// 成功したら新しく作られたデータのIDを返す
	id, err := result.LastInsertId()
	return id, err
}

// GetEntries は家計簿の記録を日付の新しい順に取得します。
func (a *App) GetEntries() ([]Entry, error) {
	// データの取得
	query := `
	SELECT id, kind, category, subcategory, title, memo, amount, account, date, receipt_path, created_at, updated_at
	FROM entries
	ORDER BY date DESC, id DESC
	`
	rows, err := a.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// データを格納するためのリストを作成
	var entries []Entry
	for rows.Next() {
		var e Entry
		// subcategory / memo / receipt_path は NULL になりうるため sql.NullString で受け取る
		var subcategory, memo, receiptPath sql.NullString
		err := rows.Scan(
			&e.Id, &e.Kind, &e.Category, &subcategory, &e.Title, &memo, &e.Amount, &e.Account, &e.Date, &receiptPath, &e.CreatedAt, &e.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		e.Subcategory = subcategory.String
		e.Memo = memo.String
		e.ReceiptPath = receiptPath.String
		// データをリストに追加
		entries = append(entries, e)
	}
	return entries, nil
}
