

package main

import (

       "encoding/json"
     _ "github.com/go-sql-driver/mysql"
       "database/sql"
       "log"
)



type Tag struct {
	ID   int
	Name string
  Age  int  `json:",string "`


}

func main() {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:nfn@(127.0.0.1:3306)/tuts")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT id, name FROM tags")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tag Tag

    err := json.Unmarshal([]byte(s), &tag)
		// for each row, scan the result into our tag composite object

		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
                // and then print out the tag's Name attribute
		log.Printf(tag.Age)
	}

}




/*
package main

import (
    "encoding/json"
    "fmt"
    "strings"
)

type Product struct {
    Name  string `json:"name"`
    Price float64 `json:"price,string"`
}

func main() {
    s := `{"name":"Galaxy Nexus","price":"3460.00"}`
    var pro Product
    err := json.NewDecoder(strings.NewReader(s)).Decode(&pro)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(pro)
} */
