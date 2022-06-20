package util

import (
	"encoding/json"
	"net/http"
	db "visdb/db"
)

type Item struct {
	Id   int
	Ans1 string
	Ans2 string
	Ans3 string
}
type showAllResponce struct {
	Items []Item `json:"items"`
}

func init() {
	db.Db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(Item{})
}

func ShowAll(w http.ResponseWriter, r *http.Request) {
	var items []Item
	db.Db.Find(&items)
	d := showAllResponce{Items: items}
	res, err := json.Marshal(d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		item := Item{
			Ans1: r.FormValue("ans1"),
			Ans2: r.FormValue("ans2"),
			Ans3: r.FormValue("ans3"),
		}
		db.Db.Create(&item)
	}
	w.WriteHeader(200)
}

func Show(w http.ResponseWriter, r *http.Request) {

}
