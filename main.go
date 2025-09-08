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

func getK(rtp float64) float64 {
	// Таблица точных значений (десятые + 0.99)
	points := []struct {
		rtp float64
		k   float64
	}{
		{0.1, 2.1},
		{0.2, 2.79},
		{0.3, 3.57},
		{0.4, 4.56},
		{0.5, 5.859},
		{0.6, 7.65},
		{0.7, 10.9},
		{0.8, 16.9},
		{0.9, 35.9},
		{0.99, 350},
	}

	// Если rtp ≤ минимального значения
	if rtp <= points[0].rtp {
		return points[0].k
	}
	// Если rtp ≥ максимального значения
	if rtp >= points[len(points)-1].rtp {
		return points[len(points)-1].k
	}

	// Находим соседние точки
	var lower, upper struct{ rtp, k float64 }
	for i := 0; i < len(points)-1; i++ {
		if rtp >= points[i].rtp && rtp <= points[i+1].rtp {
			lower = points[i]
			upper = points[i+1]
			break
		}
	}

	// Логарифмическая интерполяция
	L0 := lower.k * math.Log(lower.rtp)
	L1 := upper.k * math.Log(upper.rtp)
	t := (rtp - lower.rtp) / (upper.rtp - lower.rtp)
	L := L0 + t*(L1-L0)
	k := L / math.Log(rtp)
	return k
}

func generateMultiplier(rtp float64) float64 {
	rand.Seed(time.Now().UnixNano())

	// Ограничение RTP
	if rtp < 0.01 {
		rtp = 0.01
	}
	if rtp > 1.0 {
		rtp = 1.0
	}

	// Спец. случай rtp == 1.0
	if rtp == 1.0 {
		rawMult := 1.0 + rand.Float64()*(10000.0-1.0)
		scale := 0.025
		return 1.0 + (rawMult-1.0)*scale
	}

	// Вычисляем k
	k := getK(rtp)

	// Генерация мультипликатора
	rawMult := 1.0 + rand.Float64()*(10000.0-1.0)
	scale := math.Pow(rtp, k)
	if scale < 0.0001 {
		scale = 0.0001
	}
	return 1.0 + (rawMult-1.0)*scale
}
