package data

import (
	"encoding/json"
	"os"
	"sync"
)

var (
	GeoData map[string]interface{}
	GeoLock sync.RWMutex
)

func LoadGeoJSON(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var geo map[string]interface{}
	if err := json.Unmarshal(bytes, &geo); err != nil {
		return err
	}

	GeoLock.Lock()
	GeoData = geo
	GeoLock.Unlock()

	return nil
}
