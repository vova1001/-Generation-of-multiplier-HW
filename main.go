package main

import (
	"encoding/json"
	"flag"
	"log"
	"math"
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
	multiplier := generateMultiplier(rtp)
	resp := Response{Result: multiplier}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func generateMultiplier(rtp float64) float64 {

	// При значениях близких к 1, коф "k" сильно расходился из-за чего этот метод исключил
	

	// keyRTP := []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 0.99}
	// keyK := []float64{2.1, 2.79, 3.57, 4.56, 5.859, 7.65, 10.699, 10.741, 35.9, 408}

	// var k float64

	// // Интерполяция k
	// if rtp <= keyRTP[0] {
	// 	k = keyK[0]
	// } else if rtp >= keyRTP[len(keyRTP)-1] {
	// 	k = keyK[len(keyK)-1]
	// } else {
	// 	for i := 0; i < len(keyRTP)-1; i++ {
	// 		if rtp >= keyRTP[i] && rtp <= keyRTP[i+1] {
	// 			t := (rtp - keyRTP[i]) / (keyRTP[i+1] - keyRTP[i])
	// 			k = keyK[i] + t*(keyK[i+1]-keyK[i])
	// 			break
	// 		}
	// 	}
	// }

	var k float64
	// Диапазон 0.1-0.2
	if rtp >= 0.10 && rtp < 0.20 {
		for i := 10; i <= 19; i++ {
			checkRTP := float64(i) / 100.0
			if rtp <= checkRTP {
				switch i {
				case 10:
					k = 2.10
				case 11:
					k = 2.15
				case 12:
					k = 2.18
				case 13:
					k = 2.20
				case 14:
					k = 2.23
				case 15:
					k = 2.25
				case 16:
					k = 2.28
				case 17:
					k = 2.30
				case 18:
					k = 2.32
				case 19:
					k = 2.35
				}
				break
			}
		}
	}

	// Диапазон 0.2-0.3
	if rtp >= 0.20 && rtp < 0.30 {
		for i := 20; i <= 29; i++ {
			checkRTP := float64(i) / 100.0
			if rtp <= checkRTP {
				switch i {
				case 20:
					k = 2.79
				case 21:
					k = 2.82
				case 22:
					k = 2.85
				case 23:
					k = 2.88
				case 24:
					k = 2.91
				case 25:
					k = 2.94
				case 26:
					k = 2.97
				case 27:
					k = 3.00
				case 28:
					k = 3.03
				case 29:
					k = 3.05
				}
				break
			}
		}
	}

	// Диапазон 0.3-0.4
	if rtp >= 0.30 && rtp < 0.40 {
		for i := 30; i <= 39; i++ {
			checkRTP := float64(i) / 100.0
			if rtp <= checkRTP {
				switch i {
				case 30:
					k = 3.57
				case 31:
					k = 3.60
				case 32:
					k = 3.63
				case 33:
					k = 3.66
				case 34:
					k = 3.69
				case 35:
					k = 3.72
				case 36:
					k = 3.75
				case 37:
					k = 3.78
				case 38:
					k = 3.81
				case 39:
					k = 3.84
				}
				break
			}
		}
	}

	// Диапазон 0.4-0.5
	if rtp >= 0.40 && rtp < 0.50 {
		for i := 40; i <= 49; i++ {
			checkRTP := float64(i) / 100.0
			if rtp <= checkRTP {
				switch i {
				case 40:
					k = 4.56
				case 41:
					k = 4.60
				case 42:
					k = 4.64
				case 43:
					k = 4.68
				case 44:
					k = 4.72
				case 45:
					k = 4.76
				case 46:
					k = 4.80
				case 47:
					k = 4.84
				case 48:
					k = 4.88
				case 49:
					k = 4.92
				}
				break
			}
		}
	}

	// Диапазон 0.5-0.6
	if rtp >= 0.50 && rtp < 0.60 {
		for i := 50; i <= 59; i++ {
			checkRTP := float64(i) / 100.0
			if rtp <= checkRTP {
				switch i {
				case 50:
					k = 5.859
				case 51:
					k = 5.90
				case 52:
					k = 5.94
				case 53:
					k = 5.98
				case 54:
					k = 6.02
				case 55:
					k = 6.06
				case 56:
					k = 6.10
				case 57:
					k = 6.14
				case 58:
					k = 6.18
				case 59:
					k = 6.22
				}
				break
			}
		}
	}

	// Диапазон 0.6-0.7
	if rtp >= 0.60 && rtp < 0.70 {
		for i := 60; i <= 69; i++ {
			checkRTP := float64(i) / 100.0
			if rtp <= checkRTP {
				switch i {
				case 60:
					k = 7.65
				case 61:
					k = 7.80
				case 62:
					k = 7.95
				case 63:
					k = 8.10
				case 64:
					k = 8.25
				case 65:
					k = 8.40
				case 66:
					k = 8.55
				case 67:
					k = 8.70
				case 68:
					k = 8.85
				case 69:
					k = 9.00
				}
				break
			}
		}
	}

	// Диапазон 0.7-0.8
	if rtp >= 0.70 && rtp < 0.80 {
		for i := 70; i <= 79; i++ {
			checkRTP := float64(i) / 100.0
			if rtp <= checkRTP {
				switch i {
				case 70:
					k = 10.9
				case 71:
					k = 10.80
				case 72:
					k = 10.90
				case 73:
					k = 11.00
				case 74:
					k = 11.10
				case 75:
					k = 11.20
				case 76:
					k = 11.30
				case 77:
					k = 11.40
				case 78:
					k = 11.50
				case 79:
					k = 11.60
				}
				break
			}
		}
	}

	// Диапазон 0.8-0.9
	if rtp >= 0.80 && rtp < 0.90 {
		for i := 80; i <= 89; i++ {
			checkRTP := float64(i) / 100.0
			if rtp <= checkRTP {
				switch i {
				case 80:
					k = 16.9
				case 81:
					k = 12.0
				case 82:
					k = 14.0
				case 83:
					k = 16.0
				case 84:
					k = 18.0
				case 85:
					k = 20.0
				case 86:
					k = 22.0
				case 87:
					k = 25.0
				case 88:
					k = 28.0
				case 89:
					k = 30.0
				}
				break
			}
		}
	}

	// Диапазон 0.9-0.99
	if rtp >= 0.90 && rtp < 0.99 {
		for i := 90; i <= 98; i++ {
			checkRTP := float64(i) / 100.0
			if rtp <= checkRTP {
				switch i {
				case 90:
					k = 35.9
				case 91:
					k = 60
				case 92:
					k = 80
				case 93:
					k = 100
				case 94:
					k = 150
				case 95:
					k = 200
				case 96:
					k = 250
				case 97:
					k = 300
				case 98:
					k = 350
				}
				break
			}
		}
	}

	if rtp == 1.0 {
		rawMult := 1.0 + rand.Float64()*(10000.0-1.0)
		scale := 0.025 // маленький множитель, чтобы "сжать" к 1
		mult := 1.0 + (rawMult-1.0)*scale
		return mult

	}
	// Генерация мультипликатора
	rawMult := 1.0 + rand.Float64()*(10000.0-1.0)
	scale := math.Pow(rtp, k)
	mult := 1.0 + (rawMult-1.0)*scale
	return mult
}
