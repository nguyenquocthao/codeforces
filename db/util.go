package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"runtime"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db           *gorm.DB
	PASS         = url.QueryEscape(os.Getenv("PASS"))
	DATABASE_URL = fmt.Sprintf(os.Getenv("DATABASE_URL"), PASS)
)

// const (
// 	DATABASE_URL = "postgres://skills:249zjfvSXscRznZnnYK72tvOH5RWyh8X@localhost:5435/skills"
// )

const (

// DATABASE_URL = "postgres://skills:w2mhUEpjIuFW2ukSKw3K9lTNt1eabq5k@localhost:5435/skills"
// DATABASE_URL = "postgres://skills:w2mhUEpjIuFW2ukSKw3K9lTNt1eabq5k@localhost:5435/skills"

)

func GetDb() *gorm.DB {
	if db != nil {
		return db
	}

	db, _ = gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		QueryFields:            true,
		// Logger:                 gormLogger,
	})
	sqlDB, err := db.DB()

	if err != nil {
		// log.Print("check env DB_WRITE", util.DB_WRITE, err)
		panic("failed to connect database knowledge base: " + err.Error())
	} else {
		fmt.Println("Successfully connected to Database knowledge base!")
	}
	sqlDB.SetConnMaxIdleTime(5 * time.Second)
	// sqlDB.SetMaxIdleConns(50)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	// sqlDB.SetMaxOpenConns(450)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	return db

}

func PrintFileAndLineInfo(args ...any) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("Could not get caller information")
		return
	}
	args = append([]any{file, line}, args...)
	fmt.Println(args...)
}

func Convert[T any](v interface{}) T {
	var res T
	v2, err := json.Marshal(v)
	if err != nil {
		return res
	}
	json.Unmarshal(v2, &res)
	return res
}

// func NewJsonPsql(src interface{}) *string {
// 	if src == nil {
// 		return nil
// 	}
// 	res, err := json.Marshal(src)
// 	if err != nil {
// 		panic(fmt.Sprintf("util.NewJsonPsql marshal error: %v", err))
// 	}
// 	x := string(res)
// 	return &x
// }

type Intents struct {
	Skillid     string
	Intentid    string
	Timestamp   *time.Time
	Version     int
	Name        string
	Context     string
	Language    string
	Public      bool
	InputsCount int
	Data        *string `json:"data"`
}

type Entity struct {
	Entity    string   `json:"entity"`
	Name      string   `json:"name"`
	Questions []string `json:"questions"`
}

// type Text struct {
// 	Text  string   `json:"text"`
// 	Tones []string `json:"tones"`
// }

type Action struct {
	Language string `json:"language"`
	Texts    []any  `json:"texts"`
	Type     string `json:"type"`
}

type Data struct {
	Actions  []Action `json:"actions"`
	Inputs   []string `json:"inputs"`
	Entities []Entity `json:"entities"`
}
