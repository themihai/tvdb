package main

import (
	"fmt"

	"gopkg.in/urfave/cli.v2"
)

// Series looksup series by ID
func Series(c *cli.Context) error {
	var id string = c.Args().First()
	if len(id) == 0 {
		return fmt.Errorf("Id parameter is required")
	}

	series, err := client.Series.Get(id)
	if err != nil {
		return err
	}
	if series == nil {
		fmt.Fprintf(c.App.Writer, "Series ID %s could not be found", id)
		return nil
	}

	fmt.Fprint(c.App.Writer, series)
	return nil
}
