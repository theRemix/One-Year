package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
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

var hostKeys = make(map[string]string)

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

	if _, ok := hostKeys[join.Code]; ok {
		fmt.Println("Code is in use '%s'", join.Code)
		w.WriteHeader(http.StatusBadRequest) // or 302
		return
	}

	var key string
	keyData := make([]byte, 10)
	if _, err := rand.Read(keyData); err == nil {
		keyBytes := sha256.Sum256(keyData)
		key = hex.EncodeToString(keyBytes[:])
	}

	hostCreateStatus := &pb.HostCreateStatus{
		Code: pb.HostCreateStatus_Success,
		Key:  key,
	}
	bytes, err = proto.Marshal(hostCreateStatus)
	if err != nil {
		fmt.Println("Error encoding pb.HostCreateStatus: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	hostKeys[join.Code] = key

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

	if key, ok := hostKeys[prompt.Code]; !ok {
		fmt.Println("Code not found '%s'", prompt.Code)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if key != prompt.Key {
		fmt.Println("Host Key Mismatch for code '%s': incorrect key '%s'", prompt.Code, prompt.Key)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

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
}

func hostResolve(w http.ResponseWriter, r *http.Request) {
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

	resolve := &pb.Resolve{}
	if err = proto.Unmarshal(body, resolve); err != nil {
		fmt.Println("Error decoding pb.Resolve: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if key, ok := hostKeys[resolve.Code]; !ok {
		fmt.Println("Code not found '%s'", resolve.Code)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if key != resolve.Key {
		fmt.Println("Host Key Mismatch for code '%s': incorrect key '%s'", resolve.Code, resolve.Key)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

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

	// broadcast new scores
	pe := &pb.PlayerEvent{
		Op:           pb.PlayerEvent_Resolved,
		PlayerScores: resolve.PlayerScores,
	}
	mo := protojson.MarshalOptions{
		UseEnumNumbers: true,
	}
	peJson, err := mo.Marshal(pe)
	if err != nil {
		fmt.Println("Error encoding pb.peJson: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	events.SendMessage(playerChannel+resolve.Code, sse.SimpleMessage(string(peJson)))
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

	if _, ok := hostKeys[join.Code]; !ok {
		fmt.Println("Code not found '%s'", join.Code)
		w.WriteHeader(http.StatusNotFound)
		return
	}

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
		Op:     pb.PlayerEvent_Answered,
		Answer: pa,
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
	r.Post("/host/resolve", hostResolve)
	return r
}
