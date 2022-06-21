package repositories

import (
	"bytes"
	"e-comerce/internal/models"
	"encoding/gob"
	"log"
	"strconv"
)

func DecodeValue(info []byte) *models.Product {
	buf := bytes.NewBuffer(info)
	var product *models.Product
	dec := gob.NewDecoder(buf)
	dec.Decode(&product)
	return product
}

func EncodeValue(product *models.Product) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(product)
	return buf.Bytes()
}

func EncodeID(id int) []byte {
	iD := strconv.Itoa(id)
	return []byte(iD)
}

func DecodeID(id []byte) error {
	log.Println(id)
	_, err := strconv.Atoi(string(id))
	if err != nil {
		log.Printf("error can't convert id to integer: %v\n", err)
		return err
	}
	return nil
}
