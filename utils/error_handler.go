package utils

import (
	"log"
)

func LogError(err error) {
	// Implementasi pencatatan error ke file log
	log.Println(err)
}
