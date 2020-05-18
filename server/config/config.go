package config

import (
	"fmt"
	"os"
)

var DatabaseURL string

func init() {
	DatabaseURL = os.Getenv("DATABASE_URL") + "?parseTime=true&loc=Local"
	fmt.Println("=====", os.Getenv("DATABASE_URL"), "======")
}
