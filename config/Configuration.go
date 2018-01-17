package config

type Configuration struct {
	Harvest  HarvestConfiguration
	Database DatabaseConfiguration
	Sync     SyncConfiguration
}
