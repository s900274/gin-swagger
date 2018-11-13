package main

import (
	"define"
	"flag"
	"errors"
	"log"
	"github.com/BurntSushi/toml"
	logger "github.com/shengkehua/xlog4go"
	"sync"
	"httpservice"
	"os"
	"database/sql"
)

func initConf(flagSet *flag.FlagSet) error {

	flagSet.Parse(os.Args[1:])
	configFile := flagSet.Lookup("config").Value.String()
	if configFile != "" {
		_, err := toml.DecodeFile(configFile, &define.Cfg)
		if err != nil {
			log.Fatalf("ERROR: failed to load config file %s - %s\n", configFile, err.Error())
			return err
		}

	} else {
		log.Fatalln("ERROR: config file is nil")
		err := errors.New("ERROR: config file is nil")
		return err
	}
	return nil
}

func initLogger() error {
	err := logger.SetupLogWithConf(define.Cfg.LogFile)
	return err
}

func initSqlite() error {

	db, err := sql.Open("sqlite3", "./db/building.db")

	if err != nil {
		return err
	}

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS companies (cid INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(64) NULL, floor VARCHAR(64) NULL);")

	if err != nil {
		return err
	}

	_, err = stmt.Exec()

	if err != nil {
		return err
	}

	stmt, err = db.Prepare("INSERT INTO companies(floor, name) values(?,?)")
	if err != nil {
		return err
	}

	_, _ = stmt.Exec("3", "喜登數位")
	if err != nil {
		return err
	}

	return nil
}


func RunHttpServer() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		HServer := httpservice.NewHTTPServer()
		err := HServer.InitHttpServer()
		if nil != err {
			logger.Error("HTTP ServerStart failed, err :%v", err)
			return
		}
	}()
	wg.Wait()
}
