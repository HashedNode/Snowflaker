package services

import (
	"crystal_snowflake/generators"
	"log"
	"net/http"
)

func ServeSnowflakeId(w http.ResponseWriter, r *http.Request) {
	log.Printf("Arrived request for generate snowflake id: %s %s", r.Method, r.URL.Path)
	snowflakeId := generators.SnowflakeNode.GenerateSnowflakeId()
	w.Header().Add("X-Snowflake-Id", snowflakeId.ToString())
	w.WriteHeader(http.StatusOK)
}
