package utilities

import "os"

func GetStatsTableName() (name string) {
	name = os.Getenv("StatsTable")

	return
}
