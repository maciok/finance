package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	"pl/pietrzam/integration/bossa"
)


func main() {
	bossaProperty := bossa.FromEnv()
	metrics, err := bossa.RequestMetric(bossa.INVESTMENT_FUNDS, bossaProperty)

	if err != nil {
		logrus.Fatal("Cannot load %s", bossa.INVESTMENT_FUNDS, err)
	}

	logrus.Info(metrics)
}

