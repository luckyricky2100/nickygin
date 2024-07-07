package recitewords

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

type Word struct {
	ID     int
	Word   string
	Status string
}

var db *sql.DB
var redisClient *redis.Client
var memoryCycles = []time.Duration{
	5 * time.Minute,
	30 * time.Minute,
	12 * time.Hour,
	24 * time.Hour,
	2 * 24 * time.Hour,
	4 * 24 * time.Hour,
	7 * 24 * time.Hour,
	15 * 24 * time.Hour,
}

func main() {
	var err error
	db, err = sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	for {
		word, err := getWordFromDB()
		if err != nil {
			panic(err)
		}

		fmt.Println("Do you know this word?", word.Word)
		fmt.Println("1. Don't know")
		fmt.Println("2. Vague")
		fmt.Println("3. Know")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			setNextAppearance(word, 0)
		case 2:
			setNextAppearance(word, 0)
		case 3:
			cycle, _ := redisClient.Get(fmt.Sprintf("word_cycle_%d", word.ID)).Int()
			setNextAppearance(word, cycle+1)
		}
	}
}

func getWordFromDB() (*Word, error) {
	var word Word
	err := db.QueryRow("SELECT id, word, status FROM words ORDER BY RAND() LIMIT 1").Scan(&word.ID, &word.Word, &word.Status)
	if err != nil {
		return nil, err
	}
	return &word, nil
}

func setNextAppearance(word *Word, cycle int) {
	if cycle >= len(memoryCycles) {
		cycle = len(memoryCycles) - 1
	}
	expiration := memoryCycles[cycle]
	redisClient.Set(fmt.Sprintf("word_cycle_%d", word.ID), cycle, expiration)
}
