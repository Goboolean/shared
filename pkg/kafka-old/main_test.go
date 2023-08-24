package kafka_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)



func TestMain(m *testing.M) {

	if err := os.Chdir("../../"); err != nil {
		panic(err)
	}

	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	
	code := m.Run()

	os.Exit(code)
}
