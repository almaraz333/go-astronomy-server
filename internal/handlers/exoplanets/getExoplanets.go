package handlers

import (
	"net/http"

	"github.com/almaraz333/go-astronomy-server/internal/exoplanets"
	"github.com/almaraz333/go-astronomy-server/internal/utils"
)

func GetExoplanets(w http.ResponseWriter, r *http.Request) {
	utils.MiddlewareChain(w, r, utils.Logging, func(w http.ResponseWriter, r *http.Request) error {

		query := `
		SELECT pl_name, hostname, discoverymethod, pl_orbper, pl_rade, pl_bmasse
		FROM ps
		WHERE pl_orbper > 300 AND pl_rade < 2
		ORDER BY pl_orbper
	`

		data, err := exoplanets.QueryExoplanets(query)

		if err != nil {
			return utils.WriteJSON(w, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		}

		return utils.WriteJSON(w, http.StatusOK, data)
	})
}
