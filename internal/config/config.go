package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = "gatorconfig.json"

type Config struct {
	DB_URL string `json:"db_url"`
	DB_USR string `json:"current_username"`
}

// Resolve full path to config file
func resolveConfigPath(fileName string) (string, error) {

	// get users home directory to resolve absolute path to filename
	home_dir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return (home_dir + "/bootDev/" + fileName), nil

}

// write config settings to config file
func writeConfig(cfg Config) error {

	// resolve full filepath
	full_filepath, err := resolveConfigPath(configFileName)
	if err != nil {
		return err
	}

	// open json file
	file, err := os.OpenFile(full_filepath, os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	// create json encoder
	encoder := json.NewEncoder(file)

	encoder.SetIndent("", "	")

	// encode struct into file
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil

}

// set config struct username field and write struct to config json files
func (cfg *Config) SetUser(username string) error {

	// set struct username field
	cfg.DB_USR = username

	return writeConfig(*cfg)
}

// Read JSON config file info into Config struct
func ReadConfig() (*Config, error) {

	// resolve full filepath
	full_filepath, err := resolveConfigPath(configFileName)
	if err != nil {
		return &Config{}, err
	}
	fmt.Printf("Reading config file %s\n", full_filepath)

	// attempt to open file and read contents into byte slice
	file_contents, err := os.ReadFile(full_filepath)
	if err != nil {
		return &Config{}, err
	}

	// create empty container for config struct
	config := Config{}

	// unmarshal JSON file data into config struct
	err = json.Unmarshal(file_contents, &config)
	if err != nil {
		return &Config{}, err
	}

	return &config, err
}
