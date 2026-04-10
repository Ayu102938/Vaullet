type Transaction struct {
	ID              int    `json:"id"`
	PRODUCT         string `json:"product"`
	APPLICANT_NAME  string `json:"applicant_name"` //申請者名
	DATE            string `json:"date"`
	PURPOSE         string `json:"purpose"`     //使用用途
	EXPLANATION     string `json:"explanation"` //備考欄
	PRICE           int    `json:"price"`
	EXPENDITURE     string `json:"expenditure"`     //支出元、プール金or学友会費
	CATEGORY        string `json:"category"`        //カテゴリ
	STATUS          string `json:"status"`          //ステータス
	REJECTED_REASON string `json:"rejected_reason"` //否決理由
}
