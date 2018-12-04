package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/chilakantip/avitar/log"
	"github.com/chilakantip/avitar/pidfile"
	"github.com/chilakantip/my_retail/env"
	"github.com/chilakantip/my_retail/mg_persist"
	"github.com/chilakantip/my_retail/pg_persist"
)

func doCommonSetUp() {
	initLogging()

	if err := pidfile.Dump(); err != nil {
		log.Error("Unable to create pid file, monitoring may be affected...", err)
	}
	cfg := pg_persist.Config{
		Host:     env.DBHost,
		Port:     env.DBPort,
		User:     env.DBUser,
		Password: env.DBPassword,
		Database: env.DBDatabase,
	}
	if err := pg_persist.ConnectToPGDB(cfg); err != nil {
		abort(err)
	}
	fmt.Println("Connected to postgres DB...", cfg.Database)

	//connect to Nosql Mongo DB
	cfgmg := mg_persist.Config{
		Host:       env.DBmgHost,
		Port:       env.DBmgPort,
		Database:   env.DBmgDatabase,
		Collection: env.DBmgCollection,
		User:       env.DBmgUser,
		Password:   env.DBmgPassword,
	}
	if err := mg_persist.ConnectToMongoDB(cfgmg); err != nil {
		abort(err)
	}
	fmt.Println(fmt.Sprintf("Connected to Mongo DB: %s, Collection: %s", cfgmg.Database, cfgmg.Collection))

}

const (
	Success = iota
	SetupFailed
)

func initLogging() {
	logFile := filepath.Join("log", env.AppName+"_"+env.Varsion+"_"+env.AppEnv+".log")
	logCfg := log.Config{
		LogPrefix: env.AppName,
		LogName:   logFile,
		Debug:     false,
		AppName:   env.AppName,
		AppEnv:    env.AppEnv,
	}
	if err := log.Setup(logCfg); err != nil {
		abort(err)
	}
}

func abort(msg error) {
	log.Error(msg)
	fmt.Println(msg)
	pidfile.Drop()
	os.Exit(SetupFailed)
}
