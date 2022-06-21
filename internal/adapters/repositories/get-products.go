package repositories

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"

	"e-comerce/internal/models"
)

type store struct {
	db *bolt.DB
}

type StorageGetter interface {
	GetOneByName(productName string) (*models.Product, error)
	GetAll() ([]*models.Product, error)
}

func NewStorageGetter(dbGetter *bolt.DB) StorageGetter {
	return &store{
		db: dbGetter,
	}
}

func (store *store) GetOneByName(productName string) (*models.Product, error) {
	var product *models.Product
	err := store.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("bucket-by-name"))
		id := b.Get([]byte(productName))
		if err := DecodeID(id); err != nil {
			log.Printf("error can't find id on name: %v\n", err)
			return err
		}
		bucketById := tx.Bucket([]byte("bucket-by-id"))
		result := bucketById.Get(id)
		product = DecodeValue(result)
		return nil
	})
	if err != nil {
		log.Printf("Error in get by name VIEW method: %v\n", err)
		return nil, err
	}
	return product, nil
}

func (store *store) GetAll() ([]*models.Product, error) {
	products := []*models.Product{}
	err := store.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("bucket-by-id"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=>[%s], value=[%s]\n", k, v)
			p := DecodeValue(v)
			log.Println(p, p.Name, p.Id, p.Price)
			products = append(products, p)
		}
		return nil
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return products, nil
}
