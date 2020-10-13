package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	//"os"
)

type Test struct {
	Filename     string   `json:"filename"`
	Patient_info *Patient `json:"patient_info,omitempty"`
	Result       string   `json:"result,omitempty"`
}

type Patient struct {
	Bdate     int    `json:"bdate,omitempty"`
	Modality  string `json:"modality,omitempty"`
	Pathology string `json:"pathology,omitempty"`
	Pid       string `json:"pid,omitempty"`
	Pname     string `json:"pname,omitempty"`
	Psex      string `json:"psex,omitempty"`
	Sdate     string `json:"sdate,omitempty"`
}

type PostInfo struct {
	Pid   string `json:"pid,omitempty"`
	Pname string `json:"pname,omitempty"`
}

func PostPatient(w http.ResponseWriter, r *http.Request) {
	test := Test{Filename: "CHEST_PA_10359.png",
		Patient_info: &Patient{Bdate: 19500101,
			Modality:  "CR",
			Pathology: "Healthy",
			Pid:       "1-1559-30 NOV 19-9",
			Pname:     "NYANA SELVAN",
			Psex:      "F ",
			Sdate:     "20191130"},
		Result: "SUCCESS"}
	var newPost PostInfo
	fmt.Println("Posting patient details")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newPost)
	if err != nil {
		fmt.Println("Inside error")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	test.Patient_info.Pname = newPost.Pname
	test.Patient_info.Pid = newPost.Pid
	file, _ := json.MarshalIndent(test, "", " ")
	filename := "/tmp/" + newPost.Pid + "_" + newPost.Pname + "_report.json"
	_ = ioutil.WriteFile(filename, file, 0644)
}

func main() {
	fmt.Println("Sample")
	router := mux.NewRouter()
	router.HandleFunc("/add", PostPatient).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
