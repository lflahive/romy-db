package config

var (
	Configuration Config
)

type Config struct {
	StoragePath string
}

func NewConfig(storagePath string) {
	Configuration = Config{StoragePath: storagePath}
}
