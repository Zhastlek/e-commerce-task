package repositories

import (
	"errors"
	"log"

	"github.com/boltdb/bolt"

	"e-comerce/internal/models"
)

type editStore struct {
	db *bolt.DB
}

type StorageEditor interface {
	CreateOne(p *models.Product) (int, error)
	UpdateOne(p *models.Product) error
	DeleteOne(id int) error
}

func NewStorageEditor(dbEditor *bolt.DB) StorageEditor {
	return &editStore{
		db: dbEditor,
	}
}

func (store *editStore) CreateOne(p *models.Product) (int, error) {
	status, err := store.checkForCopy(p.Name)
	if err != nil {
		return 0, err
	}
	if !status {
		log.Println("error you can not create a copy of the product")
		return 0, errors.New("error you can not create a copy of the product")
	}
	err = store.db.Update(func(tx *bolt.Tx) error {
		// Retrieve the users bucket.
		// This should be created when the DB is first opened.
		b := tx.Bucket([]byte("bucket-by-id"))

		// Generate ID for the user.
		// This returns an error only if the Tx is closed or not writeable.
		// That can't happen in an Update() call so I ignore the error check.
		id, _ := b.NextSequence()
		p.Id = int(id)

		buff := EncodeValue(p)

		// Persist bytes to users bucket.
		return b.Put(EncodeID(p.Id), buff)
	})
	if err != nil {
		log.Fatalf("failure : %s\n", err)
	}
	if err = store.linkNameToId(p); err != nil {
		log.Printf("error unable to link name to id: %v\n", err)
		return 0, err
	}

	return p.Id, nil
}

func (store *editStore) DeleteOne(id int) error {
	if err := store.deleteLink(id); err != nil {
		log.Printf("error can't delete product by id: %v\n", err)
		return err
	}
	err := store.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("bucket-by-id"))
		if err := b.Delete(EncodeID(id)); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (store *editStore) UpdateOne(p *models.Product) error {
	status, err := store.checkForCopy(p.Name)
	if err != nil {
		return err
	}
	if status {
		log.Println("error you can not update copy of the product")
		return errors.New("error you can not update copy of the product")
	}
	err = store.db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte("bucket-by-id"))

		if err := bkt.Put(EncodeID(p.Id), EncodeValue(p)); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	if err = store.UpdateLink(p); err != nil {
		log.Printf("error unable to link name to id: %v\n", err)
		return err
	}
	return nil
}

func (store *editStore) checkForCopy(name string) (bool, error) {
	var status bool
	err := store.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("bucket-by-name"))
		log.Println("b bucket----", b)
		result := b.Get([]byte(name))
		log.Println("result---->>>>>>>>>>>>>>", result, len(result), DecodeID(result))
		if result == nil || len(result) == 0 {
			status = true
		}
		return nil
	})
	if err != nil {
		log.Printf("Error in get by name VIEW method: %v\n", err)
		return status, err
	}
	return status, nil
}

func (store *editStore) linkNameToId(p *models.Product) error {
	err := store.db.Update(func(tx *bolt.Tx) error {
		// Retrieve the users bucket.
		// This should be created when the DB is first opened.
		b := tx.Bucket([]byte("bucket-by-name"))

		// Persist bytes to users bucket.
		return b.Put([]byte(p.Name), EncodeID(p.Id))
	})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (store *editStore) deleteLink(id int) error {
	err := store.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("bucket-by-id"))
		name := b.Get(EncodeID(id))

		bucketByName := tx.Bucket([]byte("bucket-by-name"))
		err := bucketByName.Delete(name)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Printf("error in delete get by name method: %v\n", err)
		return err
	}
	return nil
}

func (store *editStore) UpdateLink(p *models.Product) error {
	err := store.db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte("bucket-by-name"))

		if err := bkt.Put([]byte(p.Name), EncodeID(p.Id)); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
