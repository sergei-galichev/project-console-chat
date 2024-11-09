package config

import "github.com/joho/godotenv"

// LoadConfig loads config from path if env variables are not set from Docker
func LoadConfig(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}
