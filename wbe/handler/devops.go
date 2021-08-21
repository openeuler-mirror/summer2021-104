package handler

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

var (
	ConfigContent Config
	DB            *sql.DB
)

type Config struct {
	Database Database `yaml:"database"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func CheckError(err error) {
	if err == sql.ErrNoRows {
		log.Println(err)
	} else if err != nil {
		panic(err)
	}
}

func (c *Config) GetConf() *Config {
	yamlFile, err := ioutil.ReadFile("config/config.yml")
	CheckError(err)
	err = yaml.Unmarshal(yamlFile, c)
	CheckError(err)
	return c
}

func InitConfig() {
	ConfigContent.GetConf()
}

func InitDB(database Database) {
	username := database.Username
	password := database.Password
	host := database.Host
	port := database.Port
	database1 := database.Database
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, host, port, database1)
	fmt.Println(conn)
	DB, _ = sql.Open("mysql", conn)
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	err := DB.Ping()
	if err != nil {
		panic(err)
	}
}
