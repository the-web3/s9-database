package database

import (
	"time"

	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/log"
)

type Addresses struct {
	Guid        string    `gorm:"primaryKey; column:guid" json:"guid"`
	Address     string    `gorm:"column:address" json:"address"`
	AddressType string    `gorm:"column:address_type" json:"address_type"`
	PublicKey   string    `gorm:"column:public_key" json:"public_key"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (Addresses) TableName() string {
	return "addresses"
}

type AddressesView interface {
	QueryAddressesByUser(userId string) (*Addresses, error)
}

type AddressesDB interface {
	AddressesView

	StoreAddresses(address []Addresses) error
}

type addressesDB struct {
	gorm *gorm.DB
}

func NewAddressesDB(db *gorm.DB) AddressesDB {
	return &addressesDB{gorm: db}
}

func (a addressesDB) QueryAddressesByUser(userId string) (*Addresses, error) {
	//TODO implement me
	panic("implement me")
}

func (a addressesDB) StoreAddresses(addressList []Addresses) error {
	if err := a.gorm.Table("addresses").CreateInBatches(addressList, len(addressList)).Error; err != nil {
		log.Error("Failed to store addresses", "error", err)
	}
	return nil
}
