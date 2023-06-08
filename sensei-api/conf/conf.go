package conf

import (
	"errors"
	"path/filepath"
	"runtime"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Conf struct {
	IP     string
	Port   string
	DB     DatabaseConf
	Mail   MailConf
	JwtKey string
}

type DatabaseConf struct {
	Host string
	Port string
	User string
	Pass string
	DB   string
}

type MailConf struct {
	User     string
	Pass     string
	SmtpHost string
	SmtpPort string
}

var conf Conf

func Setup() error {
	_, currentPath, _, _ := runtime.Caller(0)

	log.Info().Msg("Initializing conf setup")
	viper.SetConfigName("configuration")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join(filepath.Dir(currentPath), ".."))
	err := readAndUnmarshal()

	if err != nil {
		log.Error().Err(err).Msg("Error reading conf file")
		return errors.New("Error reading conf file")
	}

	log.Info().Msg("Conf file loaded succesfully")
	return nil
}

func readAndUnmarshal() error {
	err := viper.ReadInConfig()
	if err != nil {
		log.Error().Err(err)
		return err
	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Debug().Err(err)
		return err
	}
	return nil
}

func Get() *Conf {
	return &conf
}
