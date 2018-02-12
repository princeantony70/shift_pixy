package main

import (
       "encoding/json"
       "net/http"
       "fmt"
       "io"
       _ "github.com/go-sql-driver/mysql"
       "database/sql"
       "strconv"
       "log"
)
var appdatabase *sql.DB
var err error


func database(){
appdatabase, err = sql.Open("mysql", "root:nfn@/flash")
if err != nil{
  fmt.Println("db error ")
}
err = appdatabase.Ping()
if err !=nil{
  fmt.Println("ping error")
}
}


type Questions struct {
  ID int
  Name string
  Section string
  Position string
  Title string
  Submitted_value string
  Des string
  Ans string
  View_type string
  Parent_id string
  Is_required string
  Is_submit_field string
  Is_active string

}

func main(){
  defer appdatabase.Close()
  log.Println("Server is up on 8080 port")
  log.Fatalln(http.ListenAndServe(":8080", nil))
}


var qus Questions
qus.Position = strconv.Atoi()
func post(w http.ResponseWriter, r *http.Request){
  position,_ := strconv.Atoi(Questions.Position)
  view_type,_ := strconv.Atoi(QuestionsView_type)
  parent_id,_ := strconv.Atoi(Questions.Parent_id)
  is_required,_ := strconv.Atoi(Questions.Is_required)
  is_submit_field,_ := strconv.Atoi(Questions.Is_submit_field)
  is_active,_ := strconv.Atoi(Questions.Is_active)
  _, err := appdatabase.Exec("INSERT INTO questions(name, section, position,title,submitted_value,des,ans,view_type,parent_id,is_required,is_submit_field,is_active) VALUES(?, ?, ?,?,?,?,?,?,?,?,?,?)", &Questions.name ,&Questions.section , &Questions.position,&Questions.title,&Questions.submitted_value,&Questions.des,&Questions.ans,&Questions.view_type,&Questions.parent_id,&Questions.is_required,&Questions.is_submit_field,&Questions.is_active)
  if err != nil{
  fmt.Println("values are not inserting in db ")

  }
}
