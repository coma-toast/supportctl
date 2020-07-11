package drivefinder

// TableData is the struct for data for Tables
type TableData struct {
	Drive       string
	Type        string
	SSD         bool
	Serial      string
	SerialPath  string
	Hours       int64
	SMART       string
	ZPOOLErrors int32
}
