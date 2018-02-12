package main

import (
       "encoding/json"
       "net/http"
       "fmt"
       "io"
       _ "github.com/go-sql-driver/mysql"
       "database/sql"
       "strconv"
)


// _ in package means this package will not cause error if unused
//dtabase instance

var appdatabase *sql.DB
var err error

type Questions struct {
  ID int  `json: id`
  Name string  `json: name`
  Section string  `json: section`
  Position string `json: position`
  Title string  `json: title`
  Submitted_Value string `json: submitted_value`
  Des string `json: des`
  Ans string  `json: ans`
  View_type string `json: view_type`
  Parent_id string  `json: parent_id`
  Is_required string `json: is_required`
  Is_submit_field string `json: is_submit_field`
  Is_active string  `json: is_active`
}

func insertInDatabase(data Questions) error {
       position,_ := strconv.Atoi(data.Position)
       view_type,_ := strconv.Atoi(data.View_type)
       parent_id,_ := strconv.Atoi(data.Parent_id)
       is_required,_ := strconv.Atoi(data.Is_required)
       is_submit_field,_ := strconv.Atoi(data.Is_submit_field)
       is_active,_ := strconv.Atoi(data.Is_active)

       _, err := appdatabase.Exec("INSERT INTO questions(name, section, position,title,submitted_value,des,ans,view_type,parent_id,is_required,is_submit_field,is_active) VALUES(?, ?, ?,?,?,?,?,?,?,?,?,?)", data.Name ,data.Section , position,data.Title,data.Submitted_Value,data.Des,data.Ans,view_type,parent_id,is_required,is_submit_field,is_active)
       return err

}

func userAddHandler(w http.ResponseWriter, r *http.Request) {


       //make byte array
       out := make([]byte,1024)

       //
       bodyLen, err := r.Body.Read(out)

       if err != io.EOF {
              fmt.Println(err.Error())
              w.Write([]byte("{error:" + err.Error() + "}"))
              return
       }

       var k Questions

       err = json.Unmarshal(out[:bodyLen],&k)


       if err != nil {
              w.Write([]byte("{error:" + err.Error() + "}"))
              return
       }

       err = insertInDatabase(k)

       if err != nil {
              w.Write([]byte("{error:" + err.Error() + "}"))
              return
       }


       w.Write([]byte(`{"mesage ":"success"}`))

      // w.Write([]byte(err))

}




func init(){
appdatabase, err = sql.Open("mysql", "root:nfn@/flash")
if err != nil{
  fmt.Println("db error ")
}
err = appdatabase.Ping()
if err !=nil{
  fmt.Println("ping error")
}
}



func main() {

       http.HandleFunc("/adduserflash", userAddHandler)
       http.ListenAndServe(":5096", nil)
}
