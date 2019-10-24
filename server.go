package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	host     string = "mongo"
	database string = "netprotugas4"
	client   *mongo.Client
)

const (
	STATIC_DIR = "./public_html"
)

type Karyawan struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Nama   string             `json:"nama,omitempty" bson:"nama,omitempty"`
	Email  string             `json:"email,omitempty" bson:"email,omitempty"`
	Nomor  string             `json:"nomor,omitempty" bson:"nomor,omitempty"`
	Alamat string             `json:"alamat,omitempty" bson:"alamat,omitempty"`
}

type responseMessage struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []Karyawan `json:"data"`
}

func main() {
	fmt.Println("Starting the application...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	r := mux.NewRouter()
	r.HandleFunc("/api/karyawan", addKaryawan).Methods("POST")
	r.HandleFunc("/api/karyawan", addKaryawan).Methods("OPTIONS")
	r.HandleFunc("/api/karyawan", getKaryawan).Methods("GET")
	r.HandleFunc("/api/karyawan", deleteKaryawan).Methods("DELETE")
	r.HandleFunc("/api/karyawan", patchKaryawan).Methods("PATCH")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public_html/")))
	// r.PathPrefix("/").HandlerFunc(rootHandler)

	server := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 3 * time.Second,
		ReadTimeout:  3 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func addKaryawan(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("content-type", "application/json")
	var karyawan Karyawan
	err := json.NewDecoder(r.Body).Decode(&karyawan)
	if err == nil {
		collection := client.Database("netprotugas4").Collection("karyawan")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, err := collection.InsertOne(ctx, karyawan)
		if err == nil {
			json.NewEncoder(w).Encode(responseMessage{Code: http.StatusOK})
		} else {
			json.NewEncoder(w).Encode(responseMessage{Code: http.StatusInternalServerError})
		}
	} else {
		json.NewEncoder(w).Encode(responseMessage{Code: http.StatusBadRequest})
	}
}

func getKaryawan(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("content-type", "application/json")
	var karyawan []Karyawan
	collection := client.Database("netprotugas4").Collection("karyawan")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	defer cancel()
	if err != nil {
		json.NewEncoder(w).Encode(responseMessage{Code: http.StatusInternalServerError})
		return
	}
	for cursor.Next(ctx) {
		var person Karyawan
		cursor.Decode(&person)
		karyawan = append(karyawan, person)
	}
	if err := cursor.Err(); err != nil {
		json.NewEncoder(w).Encode(responseMessage{Code: http.StatusInternalServerError})
		return
	}
	json.NewEncoder(w).Encode(responseMessage{Code: http.StatusOK, Data: karyawan})
}

func deleteKaryawan(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	// params := mux.Vars(r)
	var karyawan Karyawan
	err := json.NewDecoder(r.Body).Decode(&karyawan)
	collection := client.Database("netprotugas4").Collection("karyawan")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	result, err := collection.DeleteOne(ctx, karyawan)
	defer cancel()
	if err != nil {
		json.NewEncoder(w).Encode(responseMessage{Code: http.StatusInternalServerError})
		return
	}
	if result.DeletedCount == 0 {
		json.NewEncoder(w).Encode(responseMessage{Code: http.StatusOK, Message: "No data is deleted"})
	} else {
		json.NewEncoder(w).Encode(responseMessage{Code: http.StatusOK, Message: "Data deletes successfully"})
	}
}

func patchKaryawan(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("content-type", "application/json")
	var karyawan Karyawan
	err := json.NewDecoder(r.Body).Decode(&karyawan)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(responseMessage{Code: http.StatusInternalServerError})
		return
	}
	collection := client.Database("netprotugas4").Collection("karyawan")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	filter := bson.D{primitive.E{Key: "_id", Value: karyawan.ID}}
	dataBaru := bson.D{primitive.E{
		Key: "$set",
		Value: bson.D{primitive.E{Key: "nama", Value: karyawan.Nama},
			primitive.E{Key: "email", Value: karyawan.Email},
			primitive.E{Key: "nomor", Value: karyawan.Nomor},
			primitive.E{Key: "alamat", Value: karyawan.Alamat}},
	}}
	result, err := collection.UpdateOne(ctx, filter, dataBaru)
	defer cancel()
	if result.MatchedCount == 0 {
		json.NewEncoder(w).Encode(responseMessage{Code: http.StatusOK, Message: "No data is matched"})
	} else if result.ModifiedCount == 0 {
		json.NewEncoder(w).Encode(responseMessage{Code: http.StatusOK, Message: "No data is changed"})
	} else {
		json.NewEncoder(w).Encode(responseMessage{Code: http.StatusOK, Message: "Data changes successfully"})
	}
}
