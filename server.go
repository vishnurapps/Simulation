package main

import (
    "encoding/json"
    "log"
    "net/http"
    "fmt"
    "io/ioutil"
    "github.com/gorilla/mux"
    //"os"
)

type Test struct {
    Filename  string `json:"filename"`
    Patient_info *Patient `json:"patient_info,omitempty"`
    Result string `json:"result,omitempty"`
}

type Patient struct {
    Bdate    int   `json:"bdate,omitempty"`
    Modality string   `json:"modality,omitempty"`
    Pathology  string   `json:"pathology,omitempty"`
    Pid  string `json:"pid,omitempty"`
    Pname string `json:"pname,omitempty"`
    Psex string `json:"psex,omitempty"`
    Sdate string `json:"sdate,omitempty"`
}

var test []Test


func GetPatientsResult(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(test)
    file, _ := json.MarshalIndent(test, "", " ")
    _ = ioutil.WriteFile("patient.json", file, 0644)
}

func main() {
    fmt.Println("Sample")
    router := mux.NewRouter()
    test = append(test, Test{Filename: "CHEST_PA_10359.png", Patient_info: &Patient{Bdate: 19500101, Modality: "CR", Pathology: "Healthy", Pid: "1-1559-30 NOV 19-9", Pname: "NYANA SELVAN" , Psex: "F ", Sdate: "20191130" }, Result: "SUCCESS"})
    test = append(test, Test{Filename: "3test.png", Patient_info: &Patient{Bdate: 19790225, Modality: "CR", Pathology: "Healthy", Pid: "553882", Pname: "NAME NONE", Psex: "F ", Sdate: "20181005" }, Result: "SUCCESS"})
    router.HandleFunc("/test", GetPatientsResult).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}
