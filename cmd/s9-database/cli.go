package main

import (
	"fmt"

	"github.com/the-web3/s9-database/worker"
	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/log"
	"github.com/the-web3/s9-database/common/opio"
	"github.com/the-web3/s9-database/config"
	"github.com/the-web3/s9-database/database"
	flags2 "github.com/the-web3/s9-database/flags"
)

func runMigrations(ctx *cli.Context) error {
	ctx.Context = opio.CancelOnInterrupt(ctx.Context)
	log.Info("running migrations...")
	cfg := config.NewConfig(ctx)
	db, err := database.NewDB(ctx.Context, cfg.MasterDB)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		return err
	}
	defer func(db *database.DB) {
		err := db.Close()
		if err != nil {
			log.Error("fail to close database", "err", err)
		}
	}(db)
	return db.ExecuteSQLMigration("./migrations")
}

func runGenerateBatchAddress(ctx *cli.Context) error {
	cfg := config.NewConfig(ctx)
	db, err := database.NewDB(ctx.Context, cfg.MasterDB)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		return err
	}
	return worker.CreateBatchAddress(ctx, db)
}

func runQueryAddress(ctx *cli.Context) error {
	cfg := config.NewConfig(ctx)
	db, err := database.NewDB(ctx.Context, cfg.MasterDB)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		return err
	}
	fmt.Println("runQueryAddress starting...")
	addrList, err := worker.QueryBatchAddress("0x18AC127728e20e23019F5F5f739244d64dA6c4cD", db)
	if err != nil {
		log.Error("fail to query batch address", "err", err)
		return err
	}
	for _, addr := range addrList {
		log.Info("========= address ==========")
		log.Info(addr)
		log.Info("========= address =========")
	}
	fmt.Println("runQueryAddress end...")

	return nil
}

func NewCli() *cli.App {
	flags := flags2.Flags
	return &cli.App{
		Version:              "0.0.1",
		Description:          "An database operation system",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "migrate",
				Flags:       flags,
				Description: "Run database migrations",
				Action:      runMigrations,
			},
			{
				Name:        "create-address",
				Flags:       flags,
				Description: "Run create address",
				Action:      runGenerateBatchAddress,
			},
			{
				Name:        "query-address",
				Flags:       flags,
				Description: "Run query address",
				Action:      runQueryAddress,
			},
			{
				Name:        "version",
				Description: "Show project version",
				Action: func(ctx *cli.Context) error {
					cli.ShowVersion(ctx)
					return nil
				},
			},
		},
	}
}
