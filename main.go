package main

import (
	"encoding/json"
	"flag"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
)

type Response struct {
	Result float64 `json:"result"`
}

var rtp float64

func main() {
	flag.Float64Var(&rtp, "rtp", -1.0, "Target RTP (0.0 < rtp <= 1.0)")
	flag.Parse()

	if rtp <= 0 || rtp > 1.0 {
		log.Fatal("RTP must be in (0.0, 1.0]")
	}

	http.HandleFunc("/get", handleGet)
	log.Println("Server started on :64333")
	log.Fatal(http.ListenAndServe("0.0.0.0:64333", nil))
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	multiplier := generateMultiplier(rtp)
	resp := Response{Result: multiplier}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func generateMultiplier(rtp float64) float64 {
	rand.Seed(time.Now().UnixNano())

	// Ограничение RTP
	if rtp <= 0 {
		rtp = 0.000001
	}
	if rtp >= 1.0 {
		rtp = 1.0
	}

	// Спец. случай rtp == 1.0
	if rtp == 1.0 {
		rawMult := 1.0 + rand.Float64()*(10000.0-1.0)
		scale := 0.025 // сжатие к 1
		return 1.0 + (rawMult-1.0)*scale
	}

	// Целевая средняя (targetScale) для мультипликатора
	targetScale := rtp

	// Вычисляем k динамически, чтобы scale ~= targetScale
	// scale = rtp^k => k = ln(targetScale)/ln(rtp)
	k := math.Log(targetScale) / math.Log(rtp)

	// Генерация raw multiplier
	rawMult := 1.0 + rand.Float64()*(10000.0-1.0)

	// Масштабирование
	scale := math.Pow(rtp, k)
	mult := 1.0 + (rawMult-1.0)*scale

	return mult
}
