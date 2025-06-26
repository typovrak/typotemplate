package app

import (
	"errors"
	"log"
	"os"
)

func RequireEnv() error {
	envs := []string{
		"APP_GO_TEST",
	}

	for i := 0; i < len(envs); i++ {
		_, defined := os.LookupEnv(envs[i])
		if !defined {
			err := "environment variable not defined : " + envs[i]

			if os.Getenv("APP_GO_TEST") == "false" {
				log.Fatal(err)
			}

			return errors.New(err)
		}
	}

	if os.Getenv("APP_GO_TEST") == "false" {
		log.Println("All environment variables defined")
	}

	return nil
}
