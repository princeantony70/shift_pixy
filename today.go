package main

import (
       "encoding/json"
       "net/http"
       "fmt"
       _ "github.com/go-sql-driver/mysql"
       "database/sql"
       "strconv"
       "io"
)


func db(){
appdatabase, err = sql.Open("mysql", "root:nfn@/shift_pixy")
if err != nil{
  fmt.Println("db error ")
}
err = appdatabase.Ping()
if err !=nil{
  fmt.Println("ping error")
}
}



var appdatabase *sql.DB
var err error


type Questions struct {
  ID int  `json: id`
  Name string  `json: name`
  Section string  `json: section`
  Position string `json: position`
  Title string  `json: title`
  Titlespanish string `json: titlespanish`
  Submited_Value string `json: submitted_value`
  Spanish_submitted_value string `json: spanish_submited_value`
  Des string `json: des`
  Ans string  `json: ans`
  View_type string `json: view_type`
  Parent_id string  `json: parent_id`
  Is_required string `json: isRequired`
  Is_submit_field string `json: is_submit_field`
  Is_active string  `json: is_active`
}

func userrouter(w http.ResponseWriter, r *http.Request){

  out := make([]byte,1024)
  bodyLen, err := r.Body.Read(out)

  if err != io.EOF {
         fmt.Println(err.Error())
         w.Write([]byte("{error:" + err.Error() + "}"))
         return
  }

  var user Questions

  err = json.Unmarshal(out[:bodyLen],&user)

    // Use Atoi to parse string.
    position, _ := strconv.Atoi(user.Position)
    view_type, _ := strconv.Atoi(user.View_type)
    parent_id,_ := strconv.Atoi(user.Parent_id)
    isRequired,_:= strconv.Atoi(user.Is_required)
    is_submit_field,_:=strconv.Atoi(user.Is_submit_field)
    is_active,_:=strconv.Atoi(user.Is_active)

 _, err = appdatabase.Exec("INSERT INTO questions(name, section, position,title,titlespanish,submited_value,spanish_submited_value,des,ans,view_type,parent_id,isRequired,is_submit_field,is_active) VALUES(?, ?, ?,?,?,?,?,?,?,?,?,?,?,?)", user.Name ,user.Section , position,user.Title,user.Titlespanish,user.Submited_Value,user.Spanish_submitted_value,user.Des,user.Ans,view_type,parent_id,isRequired,is_submit_field,is_active)
   if err != nil{
     fmt.Println("not inserted ")
    }


}

  func main() {

         http.HandleFunc("/insert", userrouter)
         fmt.Println("server is on port")
         http.ListenAndServe(":7010", nil)
  }
