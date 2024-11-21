package cmd

import todo "togo-list/internal"

type AddCommand struct {
	Task string
}

func (c *AddCommand) Execute() error {
	todo.AppendTask(c.Task)
	return nil
}
