package cmd

import "fmt"

type Todo struct {
}

func (t *Todo) Help() string {
	return "todo help"
}
func (t *Todo) Run(args []string) int {
	fmt.Println("todo run")
	return 0
}
func (t *Todo) Synopsis() string {
	return "todo synopsis"
}
