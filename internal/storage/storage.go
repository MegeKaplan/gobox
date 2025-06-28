package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/MegeKaplan/gobox/internal/models"
)

func GetConfigDir() (string, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(userConfigDir, "gobox"), nil
}

func GetPackagesFilePath() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "packages.json"), nil
}

func Init() error {
	configDir, err := GetConfigDir()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(configDir, 0644); err != nil {
		return err
	}

	packagesFilePath, err := GetPackagesFilePath()
	if err != nil {
		return err
	}

	if _, err := os.Stat(packagesFilePath); os.IsNotExist(err) {
		file, err := os.Create(packagesFilePath)
		if err != nil {
			return err
		}
		defer file.Close()

		if _, err = file.WriteString("[]"); err != nil {
			return err
		}
	}

	return nil
}

func LoadPackages() ([]models.Package, error) {
	packagesFilePath, err := GetPackagesFilePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(packagesFilePath)
	if err != nil {
		return nil, err
	}

	var packages []models.Package
	if err := json.Unmarshal(data, &packages); err != nil {
		return nil, err
	}

	return packages, nil
}

func FindPackage(name string) (models.Package, bool, error) {
	packages, err := LoadPackages()
	if err != nil {
		return models.Package{}, false, err
	}

	for _, pkg := range packages {
		if pkg.Name == name {
			return pkg, true, nil
		}
	}

	return models.Package{}, false, nil
}

func SavePackage(packageName string) error {
	packages, err := LoadPackages()
	if err != nil {
		return err
	}

	now := time.Now()

	for i, pkg := range packages {
		if pkg.Name == packageName {
			packages[i].UsageCount++
			packages[i].LastUsed = now
			return SaveAllPackages(packages)
		}
	}

	newPkg := models.Package{
		Name:        packageName,
		UsageCount:  1,
		LastUsed:    now,
		InstalledAt: now,
	}

	packages = append(packages, newPkg)

	return SaveAllPackages(packages)
}

func RemovePackage(name string) error {
	packages, err := LoadPackages()
	if err != nil {
		return err
	}

	newList := []models.Package{}
	for _, pkg := range packages {
		if pkg.Name != name {
			newList = append(newList, pkg)
		}
	}

	return SaveAllPackages(newList)
}

func SaveAllPackages(pkgs []models.Package) error {
	packagesFilePath, err := GetPackagesFilePath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(pkgs, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(packagesFilePath, data, 0644)
}
