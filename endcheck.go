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
  TitleSpanish string `json: titleSpanish`
  Submited_Value string `json: submited_value`
  Spanish_submited_value string `json: spanish_submited_value`
  Des string `json: des`
  Ans string  `json: ans`
  View_type string `json: view_type`
  Parent_id string  `json: parent_id`
  IsRequired string `json: isRequired`
  Is_submit_field string `json: is_submit_field`
  Is_active string  `json: is_active`
}




func insertInDatabase(data Questions) error {
       position,_ := strconv.Atoi(data.Position)
       view_type,_ := strconv.Atoi(data.View_type)
       parent_id,_ := strconv.Atoi(data.Parent_id)
       isRequired,_ := strconv.Atoi(data.IsRequired)
       is_submit_field,_ := strconv.Atoi(data.Is_submit_field)
       is_active,_ := strconv.Atoi(data.Is_active)
       _, err := appdatabase.Exec("INSERT INTO profile_questions(name, section, position,title,titleSpanish,submited_value,spanish_submited_value,des,ans,view_type,parent_id,isRequired,is_submit_field,is_active) VALUES(?, ?, ?,?,?,?,?,?,?,?,?,?,?,?)", data.Name ,data.Section , position,data.Title,data.TitleSpanish,data.Submited_Value,data.Spanish_submited_value,data.Des,data.Ans,view_type,parent_id,isRequired,is_submit_field,is_active)
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
appdatabase, err = sql.Open("mysql", "root:nfn@/shift_pixy")
if err != nil{
  fmt.Println("db error ")
}
err = appdatabase.Ping()
if err !=nil{
  fmt.Println("ping error")
}
}



func main() {

       http.HandleFunc("/addusercheck", userAddHandler)
       http.ListenAndServe(":6008", nil)

}
