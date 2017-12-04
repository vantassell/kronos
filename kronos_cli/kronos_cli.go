package main

import (
	"fmt"
	"log"
	"os"
	"github.com/urfave/cli"
	"github.com/parnurzeal/gorequest"
	"encoding/json"
	"text/tabwriter"
)

// models
type Task struct {
	ID int `json:"id,omitempty"`
	Image *Image `json:"image"`
	Frequency string `json:"frequency"`
}

type Image struct {
	Tag  string `json:"tag"`
}

func main() {
	app := cli.NewApp()
	app.Name = "kronos"
	app.Usage = "Schedule and run short-lived Docker containers."

	app.Commands = []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "Schedule a task to be run",
			Action: func(c *cli.Context) error {

				// prepare request
				host := fmt.Sprintf("%s/tasks", c.Args().Get(0))
				tag := c.Args().Get(1)
				frequency := c.Args().Get(2)
				request := gorequest.New()
				body := fmt.Sprintf("{ \"image\": {\"tag\": \"%s\"}, \"frequency\": \"%s\"}",tag, frequency)
				
				// fire request
				_, body, errs := request.Post(host).Send(body).End()
				for _, err := range errs {
					check(err)
				}

				// check results
				var tasks []Task
				err := json.Unmarshal([]byte(body), &tasks)
				check(err)

				// print results
				printTasks(tasks)

				return nil
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "Delete a scheduled task",
			Action: func(c *cli.Context) error {

				// prepare request
				host := fmt.Sprintf("%s/%s/%s", c.Args().Get(0), "tasks", c.Args().Get(1))
				request := gorequest.New()
				
				// fire request
				_, body, errs := request.Delete(host).End()
				for _, err := range errs {
					check(err)
				}

				// check results
				var tasks []Task
				err := json.Unmarshal([]byte(body), &tasks)
				check(err)

				// print results
				printTasks(tasks)

				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "Show all scheduled tasks",
			Action: func(c *cli.Context) error {
				
				// prepare request
				host := fmt.Sprintf("%s/tasks", c.Args().Get(0))
				request := gorequest.New()
				
				// fire request
				_, body, errs := request.Get(host).End()
				for _, err := range errs {
					check(err)
				}

				// check results
				var tasks []Task
				err := json.Unmarshal([]byte(body), &tasks)
				check(err)

				// print results
				printTasks(tasks)

				return nil
			},
		},
	}

	app.Run(os.Args)
}

func printTasks(tasks []Task) {
	w := tabwriter.NewWriter(os.Stdout, 4, 1, 2, ' ', 0)
	fmt.Println("-")
	fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%s", "id", "tag", "frequency"))
	for _, v := range tasks {
		fmt.Fprintln(w, fmt.Sprintf("%v\t%s\t%s", v.ID, v.Image.Tag, v.Frequency))
	}
	w.Flush()
	fmt.Println("-")
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}