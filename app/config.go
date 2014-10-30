package app

type appConfig struct {
	DbContainerName string
}

var (
	Config = appConfig{
		DbContainerName: "/vagrant_db_1",
	}
)
