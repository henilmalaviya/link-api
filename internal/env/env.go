package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	MongoDBURI  string
	MongoDBName string
	Port        int
}

var env *Env

func Get() *Env {
	return env
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	if port == 0 {
		port = 8080
	}

	env = &Env{
		MongoDBURI:  os.Getenv("MONGODB_URI"),
		MongoDBName: os.Getenv("MONGODB_NAME"),
		Port:        port,
	}

	if env.MongoDBURI == "" || env.MongoDBName == "" {
		log.Fatal("Missing required environment variables")
	}
}
