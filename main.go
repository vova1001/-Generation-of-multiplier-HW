package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net/http"
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
	multiplier := sampleMultiplier(rtp)
	resp := Response{Result: multiplier}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

const (
	MinM = 1.0
	MaxM = 10000.0
)

// sampleMultiplier возвращает множитель [1,10000]
// с плавным переходом между зонами
func sampleMultiplier(rtp float64) float64 {
	if rtp <= 0 {
		return MinM
	}
	if rtp > 1 {
		rtp = 1
	}

	u := rand.Float64()

	// "сырые" веса для зон
	w1 := 1.0 - rtp  // для Zone1
	w2 := rtp        // для Zone2
	w3 := rtp / MaxM // для Zone3 (очень маленький)

	// нормализуем веса, чтобы сумма = 1
	sum := w1 + w2 + w3
	massZone1 := w1 / sum
	massZone2 := w2 / sum
	// massZone3 = w3 / sum (оставшаяся вероятность)

	// --- Zone1 ---
	if u < massZone1 {
		return MinM
	}
	u -= massZone1

	// --- Zone2 ---
	if u < massZone2 {
		u2 := u / massZone2
		den := 1.0 - u2*(1.0-1.0/MaxM)
		m := 1.0 / den
		if m < MinM {
			m = MinM
		}
		if m > MaxM {
			m = MaxM
		}
		return m
	}

	// --- Zone3 ---
	return MaxM
}
