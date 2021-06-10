package common

import (
	"encoding/json"
	"net/http"
)

// HTTPResponse is the common response format
type HTTPResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorData  `json:"error,omitempty"`
}

// ErrorData is the data format returned in case of error
type ErrorData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

//GenerateOKSuccessResponse is helper to generate success response
func GenerateOKSuccessResponse(w http.ResponseWriter, message string, data interface{}) {
	resp := HTTPResponse{
		Message: message,
		Data:    data,
		Error:   nil,
	}
	bytes, err := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		GenerateFailedResponse(w, "Server error.", err)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Write(bytes)
}

//GenerateFailedResponse is helper to generate failed response
func GenerateFailedResponse(w http.ResponseWriter, message string, err error) {
	statusCode := http.StatusBadRequest
	var errResp ErrorData
	if err != nil {
		errCode := "400"
		errMsg := err.Error()
		errResp = ErrorData{
			Code:    errCode,
			Message: errMsg,
		}
	}
	resp := HTTPResponse{
		Message: message,
		Data:    nil,
		Error:   &errResp,
	}
	bytes, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(bytes)
}
