package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/Saiteja93/Students-Api-Go/internal/types"
	"github.com/Saiteja93/Students-Api-Go/internal/utils/response"
)

func new() http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		var student types.Student
		 err := json.NewDecoder(r.Body).Decode(&student)
		 if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return 
		 }
		slog.Info("craeting student data")

		response.WriteJson(w, http.StatusCreated, map[string]string{"success":"OK"})
	}
}