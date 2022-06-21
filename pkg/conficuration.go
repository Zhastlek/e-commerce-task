package pkg

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	NameDB       string
	BucketByName string
	BucketById   string
}

func Configuration() *Config {
	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
		fmt.Printf(".env file load error: %v\n", err)
	}
	return &Config{
		NameDB:       os.Getenv("nameDB"),
		BucketByName: os.Getenv("bucket_by_name"),
		BucketById:   os.Getenv("bucket_by_id"),
	}
}
