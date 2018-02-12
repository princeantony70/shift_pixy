

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

//database instance
var appdatabase *sql.DB
var err error


// user struct
type Userstruct struct {
       Username string
       Age string
       Salary string
}


func insertInDatabase(data Userstruct) error {
       //convert age to int
       age,_ := strconv.Atoi(data.Age)

       //convert salary to int
       salary,_ := strconv.Atoi(data.Salary)

      //execute statement
       _, err := appdatabase.Exec("INSERT INTO new_table(name, age, salary)  VALUES(?, ?, ?)", data.Username , age, salary)
       return err

}

func getFromdatabase(uname string, w http.ResponseWriter) error{

       out := Userstruct{}

       err := appdatabase.QueryRow("SELECT * FROM new_table WHERE name=?", uname).
                            Scan(&out.Username, &out.Age, &out.Salary)

       if err != nil {
              return err
       }

       //create json encoder and assign http response ,
      //which implements the IO interface
       enc := json.NewEncoder(w)

       err = enc.Encode(&out)

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

       var k Userstruct

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

       w.Write([]byte(`{"status":"true","message":"success","details":err}`))

}

func userGetHandler(w http.ResponseWriter, r *http.Request) {

       type Userstructlocal struct {
              Username string
       }

       //make byte array
       out := make([]byte,1024)

       //
       bodyLen, err := r.Body.Read(out)

       if err != io.EOF {
              fmt.Println(err.Error())
              w.Write([]byte(`{"error":"bodyRead"}`))
              return
       }

       var k Userstructlocal

       err = json.Unmarshal(out[:bodyLen],&k)


       if err != nil {
              w.Write([]byte(err.Error()))
              return
       }

       err = getFromdatabase(k.Username, w)

       if err != nil {
              w.Write([]byte("{error:" + err.Error() + "}"))
              return
       }


}


func init(){
appdatabase, err = sql.Open("mysql", "root:nfn@/employee")
if err != nil{
  fmt.Println("db error ")
}
err = appdatabase.Ping()
if err !=nil{
  fmt.Println("ping error")
}
}


func main() {

       http.HandleFunc("/adduser", userAddHandler)
       http.HandleFunc("/getuser", userGetHandler)
       http.ListenAndServe(":8084", nil)
}
