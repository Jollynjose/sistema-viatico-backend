package helpers

import "github.com/Jollynjose/sistema-viatico-backend/internal/config"

func IsProduction(cfg *config.Config) bool {
	return cfg.ENV == config.PRODUCTION
}
