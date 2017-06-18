package main

import (
     "database/sql"
     "fmt"
     _"github.com/mattn/go-sqlite3"
     "log"
     "time"
     "net/http"
     "github.com/gorilla/mux"
     "encoding/json"
 )


 func GetAllDataEndpoint(w http.ResponseWriter, req *http.Request) {
   var data []DataStruct
   data = DBQuery("SELECT * FROM data")
   json.NewEncoder(w).Encode(data)
 }

 func GetDataByIdEndpoint(w http.ResponseWriter, req *http.Request) {
   var data []DataStruct
   params := mux.Vars(req)
   query := fmt.Sprint("SELECT * FROM data WHERE devid == '", params["devid"], "'")
   data = DBQuery(query)
   json.NewEncoder(w).Encode(data)
 }

 func GetTempDataByIdEndpoint(w http.ResponseWriter, req *http.Request) {
   var data []DataStructTemp
   params := mux.Vars(req)
   query := fmt.Sprint("SELECT timestamp, temp FROM data WHERE devid == '", params["devid"], "'")
   data = GetTempFromDb(query)
   json.NewEncoder(w).Encode(data)
 }

 func GetHumDataByIdEndpoint(w http.ResponseWriter, req *http.Request) {
   var data []DataStructHum
   params := mux.Vars(req)
   query := fmt.Sprint("SELECT timestamp, hum FROM data WHERE devid == '", params["devid"], "'")
   data = GetHumFromDb(query)
   json.NewEncoder(w).Encode(data)
 }

 func PutDataEndpoint(w http.ResponseWriter, req *http.Request) {
   var data []DataStruct
   var timestamp float64
   params := mux.Vars(req)

   db, err := sql.Open("sqlite3", "./foo.db")
   checkErr(err)
   stmt, err := db.Prepare("INSERT INTO data(devid, temp, hum,timestamp) values(?,?,?,?)")
   checkErr(err)
   timestamp = float64(time.Now().Unix())
   res, err := stmt.Exec(params["devid"], params["temp"], params["hum"], timestamp)
   checkErr(err)
   id, err := res.LastInsertId()
   checkErr(err)
   db.Close()

   query := fmt.Sprint("SELECT * FROM data WHERE id == '", id, "'")
   data = DBQuery(query)
   json.NewEncoder(w).Encode(data)
   fmt.Println("inserted data")

   MailAlert (params["temp"], params["hum"])
 }

 func main() {
     router := mux.NewRouter()
     router.HandleFunc("/data", GetAllDataEndpoint).Methods("GET")
     router.HandleFunc("/data/{devid}", GetDataByIdEndpoint).Methods("GET")
     router.HandleFunc("/data/{devid}/temp", GetTempDataByIdEndpoint).Methods("GET")
     router.HandleFunc("/data/{devid}/hum", GetHumDataByIdEndpoint).Methods("GET")
     router.HandleFunc("/data/{devid}/{temp}/{hum}", PutDataEndpoint).Methods("PUT")
     router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
     http.Handle("/", router)
     log.Fatal(http.ListenAndServe(":8080", router))

 }

 func checkErr(err error) {
     if err != nil {
         panic(err)
     }
 }
