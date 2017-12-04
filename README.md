# Kronos - a docker scheduler
An API and CLI in Go for a cron-like task scheduler for running recurring, short-lived Docker containers.

## Usage
1. Download dependencies
2. Build desired images (simple image that prints to console is included)
3. Build and start API
4. Build/Install CLI
5. Try some commands out!

### Install dependencies
```
go get "github.com/gorilla/mux"
go get "github.com/parnurzeal/gorequest"
go get "github.com/urfave/cli"
go get "gopkg.in/robfig/cron.v2"
```

### Demo image
```
cd .../kronos
docker build -t selfprintingimage .
```

### Start API
```
cd .../kronos/kronos_api
go build
./kronos_api
```

### Use CLI
Supports `list`, `create`, and `delete` operations.

```
cd .../kronos/kronos
go install
kronos_cli list http://localhost:8080
kronos_cli create http://localhost:8080 selfprintingimage "0-59/5 * * * * *"
kronos_cli create http://localhost:8080 selfprintingimage "0-59/5 * * * * *"
kronos_cli delete http://localhost:8080 1

# kronos_cli <command> <host IP> <args...>
# kronos_cli list <host IP>
# kronos_cli create <host IP> <image Tag> "<frequency>"
# kronos_cli delete <host IP> <image Tag> <job id>
```

## Future enhancements

### More expressive docker commands
As is, the scheduler simply runs the container without any extra input. It'd be useful to be able to pass in the full set of `docker run ...` command set, allowing users to set environment variables or anything else they'd want.

### Error handling
Right now there's a good chance you'll break things if you enter a bad command in the CLI or if the API runs a bad image. I'd want to add some checks around deleting jobs on the API. The CLI could use some improvements for checking the arguments before trying to run them.

### CLI help
There's a small bit of boilerplate in the CLI to help users know how to use the CLI, but it could use a lot more. I'd like to flesh out help for args for the commands to help guide the user. I'd want to also add the ability to save a host, so you don't have to type in the IP for each command. It'd also be cool to add some auto-completion for commands. 

### Frequency help
Most people are familiar with schdeuling cron jobs, but it'd be great to add some verification or help within the CLI. Even something as sample as a few example inputs. Right now it just takes a string and passes it along to the API, which isn't great.

### WebHooks
Adding a callback for the API on the CLI's local machine with the output of the docker task would be be a helpful addition.
