package algo_test

import (
	"bytes"
	"fmt"
	"log"
	"testing"

	"github.com/boltdb/bolt"
)

func init() {

}

func TestBolt(t *testing.T) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	tx, err := db.Begin(true)
	if err != nil {
		fmt.Println(err)
	}
	defer tx.Rollback()

	// Use the transaction...
	_, err = tx.CreateBucket([]byte("MyBucket"))
	if err != nil {
		fmt.Println(err)
	}
	// Commit the transaction and check for error.
	if err := tx.Commit(); err != nil {
		fmt.Println(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		err := b.Put([]byte("answer"), []byte("43"))
		return err
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte("answer"))
		fmt.Printf("The answer is: %s\n", v)
		return nil
	})

	// db.Update(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("MyBucket"))
	// 	var err error
	// 	for i := 0; i < 1000000; i++ {
	// 		err = b.Put([]byte("answer"+strconv.Itoa(i)), []byte(strconv.Itoa(i)))
	// 	}
	//
	// 	return err
	// })

	// db.Update(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("MyBucket"))
	// 	var err error
	// 	for i := 10000; i < 1000000; i++ {
	// 		err = b.Delete([]byte("answer" + strconv.Itoa(i)))
	// 	}
	//
	// 	return err
	// })

	// db.View(func(tx *bolt.Tx) error {
	// 	// Assume bucket exists and has keys
	// 	b := tx.Bucket([]byte("MyBucket"))
	//
	// 	c := b.Cursor()
	//
	// 	for k, v := c.First(); k != nil; k, v = c.Next() {
	// 		fmt.Printf("key=%s, value=%s\n", k, v)
	// 	}
	//
	// 	return nil
	// })

	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		c := tx.Bucket([]byte("MyBucket")).Cursor()

		prefix := []byte("answer999")
		for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	})
}
