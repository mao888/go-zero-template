package models

import (
	"time"
)

type BusinessUser struct {
	BifUserBid            string    `gorm:"bif_user_bid"` // 企业用户星火通账户bid
	CreditCode            string    `gorm:"credit_code"`  // 企业统一信用代码
	CreateTime            time.Time `gorm:"create_time"`  // 创建时间
	Name                  string    `gorm:"name"`         // 企业名称
	FullName              string    `gorm:"full_name"`    // 企业全称
	ImageUrl              string    `gorm:"image_url"`    // 企业头像
	Dna721Account         string    `gorm:"dna721_account"`
	Dna721PrivateKey      string    `gorm:"dna721_private_key"`
	Dna721ContractAddress string    `gorm:"dna721_contract_address"`
}
