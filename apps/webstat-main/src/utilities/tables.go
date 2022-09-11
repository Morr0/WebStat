package utilities

import "os"

func GetStatsTableName() string {
	name := os.Getenv("StatsTable")
	return name
}
