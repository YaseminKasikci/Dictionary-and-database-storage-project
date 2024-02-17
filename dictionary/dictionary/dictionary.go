package dictionary

import (
	"fmt"
	"time"

	"github.com/dgraph-io/badger"
)

// dict db cle / valeur (badger)
type Dictionary struct {
	db *badger.DB
}

// pour stocké les entrées dans le dict
type Entry struct {
	Word string
	Definition string
	CreateAt time.Time
}

//affichage de l'Entry
func (e Entry) String() string {
	created := e.CreateAt.Format(time.Stamp)
	return fmt.Sprintf("%-10v\t%-50v%-6v", e.Word, e.Definition, created)
}

// ouvre la db et stock
func New(dir string) (*Dictionary, error) {
	// logger := badger.DefaultLogger()
	// logger.SetLogLevel(badger.WARNING)

	opts := badger.DefaultOptions("") //or dir ?
	opts.Dir = dir
	opts.ValueDir = dir

 db, err :=	badger.Open(opts)
 if err != nil {
	return nil, err
 }

 dict := &Dictionary{
	db: db, 
 }
 return dict, nil
}

//Ferme la db
func (d *Dictionary) Close() {
	d.db.Close()
}
