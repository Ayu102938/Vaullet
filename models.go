package main

// Entry は家計簿の1記録を表します。
// Amount は小数（円）です。Kind で収入/支出を区別し、Amount は常に正値を持ちます。
type Entry struct {
	Id          int     `json:"id"`
	Kind        string  `json:"kind"`         // income | expense
	Category    string  `json:"category"`     // 食費 / 住居 / 交通 ...
	Subcategory string  `json:"subcategory"`  // 細分類（任意）
	Title       string  `json:"title"`        // 項目名
	Memo        string  `json:"memo"`         // 備考
	Amount      float64 `json:"amount"`       // 金額（円、小数可）
	Account     string  `json:"account"`      // cash / card / savings ...
	Date        string  `json:"date"`         // YYYY-MM-DD
	ReceiptPath string  `json:"receipt_path"` // 領収書の保存パス
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}
