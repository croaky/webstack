package api

import (
	"fmt"
	"net/http"
)

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	var col int
	s.db.QueryRow(r.Context(), "SELECT 1").Scan(&col)
	fmt.Fprintf(w, "{\"status\":\"ok\"}")
}
