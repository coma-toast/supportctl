package system

// SmartData is the relevant SMART data for a disk
type SmartData struct {
	DeviceInfo   Device      `json:"device"`
	ModelFamily  string      `json:"model_family"`
	ModelName    string      `json:"model_name"`
	SerialNumber string      `json:"serial_number"`
	Status       SmartStatus `json:"smart_status"`
	PowerOnHours PowerOnTime `json:"power_on_time"`
	RotationRate string      `json:"rotation_rate"`
}

// Device is general device data
type Device struct {
	Name       string `json:"name"`
	InfoName   string `json:"info_name"`
	DeviceType string `json:"type"`
	Protocol   string `json:"protocol"`
}

// SmartStatus is the overall SMART status
type SmartStatus struct {
	Passed bool `json:"passed"`
}

// PowerOnTime is the power on hours
type PowerOnTime struct {
	Hours int64 `json:"hours"`
}
