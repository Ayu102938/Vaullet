package main

type Transaction struct {
	Id                   int    `json:"id"`
	Product              string `json:"product"`
	ApplicantName        string `json:"applicant_name"` //申請者名
	Date                 string `json:"date"`
	Purpose              string `json:"purpose"`     //使用用途
	Explanation          string `json:"explanation"` //備考欄
	Price                int    `json:"price"`
	Expenditure          string `json:"expenditure"`            //支出元、プール金or学友会費
	Category             string `json:"category"`               //カテゴリ
	Status               string `json:"status"`                 //ステータス
	RejectedReason       string `json:"rejected_reason"`        //否決理由
	Receipt              string `json:"receipt"`                //領収書
	IsAdvanceApplication string `json:"is_advance_application"` // 先払い申請
}
