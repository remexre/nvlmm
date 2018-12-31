package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path"

	"github.com/docker/docker/pkg/mount"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "nvlmm"
	app.Usage = "New Vegas Linux Mod Manager"
	app.Commands = []cli.Command{
		{
			Name:   "mount",
			Usage:  "Mounts the OverlayFS",
			Action: doMount,
		},
		{
			Name:   "unmount",
			Usage:  "Unmounts the OverlayFS",
			Action: doUnmount,
		},
		{
			Name:   "setup",
			Usage:  "Performs first-time setup",
			Action: doSetup,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func doSetup(c *cli.Context) error {
	user, err := user.Current()
	if err != nil {
		return err
	}

	err = os.MkdirAll(path.Join(user.HomeDir, ".nvlmm/profiles/Default"), 0750)
	if err != nil {
		return err
	}

	err = os.MkdirAll(path.Join(user.HomeDir, ".nvlmm/workdir"), 0750)
	if err != nil {
		return err
	}

	_, err = os.Stat(path.Join(user.HomeDir, ".steam/steam/steamapps/common/Fallout New Vegas/FalloutNV.exe"))
	if err != nil {
		log.Println("Warning: Couldn't find FalloutNV.exe -- is Fallout: New Vegas installed?")
	}

	return nil
}

func doMount(c *cli.Context) error {
	profile := "Default" // TODO Support selecting a profile

	user, err := user.Current()
	if err != nil {
		return err
	}

	dataDir := path.Join(user.HomeDir, ".steam/steam/steamapps/common/Fallout New Vegas/Data")
	profileDir := path.Join(user.HomeDir, ".nvlmm/profiles", profile)
	workDir := path.Join(user.HomeDir, ".nvlmm/workdir")

	opts := fmt.Sprintf("lowerdir=%s,upperdir=%s,workdir=%s", dataDir, profileDir, workDir)
	return mount.Mount("overlay", dataDir, "overlay", opts)
}

func doUnmount(c *cli.Context) error {
	user, err := user.Current()
	if err != nil {
		return err
	}

	dataDir := path.Join(user.HomeDir, ".steam/steam/steamapps/common/Fallout New Vegas/Data")

	return mount.Unmount(dataDir)
}
