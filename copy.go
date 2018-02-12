package main

import (
       "encoding/json"
       "net/http"
       _ "github.com/go-sql-driver/mysql"
       "database/sql"
)


func init(){
appdatabase, err = sql.Open("mysql", "root:nfn@/pixy")
if err != nil{
  fmt.Println("db error ")
}
err = appdatabase.Ping()
if err !=nil{
  fmt.Println("ping error")
}
}
