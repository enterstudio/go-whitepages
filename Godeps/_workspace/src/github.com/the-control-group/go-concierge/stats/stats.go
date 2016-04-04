package stats

// StatHandler is for sending stats to a time-series database
type StatHandler interface {
	SendStats(name string, stats map[string]interface{}) error
}
