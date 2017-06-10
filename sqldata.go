package main

import (
     "database/sql"
     _"github.com/mattn/go-sqlite3"
     "time"
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


type DataStructTemp struct {
    Time             float64      `json:"time,omitempty"`
    Temperature      float64      `json:"temperature,omitempty"`
}

func GetTempFromDb (query string) (data []DataStructTemp) {
  var timestamp float64
  var temp float64

  db, err := sql.Open("sqlite3", "./foo.db")
  checkErr(err)
  rows, err := db.Query(query)
  checkErr(err)
  for rows.Next() {
    err = rows.Scan(&timestamp, &temp)
    checkErr(err)
    data = append(data, DataStructTemp{Time: timestamp, Temperature: temp})
  }
    rows.Close()
    db.Close()
    return
}


type DataStructHum struct {
    Time             float64      `json:"time,omitempty"`
    Humidity         float64      `json:"humidity,omitempty"`
}

func GetHumFromDb (query string) (data []DataStructHum) {
  var timestamp float64
  var hum float64

  db, err := sql.Open("sqlite3", "./foo.db")
  checkErr(err)
  rows, err := db.Query(query)
  checkErr(err)
  for rows.Next() {
    err = rows.Scan(&timestamp, &hum)
    checkErr(err)
    data = append(data, DataStructHum{Time: timestamp, Humidity: hum})
  }
    rows.Close()
    db.Close()
    return
}
