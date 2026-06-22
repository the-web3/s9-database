package worker

import (
	"math/big"
	"time"

	"github.com/google/uuid"
	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/log"

	"github.com/the-web3/s9-database/database"
	"github.com/the-web3/s9-database/ethereum"
)

func CreateBatchAddress(ctx *cli.Context, db *database.DB) error {
	var addressList []database.Addresses
	var balanceList []database.Balances
	for index := 0; index < 100; index++ {
		addressStruct, err := ethereum.CreateAddressByKeyPairs()
		if err != nil {
			log.Error("create address error", err)
			return err
		}
		var AddressType string
		if index == 1 {
			AddressType = "hot"
		} else if index == 2 {
			AddressType = "cold"
		} else {
			AddressType = "user"
		}
		addressItem := database.Addresses{
			Guid:        uuid.New().String(),
			Address:     addressStruct.Address,
			AddressType: AddressType,
			PublicKey:   addressStruct.PublicKey,
			CreatedAt:   time.Now(),
		}
		addressList = append(addressList, addressItem)

		balanceItem := database.Balances{
			Guid:         uuid.New().String(),
			Address:      addressStruct.Address,
			TokenAddress: addressStruct.Address,
			Balance:      big.NewInt(0),
			CreatedAt:    time.Now(),
		}
		balanceList = append(balanceList, balanceItem)
	}
	err := db.Addresses.StoreAddresses(addressList)
	if err != nil {
		log.Error("store address error", err)
		return err
	}
	err = db.Balances.StoreBalances(balanceList)
	if err != nil {
		log.Error("store balances error", err)
		return err
	}
	return nil
}
