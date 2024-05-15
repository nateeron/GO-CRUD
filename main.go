package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// install go get -u github.com/gorilla/mux
// Define a struct to represent an item
type Item struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// Simulated database
var items []Item

// Get all items
func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// Get a single item by ID
func getItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range items {
		if item.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

// Create a new item
func createItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	_ = json.NewDecoder(r.Body).Decode(&newItem)
	items = append(items, newItem)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newItem)
}

// Update an existing item
func updateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range items {
		if item.ID == params["id"] {
			items = append(items[:index], items[index+1:]...)
			var updatedItem Item
			_ = json.NewDecoder(r.Body).Decode(&updatedItem)
			updatedItem.ID = params["id"]
			items = append(items, updatedItem)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}
	http.NotFound(w, r)
}

// Delete an item
func deleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range items {
		if item.ID == params["id"] {
			items = append(items[:index], items[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(items)
			return
		}
	}
	http.NotFound(w, r)
}

func main() {
	// Initialize router
	r := mux.NewRouter()

	// Mock data
	items = append(items, Item{ID: "1", Name: "Item One", Price: 100})
	items = append(items, Item{ID: "2", Name: "Item Two", Price: 200})

	// Route handles & endpoints
	r.HandleFunc("/items", getItems).Methods("GET")
	r.HandleFunc("/items/{id}", getItem).Methods("GET")
	r.HandleFunc("/items", createItem).Methods("POST")
	r.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	// Start server
	log.Println("Starting server on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// RUN
//-----------------------------------------------------
//[ Create mod File ]
// go mod init CRUD
//[ install
// go get -u github.com/gorilla/mux
//-----------------------------------------------------------
// go run main.go
//-------------------------------------------------------
//For TEST
//curl http://localhost:8000/items
// curl http://localhost:8000/items/1

// curl -X POST -H "Content-Type: application/json" -d '{"id":"3","name":"Item Three","price":300}' http://localhost:8000/items
// curl -X DELETE http://localhost:8000/items/1

//*******[Build Deploy IIS]*******************
// go build -o main.exe G:\M_save\GO\CRUD
// use NSSM ต้อง เซ็ต Parth c]h; Run บน Powershell Administrator
// nssm install MyGoApp "G:\M_save\GO\CRUD\main.exe"
// Ountput : Service "MyGoApp" installed successfully!
// nssm start MyGoApp
// Ountput : MyGoApp: START: The operation completed successfully.
// หลังจากนั้นให้ใช้
//**********************************************
//** nssm list --> Check All
//** nssm status MyGoApp
//** nssm stop MyGoApp
//----------------------------------------------------------------------------
// Exeample Powershell
//----------------------------------------------------------------------------
// PS C:\WINDOWS\system32> nssm install MyGoApp "G:\M_save\GO\CRUD\main.exe"
// Service "MyGoApp" installed successfully!
// PS C:\WINDOWS\system32> nssm start MyGoApp
// MyGoApp: START: The operation completed successfully.
// PS C:\WINDOWS\system32>
//----------------------------------------------------------------------------
