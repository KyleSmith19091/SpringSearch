package main

import (
    "log"

	"github.com/KyleSmith19091/SpringSearch/db"
    "github.com/KyleSmith19091/SpringSearch/query"
)

func main() {
    db, err := db.NewDB("./mylog", "./data")
    
    if err != nil {
        log.Fatal(err)
    }

    res := db.Search(&query.TextQuery{Term: "world"})
    log.Printf("%v results found", res.NumResults())
    
    for res.HasNext() {
        data, err := res.Next() 
    
        if err != nil {
            break
        }

        log.Println(data)
    }
}
