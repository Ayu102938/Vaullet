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

	//下からテスト用
	dummyData := Transaction{
		Id:                   1,
		Product:              "テスト商品",
		ApplicantName:        "テスト太郎",
		Date:                 "2024-01-01",
		Purpose:              "テスト用途",
		Explanation:          "テスト備考",
		Price:                2500,
		Expenditure:          "テスト支出元",
		Category:             "テストカテゴリ",
		Status:               "テストステータス",
		RejectedReason:       "テスト拒否理由",
		Receipt:              "テスト領収書",
		IsAdvanceApplication: "テスト先払い申請",
	}
	//CreateTransaction関数を呼び出してみる
	id, err := a.CreateTransaction(dummyData)
	if err != nil {
		fmt.Println("テスト失敗:", err)
	} else {
		fmt.Println("テスト成功:", id)
	}
	// ここまで
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

// Greet returns a greeting for the given name
func (a *App) CreateTransaction(t Transaction) (int64, error) {
	// 安全なSQL文の準備
	query := `
	INSERT INTO transactions (
		id, product, applicant_name, date, purpose, explanation, price, expenditure, category, status, rejected_reason, receipt, is_advance_application
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	// SQL文を実行し、？の部分に構造体のデータを当てはめる
	result, err := a.db.Exec(query, t.Id, t.Product, t.ApplicantName, t.Date, t.Purpose, t.Explanation, t.Price, t.Expenditure, t.Category, t.Status, t.RejectedReason, t.Receipt, t.IsAdvanceApplication)

	// エラー情報をフロントエンドに返す
	if err != nil {
		return 0, err
	}

	//成功したら新しく作られたデータのIDを返す
	id, err := result.LastInsertId()
	return id, err
}
