package database

import (
	"math/big"
	"time"

	"gorm.io/gorm"
)

type Tokens struct {
	Guid           string    `gorm:"primaryKey;column:guid" json:"guid"`
	TokenAddress   string    `gorm:"column:token_address" json:"token_address"`
	Decimals       uint8     `gorm:"column:decimals" json:"decimals"`
	TokenName      string    `gorm:"column:token_name" json:"token_name"`
	TokenSymbol    string    `gorm:"column:token_symbol" json:"token_symbol"`
	CollectAmount  *big.Int  `gorm:"type:numeric;not null;default:0;check:collect_amount >= 0;serializer:u256" json:"collect_amount"`
	ColdAmount     *big.Int  `gorm:"type:numeric;not null;default:0;check:cold_amount >= 0;serializer:u256" json:"cold_amount"`
	HotAlertAmount *big.Int  `gorm:"type:numeric;not null;default:0;check:hot_alert_amount >= 0;serializer:u256" json:"hot_alert_amount"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Tokens) TableName() string {
	return "tokens"
}

type TokensView interface {
	QueryTokensList(string) ([]*Tokens, error)
}

type TokensDB interface {
	TokensView
	StoreTokens([]*Tokens) error
}

type tokensDB struct {
	gorm *gorm.DB
}

func NewTokensDB(db *gorm.DB) TokensDB {
	return &tokensDB{gorm: db}
}

func (t tokensDB) QueryTokensList(s string) ([]*Tokens, error) {
	//TODO implement me
	panic("implement me")
}

func (t tokensDB) StoreTokens(tokens []*Tokens) error {
	//TODO implement me
	panic("implement me")
}
