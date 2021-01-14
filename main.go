package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"time"
	"encoding/json"
)

var session, _ = mgo.Dial("127.0.0.1")
var c = session.DB("TutDb").C("ToDo")

type ToDoItem struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Date time.Time
	Description string
	Done bool
}

func AddToDo(w http.ResponseWriter, r *http.Request) {
	_ = c.Insert(ToDoItem{
			bson.NewObjectId(),
			time.Now(),
			r.FormValue("description"),
			false,
	})
	result := ToDoItem{}
	_ = c.Find(bson.M{"description": r.FormValue("description")}).One(&result)
	json.NewEncoder(w).Encode(result)
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // Set 200 OK
	w.Header().Set("Content-Type", "application/json") 
	io.WriteString(w, `{"alive": true}`) // Send json to the ResponseWriter
}

func GetByID(id string) []ToDoItem {
	var result ToDoItem
	var res []ToDoItem
	_ = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	res = append(res, result)
	return res
}

func GetToDo(w http.ResponseWriter, r *http.Request) {
	var res []ToDoItem

	vars := mux.Vars(r)
	id := vars["id"]
	if id != "" {
			res = GetByID(id)
	} else {
			_ = c.Find(nil).All(&res)
	}

	json.NewEncoder(w).Encode(res)
}

func MarkDone(w http.ResponseWriter, r*http.Request) {
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	err := c.Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"done": true}})
	if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json")
			io.WriteString(w, `{"updated": false, "error": ` + err.Error() + `}` )
	} else {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			io.WriteString(w, `{"updated": true}`)
	}
}

func DelToDo(w http.ResponseWriter, r*http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := c.RemoveId(bson.ObjectIdHex(id))
	if err == mgo.ErrNotFound  {
			json.NewEncoder(w).Encode(err.Error())
	} else {
			io.WriteString(w, "{result: 'OK'}")
	}
}

func main() {
	session.SetMode(mgo.Monotonic, true)
	defer session.Close()

	router := mux.NewRouter()
	router.HandleFunc("/todo", GetToDo).Methods("GET")
	router.HandleFunc("/todo", AddToDo).Methods("POST", "PUT")
	router.HandleFunc("/todo/{id}", GetToDo).Methods("GET") // Define a route with an id Variable
	router.HandleFunc("/todo/{id}", MarkDone).Methods("PATCH")
	router.HandleFunc("/health", Health).Methods("GET")
	router.HandleFunc("/todo/{id}", DelToDo).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":8000", router))
}