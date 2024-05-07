package model

import "fmt"

type UpdateSchedule struct {
	Interval string `json:"interval"`
}

type DependencyUpdate struct {
	PackageEcosystem string         `json:"package_ecosystem"`
	Directory        string         `json:"directory"`
	Schedule         UpdateSchedule `json:"schedule"`
}

type DependabotConfig struct {
	Version int                `json:"version"`
	Updates []DependencyUpdate `json:"updates"`
}

func (d *DependabotConfig) String() string {
	return fmt.Sprintf("DependabotConfig{Version: %d, Updates: %v}", d.Version, d.Updates)
}
