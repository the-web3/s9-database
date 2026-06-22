package database

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"gorm.io/gorm"
)

type Balances struct {
	Guid         string    `gorm:"primaryKey; column:guid"`
	Address      string    `gorm:"column:address"`
	TokenAddress string    `gorm:"column:token_address"`
	Balance      *big.Int  `gorm:"type:numeric;not null;default:0;check:balance >= 0;serializer:u256" json:"balance"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

func (Balances) TableName() string {
	return "balances"
}

type BalancesView interface {
	QueryBalancesByAddress() ([]*Balances, error)
	QueryBalancesByAddressAndToken(address string, token string) (*Balances, error)
}

type BalancesDB interface {
	BalancesView

	StoreBalances(balances []Balances) error
}

type balancesDB struct {
	gorm *gorm.DB
}

func NewBalancesDB(db *gorm.DB) BalancesDB {
	return &balancesDB{gorm: db}
}

func (b balancesDB) QueryBalancesByAddress() ([]*Balances, error) {
	var balancesList []*Balances
	if err := b.gorm.Table("balances").Find(&balancesList).Error; err != nil {
		log.Error("Failed to query balances", "error", err)
		return nil, err
	}
	return balancesList, nil
}

func (b balancesDB) QueryBalancesByAddressAndToken(address string, token string) (*Balances, error) {
	var balanceItem Balances
	if err := b.gorm.Table("balances").Where("address = ? and token_address", address, token).Take(&balanceItem).Error; err != nil {
		log.Error("Failed to query balances", "error", err)
		return nil, err
	}
	return &balanceItem, nil
}

func (b balancesDB) StoreBalances(balances []Balances) error {
	if err := b.gorm.Table("balances").CreateInBatches(balances, len(balances)).Error; err != nil {
		log.Error("failed to store balances", "error", err)
	}
	return nil
}
