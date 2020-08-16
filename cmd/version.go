package cmd

import "fmt"

const (
	version = "0.0.0"
	commit  = "n/a"
)

func Vers() string {
	return version
}

func Commit() string {
	return commit
}

type Version struct {
}

func (v *Version) Help() string {
	return "print version and exit"
}
func (v *Version) Run(args []string) int {
	fmt.Printf("%s (%s)\n", Vers(), Commit())
	return 0
}
func (v *Version) Synopsis() string {
	return v.Help()
}
