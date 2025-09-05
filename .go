package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Response struct {
	Result float64 `json:"result"`
}

var rtp float64

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Получаем параметр RTP при запуске
	flag.Float64Var(&rtp, "rtp", 0.95, "Target RTP (0.0 < rtp <= 1.0)")
	flag.Parse()

	if rtp <= 0 || rtp > 1.0 {
		log.Fatal("RTP must be in (0.0, 1.0]")
	}

	http.HandleFunc("/get", handleGet)
	log.Println("Server started on :64333")
	log.Fatal(http.ListenAndServe(":64333", nil))
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	multiplier := generateMultiplier(rtp)

	resp := Response{Result: multiplier}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// Генерация multiplier с распределением, зависящим от rtp
func generateMultiplier(rtp float64) float64 {
	p := rand.Float64() // [0,1)

	if p <= rtp {
		// «Большое» значение, почти всегда больше x
		return 1000 + rand.Float64()*(10000-1000)
	} else {
		// «Маленькое» значение, почти всегда меньше x
		return 1 + rand.Float64()*(50-1)
	}
}
