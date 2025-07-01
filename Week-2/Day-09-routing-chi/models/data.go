package models

type EchoData struct {
	Data string `json:"data"`
}

type StatusResponse struct {
	Uptime string `json:"uptime"`
	Status string `json:"status"`
}
