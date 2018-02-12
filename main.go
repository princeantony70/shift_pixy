package main

import (

          "net/http"
          "database/sql"
          "log"
         _ "github.com/go-sql-driver/mysql"

          )


var db *sql.DB
var err error

type questions struct {
  ID int
  Name string
  Section string
  Position int
  Title string
  Submitted_Value string
  Des string
  Ans string
  View_type int
  Parent_id int
  IsRequired int
  Is_submit_field int
  Is_active int
}


func checkErr(err error) {
if err != nil {
log.Fatalln(err*)
}
}


func init(){
db, err = sql.Open("mysql", "root:nfn@/pixy")
checkErr(err)
err = db.Ping()
checkErr(err)
}



func main(){
http.HandleFunc("/",index)
log.Println("Server is up on  port")
log.Fatalln(http.ListenAndServe(":8080", nil))
}


func index(w http.ResponseWriter, req *http.Request){
if req.Method == http.MethodPost{
qus := questions{}
    qus.Name = req.FormValue("name")
    qus.Section  = req.FormValue("section")
    qus.Position  = req.FormValue("position")
    qus.Title = req.FormValue("title")
    qus.Submitted_Value  = req.FormValue("submitted_value")
    qus.Des  = req.FormValue("des")
    qus.Ans  = req.FormValue("ans")
    qus.View_type  = req.FormValue("view_type")
    qus.Parent_id  = req.FormValue("parent_id")
    qus.IsRequired  = req.FormValue("isRequired")
    qus.Is_submit_field  = req.FormValue("is_submit_field")
    qus.Is_active  = req.FormValue("is_active")
    checkErr(err)
    _,err = db.Exec(
      "INSERT INTO  questions (name,section,position,title,submitted_value,des,ans,view_type,parent_id,isRequired,is_submit_field,is_active) VALUES (?, ?, ?, ?,?,?,?,?,?,?,?,?)",
      qus.Name,
      qus.Section,
      qus.Position,
      qus.Title,
      qus.Submitted_Value,
      qus.Des,
      qus.Ans,
      qus.View_type,
      qus.Parent_id,
      qus.IsRequired,
      qus.Is_submit_field,
      qus.Is_active,
    )
    checkErr(err)
    return
    }
    http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
    }
