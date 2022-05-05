package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/GomaGames/OneYear/pb"
	"github.com/alexandrevicenzi/go-sse"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var events *sse.Server

const (
	playerChannel = "/events/player/"
	hostChannel   = "/events/host/"
	stageChannel  = "/events/stage/"
)

func host(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-protobuf")

	switch action := chi.URLParam(r, "action"); action {
	case "":
	}

	// w.Write(bytes)
}

func join(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-protobuf")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	join := &pb.Join{}
	if err := proto.Unmarshal(body, join); err != nil {
		fmt.Println("Error decoding pb.Join: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("[POST] /api/join %+v", join)

	status := &pb.Status{
		Code: pb.Status_Success,
	}
	bytes, err := proto.Marshal(status)
	if err != nil {
		fmt.Println("Error encoding pb.Status: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(bytes)
}

func playerReady(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-protobuf")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	join := &pb.Join{}
	if err := proto.Unmarshal(body, join); err != nil {
		fmt.Println("Error decoding pb.Join: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("[POST] /api/player/ready %+v", join)

	status := &pb.Status{
		Code: pb.Status_Success,
	}
	statusBytes, err := proto.Marshal(status)
	if err != nil {
		fmt.Println("Error encoding pb.Status: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(statusBytes)

	time.Sleep(1 * time.Second)

	pe := &pb.PlayerEvent{
		Op:   pb.PlayerEvent_Ready,
		Name: join.Name,
	}

	mo := protojson.MarshalOptions{
		UseEnumNumbers: true,
	}
	peJson, err := mo.Marshal(pe)

	events.SendMessage(playerChannel+join.Code, sse.SimpleMessage(string(peJson)))
}

func create(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal([]string{})
	if err != nil {
		fmt.Println("Error json stringifying stageRes : ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(bytes)
}

func apiRouter() http.Handler {

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))

	r.Post("/join", join)
	r.Post("/player/ready", playerReady)
	r.Post("/create", create)
	r.Post("/host/{action}", host)
	return r
}
