package actions

import (
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli"
)

func Patch(scanPath *string) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		errMsg := "Method not yet implemented!"
		log.Error(errMsg)
		return errors.New(errMsg)
	}
}
