package main

import ( "fmt"
          "encoding/json"
          "net/http"
        _ "github.com/go-sql-driver/mysql"
          "database/sql"
          "log"
          "strconv"
         )


var db *sql.DB
var err error


fun dbconnection(){
   db, err = sql.Open("mysql", "root:nfn@/pixy")
   if err != nil{
     fmt.Println("db connection error ")
   }
   err = db.ping()
     if err !=nil{
       fmt.Println("db is not pinged properly")
     }
   }


type Questions struct{
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
  defer db.Close()
  log.Println("Server is up on 8080 port")
  log.Fatalln(http.ListenAndServe(":8080", nil))
}



func post(w http.ResponseWriter, r *http.Request)(data Questions){
  position,_ := strconv.Atoi(data.Position)
  view_type,_ := strconv.Atoi(data.View_type)
  parent_id,_ := strconv.Atoi(data.Parent_id)
  is_required,_ := strconv.Atoi(data.IsRequired)
  is_submit_field,_ := strconv.Atoi(data.Is_submit_field)
  is_active,_ := strconv.Atoi(data.Is_active)
  _, err := db.Exec("INSERT INTO questions(name, section, position,title,submitted_value,des,ans,view_type,parent_id,is_required,is_submit_field,is_active) VALUES(?, ?, ?,?,?,?,?,?,?,?,?,?)", data.Name ,data.Section , position,data.Title,data.Submitted_Value,data.Des,data.Ans,view_type,parent_id,is_required,is_submit_field,is_active)
  if err != nil{
  fmt.Println("values are not inserting in db ")
  }
  rows,err := db.Query(
    `SELECT name,
     section,
      position,
      title,
      submitted_value,
      des,
      ans,
      view_type,
      parent_id,
      is_required,
      is_submit_field,
      is_active
       FROM Questions;
       `)
  if err ! = nil{
    fmt.Println("select query not working")

  }
  data := make([]dat, 0)
  for rows.Next() {
  dt:= dat{}
  rows.Scan(&dt.Id,&dt.Name,&dt.Section,&dt.Position,&dt.Title,&dt.Submitted_Value,&dt.Des,&dt.Ans,&dt.View_type,&dt.Parent_id,&dt.Is_required,&dt.Is_submit_field,&dt.Is_active)
  data = append(data,dt)
  if (&dt.parent_id == NULL){

        _, er := db.Exec(
                  "UPDATE questions SET parent_id = ?  ",
  }

}

func return(w http.ResponseWriter, r *http.Request){

  json.NewEncoder(w).Encode(data)


}
}
