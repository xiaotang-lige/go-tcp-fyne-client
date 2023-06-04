package tool

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"go_gui/file"
)

type boltMy struct {
	table string
}

func (t *boltMy) Table(table string) *boltMy {
	return &boltMy{table: table}
}
func (t *boltMy) Set(id string, v interface{}) error {
	tx, err := file.Db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var c *bolt.Bucket
	c, err = tx.CreateBucket([]byte(t.table))
	if err != nil {
		c = tx.Bucket([]byte(t.table))
	}
	va, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = c.Put([]byte(id), va)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}
func (t *boltMy) Get(id string) ([]byte, error) {
	tx, err := file.Db.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	var c *bolt.Bucket
	c, err = tx.CreateBucket([]byte(t.table))
	if err != nil {
		c = tx.Bucket([]byte(t.table))
	}

	va := c.Get([]byte(id))
	if err != nil {
		return nil, err
	}
	return va, nil
}
func (t *boltMy) Delete(id string) error {
	tx, err := file.Db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var c *bolt.Bucket
	c, err = tx.CreateBucket([]byte(t.table))
	if err != nil {
		c = tx.Bucket([]byte(t.table))
	}

	err = c.Delete([]byte(id))
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}
