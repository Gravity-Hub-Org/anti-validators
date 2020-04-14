package main

import (
	"anti-validators/cacher"
	"anti-validators/config"
	"anti-validators/models"
	"anti-validators/server"
	"anti-validators/signer"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	defaultConfigFileName = "config.json"
)

func openDb(dbName string) *gorm.DB {
	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Request{})
	db.AutoMigrate(&models.Signs{})
	return db
}

func main() {

	ctx := context.Background()

	var confFileName string
	flag.StringVar(&confFileName, "config", defaultConfigFileName, "set host")
	flag.Parse()

	cfg, err := config.Load(confFileName)
	if err != nil {
		panic(err)
	}

	cfg.Ethereum.OracleAddress = strings.ToLower(cfg.Ethereum.OracleAddress)

	db := openDb(cfg.Db)
	defer db.Close()
	//---------

	go server.StartServer(cfg.Host, cfg.Waves.OracleAddress, cfg.Ethereum.OracleAddress, db)

	if err := cacher.StartEthCacher(cfg.Ethereum.ContractAddress, cfg.Ethereum.NodeURL, int64(cfg.RqTimeout), int64(cfg.Ethereum.BlockInterval), ctx, db); err != nil {
		panic(err)
	}
	cacher.StartWavesCacher(cfg.Waves.ContractAddress, cfg.Waves.NodeURL, cfg.RqTimeout, cfg.Waves.BlockInterval, ctx, db)

	if err := signer.StartSigner(cfg, cfg.Ethereum.OracleAddress, cfg.Waves.OracleAddress, ctx, db); err != nil {
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
