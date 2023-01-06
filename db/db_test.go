package db

import (
	"log"
	"os"
	"testing"

	"github.com/KyleSmith19091/SpringSearch/query"
)

const (
    WAL_PATH = "./test_log"
    DATA_PATH = "./test_data"
)

var (
    db *DB
)

func cleanupTestLogFiles() {
    os.RemoveAll(WAL_PATH)
    os.RemoveAll(DATA_PATH)
}

func createDB() (*DB) {
    if db == nil {
        db, _ = NewDB(WAL_PATH, DATA_PATH)
        db.Insert("SomeTitle", "Kyle Smith is my name and this is some more stuff")
        db.Insert("Complex", "{ \"name\": \"Ben Stone\", \"age\": 19 }")
        db.Insert("Numbers", "190912928923203298902090 19091")
    }

    return db
}

func TestSearchMultipleReturn(t *testing.T) {
    db := createDB()
    
    res := db.Search(&query.TextQuery{Term: "name"})

    if res.NumResults() != 2 {
        log.Fatalf("Number of results incorrect %v", res.NumResults())
        return
    }
    
    data1, err1 := res.Next() 
    
    if err1 != nil {
        log.Fatalf("Error calling cursor Next() %v", err1)
        return
    }
        
    if data1 != "Kyle Smith is my name and this is some more stuff" &&
        data1 != "{ \"name\": \"Ben Stone\", \"age\": 19 }" {
        log.Fatalf("Error returned from cursor Next()")
    }

    data2, err2 := res.Next()


    if err2 != nil {
        log.Fatalf("Error calling cursor Next() %v", err1)
        return
    }
        
    if data2 != "Kyle Smith is my name and this is some more stuff" &&
        data2 != "{ \"name\": \"Ben Stone\", \"age\": 19 }" {
        log.Fatalf("Error returned from cursor Next()")
    }

    cleanupTestLogFiles()
}

func TestNumberSearch(t *testing.T) {
    db := createDB()

    res := db.Search(&query.TextQuery{Term: "19"})

    if res.NumResults() != 1 {
        log.Fatalf("Number of results incorrect %v", res.NumResults())
        return
    }

    data1, err1 := res.Next() 
    
    if err1 != nil {
        log.Fatalf("Error calling cursor Next() %v", err1)
        return
    }

    if data1 != "{ \"name\": \"Ben Stone\", \"age\": 19 }" {
        log.Fatalf("Error returned from cursor Next()")
        return
    }

    cleanupTestLogFiles()
}

func TestNoResult(t *testing.T) {
    cleanupTestLogFiles()
    db := createDB()

    res := db.Search(&query.TextQuery{Term: "NOPE"})

    if res.NumResults() != 0 {
        log.Fatalf("Number of results incorrect %v", res.NumResults())
        return
    }

    cleanupTestLogFiles()
}
