package config

import (
	"db"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Parser must implement ParseJSON
type Parser interface {
	ParseJSON([]byte) error
}

// configuration contains the application settings
type Configuration struct {
	Database db.Info `json:"database"`
	Port     int     `json:"httpport"`
}

// ParseJSON unmarshals bytes to structs
func (c *Configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

// Load the JSON config file
func Load(configFile string, p Parser) {
	var err error
	var absPath string
	var input = io.ReadCloser(os.Stdin)
	if absPath, err = filepath.Abs(configFile); err != nil {
		log.Fatalln(err)
	}

	if input, err = os.Open(absPath); err != nil {
		log.Fatalln(err)
	}

	// Read the config file
	jsonBytes, err := ioutil.ReadAll(input)
	input.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// Parse the config
	if err := p.ParseJSON(jsonBytes); err != nil {
		log.Fatalln("Could not parse %q: %v", configFile, err)
	}
}
