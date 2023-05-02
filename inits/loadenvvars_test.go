package inits

import (
	"testing"

	"github.com/joho/godotenv"
)

func Test_Loadenvvars(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fail()
		t.Log(err)
	}
}
