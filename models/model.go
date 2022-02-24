package models

import (
	"math/rand"
	"os"
	"text/tabwriter"
	"time"
)

type Water struct {
	Status  int
	Message string
}

type Wind struct {
	Status  int
	Message string
}

type Status struct {
	Water Water
	Wind  Wind
}

func GetStatus() Status {
	var result Status
	// fmt.Println("Masuk Model")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	defer w.Flush()

	waterRand := r.Intn(100)
	windRand := r.Intn(100)

	result.Water.Status = waterRand
	result.Wind.Status = windRand

	if waterRand < 5 {
		result.Water.Message = "aman"
	} else if waterRand > 8 {
		result.Water.Message = "bahaya"
	} else {
		result.Water.Message = "siaga"
	}

	if windRand < 6 {
		result.Wind.Message = "aman"
	} else if windRand > 15 {
		result.Wind.Message = "bahaya"
	} else {
		result.Wind.Message = "siaga"
	}

	return result
}
