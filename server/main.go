package main

import (
	"github.com/YUDAI-HIZU/gin-api/server/database"
	"github.com/YUDAI-HIZU/gin-api/server/handler"
)

func main() {
	database.GormConnect()
	handler.ListenAndServe()
}
