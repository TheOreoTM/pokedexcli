package main

import "os"

func runExit(conf *config, args ...string) error {
	os.Exit(0)

	return nil
}
