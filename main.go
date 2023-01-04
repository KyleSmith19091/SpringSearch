package main

import (
    "log"

	"springsearch/db"
)

func main() {
    db, err := db.NewDB("./mylog", "./data")
    
    if err != nil {
        log.Fatal(err)
    }

    res := db.Search("world")
    log.Printf("%v results found", res.NumResults())
    
    for res.HasNext() {
        data, err := res.Next() 
    
        if err != nil {
            break
        }

        log.Println(data)
    }
}
