package repositories

import (
	"e-comerce/pkg"
	"log"

	"github.com/boltdb/bolt"
)

func InitDB(config *pkg.Config) *bolt.DB {
	db, err := bolt.Open("./internal/adapters/repositories/"+config.NameDB, 0600, nil)
	if err != nil {
		log.Printf("error while opening-->%v\n", err)
		log.Fatal(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(config.BucketById))
		if err != nil {
			log.Printf("error while create bucket-by-id-->%v\n", err)
			log.Fatal(err)
		}
		_, err = tx.CreateBucketIfNotExists([]byte(config.BucketByName))
		if err != nil {
			log.Printf("error while create bucket-by-name-->%v\n", err)
			log.Fatal(err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
