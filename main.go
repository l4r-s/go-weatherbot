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

 type DataStruct struct {
     ID        string       `json:"id,omitempty"`
     DeviceID  string       `json:"devid,omitempty"`
     Temp      float64      `json:"temp,omitempty"`
     Hum       float64      `json:"hum,omitempty"`
     Timestamp time.Time    `json:"timestamp,omitempty"`
 }


 func DBQuery (query string) (data []DataStruct) {
   var dbid string
   var devid string
   var temp float64
   var hum float64
   var timestamp float64

   db, err := sql.Open("sqlite3", "./foo.db")
   checkErr(err)
   rows, err := db.Query(query)
   checkErr(err)
   for rows.Next() {
     err = rows.Scan(&dbid, &devid, &temp, &hum, &timestamp)
     checkErr(err)
     timestamp_human := time.Unix(int64(timestamp), 0)
     data = append(data, DataStruct{ID: dbid, DeviceID: devid, Temp: temp, Hum: hum, Timestamp: timestamp_human})
   }
     rows.Close()
     db.Close()
     return
 }

 func GetAllDataEndpoint(w http.ResponseWriter, req *http.Request) {
   var data []DataStruct
   data = DBQuery("SELECT * FROM data")
   json.NewEncoder(w).Encode(data)
 }

 func GetDataEndpoint(w http.ResponseWriter, req *http.Request) {
   var data []DataStruct
   params := mux.Vars(req)
   query := fmt.Sprint("SELECT * FROM data WHERE devid == '", params["devid"], "'")
   data = DBQuery(query)
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
 }

 func main() {
     router := mux.NewRouter()
     router.HandleFunc("/data", GetAllDataEndpoint).Methods("GET")
     router.HandleFunc("/data/{devid}", GetDataEndpoint).Methods("GET")
     router.HandleFunc("/data/{devid}/{temp}/{hum}", PutDataEndpoint).Methods("PUT")
     log.Fatal(http.ListenAndServe(":12345", router))

 }

 func checkErr(err error) {
     if err != nil {
         panic(err)
     }
 }
