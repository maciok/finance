package bossa

import (
	log "github.com/sirupsen/logrus"
	"os"
)

type BossaProperty struct {
	BaseUrl string
}

func FromEnv() BossaProperty {
	baseUrl, exists := os.LookupEnv("BOSSA_BASE_URL")
	if !exists {
		log.Fatal("Cannot find base url for Bossa service")
	}

	return BossaProperty{baseUrl}
}
