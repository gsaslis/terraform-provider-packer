package main

import (
    	"log"
)

type Config struct {
	WorkDir  string
	JsonPath string
}

func (c *Config) New() (Config, error) {
	conf := Config{ WorkDir: c.WorkDir, JsonPath:c.JsonPath}

	log.Printf("[INFO] Packer configured for dir: %s, with JSON path: %s", c.WorkDir, c.JsonPath)
	return conf, nil
}


/*

// Client() returns a new Service for accessing Heroku.
//
func (c *Config) Client() (*heroku.Service, error) {
	service := heroku.NewService(&http.Client{
		Transport: &heroku.Transport{
			Username:  c.Email,
			Password:  c.APIKey,
			UserAgent: heroku.DefaultUserAgent,
		},
	})

	log.Printf("[INFO] Heroku Client configured for user: %s", c.Email)

	return service, nil
}*/
