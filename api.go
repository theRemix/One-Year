package main

import (
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

func hostCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-protobuf")

	var body []byte
	var bytes []byte
	var err error

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	join := &pb.Join{}
	if err = proto.Unmarshal(body, join); err != nil {
		fmt.Println("Error decoding pb.Join: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// @TODO check if join code is available

	status := &pb.Status{
		Code: pb.Status_Success,
	}
	bytes, err = proto.Marshal(status)
	if err != nil {
		fmt.Println("Error encoding pb.Status: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(bytes)
}

func hostPrompt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-protobuf")

	var body []byte
	var bytes []byte
	var err error

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	prompt := &pb.Prompt{}
	if err = proto.Unmarshal(body, prompt); err != nil {
		fmt.Println("Error decoding pb.Prompt: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// @TODO validate key

	status := &pb.Status{
		Code: pb.Status_Success,
	}
	bytes, err = proto.Marshal(status)
	if err != nil {
		fmt.Println("Error encoding pb.Status: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pe := &pb.PlayerEvent{
		Op:     pb.PlayerEvent_Prompt,
		Prompt: prompt,
	}

	mo := protojson.MarshalOptions{
		UseEnumNumbers: true,
	}
	peJson, err := mo.Marshal(pe)
	if err != nil {
		fmt.Println("Error encoding pb.PlayerEvent: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	events.SendMessage(playerChannel+prompt.Code, sse.SimpleMessage(string(peJson)))

	w.Write(bytes)
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
		Op:         pb.PlayerEvent_Ready,
		PlayerName: join.PlayerName,
	}

	mo := protojson.MarshalOptions{
		UseEnumNumbers: true,
	}
	peJson, err := mo.Marshal(pe)

	events.SendMessage(playerChannel+join.Code, sse.SimpleMessage(string(peJson)))
	events.SendMessage(hostChannel+join.Code, sse.SimpleMessage(string(peJson)))
}

func playerAnswer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-protobuf")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	pa := &pb.PromptAnswer{}
	if err := proto.Unmarshal(body, pa); err != nil {
		fmt.Println("Error decoding pb.PromptAnswer: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("[POST] /api/player/answer %+v", pa)

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

	pe := &pb.PlayerEvent{
		Op:         pb.PlayerEvent_Answered,
		PlayerName: pa.PlayerName,
	}

	mo := protojson.MarshalOptions{
		UseEnumNumbers: true,
	}
	peJson, err := mo.Marshal(pe)

	events.SendMessage(hostChannel+pa.Code, sse.SimpleMessage(string(peJson)))
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
	r.Post("/player/answer", playerAnswer)
	r.Post("/host/create", hostCreate)
	r.Post("/host/prompt", hostPrompt)
	return r
}
