package cli

import (
	log "github.com/Sirupsen/logrus"
	"github.com/mdy/melody/project"
	"github.com/mdy/melody/provider"
	"github.com/urfave/cli"
	"os"
)

func install(c *cli.Context) error {
	if len(c.Args()) != 0 {
		log.Fatalf("`install` command takes no arguments. See '%s install --help'.", c.App.Name)
	}

	wDir, _ := os.Getwd()
	project, err := project.Load(wDir)
	log.Info("Project", project, " -- ", err)
	if err != nil {
		return err
	}

	// Convert Project.Config to Requested
	source := provider.NewMelody(project.Locked)
	return project.UpdateWithBase(source, project.Locked)
}
