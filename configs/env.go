package config

import (
	"os"
	
)



func GetEnv(key , fallbackValue string) string{
	if value, done := os.LookupEnv(key); done{

		return value
	}

	return fallbackValue
}
