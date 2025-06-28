package models

import "time"

type Package struct {
	Name        string    `json:"name"`
	UsageCount  int       `json:"usage_count"`
	LastUsed    time.Time `json:"last_used"`
	InstalledAt time.Time `json:"installed_at"`
}
