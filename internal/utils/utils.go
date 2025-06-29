package utils

import (
	"os"
	"sort"

	"github.com/MegeKaplan/gobox/internal/models"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func SortPackages(packages *[]models.Package,sortBy string, ascending bool) []models.Package {	
	p := *packages

	sort.Slice(p, func(i, j int) bool {
		var less bool

		switch sortBy {
			case "name":
				less = p[i].Name < p[j].Name
			case "usage_count":
				less = p[i].UsageCount < p[j].UsageCount
			case "last_used":
				less = p[i].LastUsed.Before(p[j].LastUsed) // (is date i before date j) ? true : false
			case "installed_at":
				less = p[i].InstalledAt.Before(p[j].InstalledAt) // (is date i before date j) ? true : false
			default:
				less = p[i].Name < p[j].Name
		}

		if ascending {
			return less
		}
	
		return !less
	})

	return p
}