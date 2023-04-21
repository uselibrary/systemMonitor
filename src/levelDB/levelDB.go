package levelDB

import (
	"log"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
)

// function to check if a levelDB database exists
// variable: path to the levelDB database
// return: True if the levelDB database exists, False if not
func CheckLevelDB(path string) bool {
	// check if the levelDB database exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// the levelDB database does not exist
		return false
	} else {
		// the levelDB database exists
		return true
	}
}

// function to create a levelDB database
// variable: path to the levelDB database
// return: none
func CreateLevelDB(path string) {
	// create the levelDB database
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

// function to get the value of a key in a levelDB database
// variable: path to the levelDB database, key
// return: value of the key
func GetLevelDBValue(path string, key string) string {
	// open the levelDB database
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// get the value of the key
	value, err := db.Get([]byte(key), nil)
	if err != nil {
		log.Fatal(err)
	}

	// return the value
	return string(value)
}

// function to update the value of a key in a levelDB database
// variable: path to the levelDB database, key, value
// return: none
func UpdateLevelDBValue(path string, key string, value string) {
	// open the levelDB database
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// update the value of the key
	err = db.Put([]byte(key), []byte(value), nil)
	if err != nil {
		log.Fatal(err)
	}
}

// function to write a key-value pair to a levelDB database
// variable: path to the levelDB database, key, value
// key value format: CPU: True, Memory: True, Disk: True, Network: True
// return: none
func WriteLevelDB(path string, key string, value string) {
	// open the levelDB database
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// write a key-value pair to the levelDB database
	err = db.Put([]byte(key), []byte(value), nil)
	if err != nil {
		log.Fatal(err)
	}
}
