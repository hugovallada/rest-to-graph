package response

import "time"

type ApplicationStatusResponse struct {
	Message   string `json:"message"`
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"`
}

func New() ApplicationStatusResponse {
	return ApplicationStatusResponse{
		Message:   "ok",
		Version:   "1.0",
		Timestamp: time.Now().String(),
	}
}
