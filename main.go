package main

import (
	"anti-validators/cacher"
	"anti-validators/config"
	"anti-validators/db/models"
	"anti-validators/server"
	"anti-validators/signer"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	defaultConfigFileName = "config.json"
	defaultHost           = "127.0.0.1:8080"
)

func openDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "anti.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Request{})
	db.AutoMigrate(&models.Signs{})
	return db
}

func main() {
	db := openDb()
	defer db.Close()

	ctx := context.Background()

	var host, confFileName, oracleAddress string
	flag.StringVar(&confFileName, "config", defaultConfigFileName, "set host")
	flag.StringVar(&oracleAddress, "oracleAddress", "", "set oracle address")
	flag.StringVar(&host, "host", defaultHost, "set host")
	flag.Parse()

	cfg, err := config.Load(confFileName)
	if err != nil {
		panic(err)
	}

	//---------

	go server.StartServer(host, oracleAddress, db)

	if err := cacher.StartEthCacher(cfg.Ethereum.ContractAddress, cfg.Ethereum.NodeURL, ctx, db); err != nil {
		panic(err)
	}
	cacher.StartWavesCacher(cfg.Waves.ContractAddress, cfg.Waves.NodeURL, ctx, db)

	if err := signer.StartWavesSigner(cfg, oracleAddress, ctx, db); err != nil {
		panic(err)
	}
	if err := signer.StartEthSigner(cfg, oracleAddress, ctx, db); err != nil {
		panic(err)
	}
	//---------

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("Started...")
	<-done
	fmt.Println("Stopped...")
}
