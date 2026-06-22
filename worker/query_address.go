package worker

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/the-web3/s9-database/database"
)

func QueryBatchAddress(address string, db *database.DB) ([]string, error) {
	var addrList []string
	addressList, err := db.Balances.QueryBalancesByAddress()
	if err != nil {
		log.Error("Failed to query balances by address", "address", address, "err", err)
		return nil, err
	}
	for _, addressItem := range addressList {
		addrList = append(addrList, addressItem.Address)
	}
	return addrList, nil
}
