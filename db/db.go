package db

import (
	"appname/conf"
	"appname/ent"
	"appname/ent/migrate"
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

var db *ent.Client

func GetClient() *ent.Client {
	return db
}

func Setup() error {
	err := StartConnection()
	if err != nil {
		log.Debug().Err(err).Msg("Failed opening connection to DB")
		return err
	}
	ctx := context.Background()

	err = db.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatal().Msgf("failed creating schema resources: %v", err)
	}

	return nil
}

func getConnection() string {
	conf := conf.Get()
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", conf.DB.User, conf.DB.Pass, conf.DB.Host, conf.DB.Port, conf.DB.DB)
	return connString
}

func StartConnection() error {
	connString := getConnection()
	var err error
	db, err = ent.Open("mysql", connString)
	if err != nil {
		return err
	}
	return nil
}
