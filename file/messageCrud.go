package file

import (
	"encoding/binary"
	"github.com/boltdb/bolt"
	"log"
)

type mesage struct {
}

func (*mesage) Set(id string, mes []byte) {
	t, err := Db.Begin(true)
	defer t.Rollback()
	var b *bolt.Bucket
	if b = t.Bucket([]byte(id)); b == nil {
		b, err = t.CreateBucket([]byte(id))
	}
	i, _ := b.NextSequence()

	err = b.Put(itob(int(i)), mes)
	if err != nil {
		log.Println(err)
		t.Rollback()
		return
	}
	t.Commit()
}
func (*mesage) Get() {

}
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
