package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func mockMissingParameterSubjectResponse() *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()

	recorder.Header().Set("Content-Type", "application/json")
	recorder.WriteHeader(http.StatusBadRequest) // 400

	// Write the response body
	response := map[string]string{"message": "Missing parameter subject"}
	json.NewEncoder(recorder).Encode(response)

	return recorder
}
