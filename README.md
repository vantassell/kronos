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

### Start api
```
cd .../kronos/kronos_api
go build
./kronos_api
```

### Use cli
Supports `list`, `create`, and `delete` operations.

```
cd .../kronos/kronos
go install
# kronos <command> <host IP> <args...>
kronos_cli list http://localhost:8080
kronos_cli create http://localhost:8080 selfprintingimage "0-59/5 * * * * *"
kronos_cli create http://localhost:8080 selfprintingimage "0-59/5 * * * * *"
kronos_cli delete http://localhost:8080 1

# kronos_cli list <host IP>
# kronos_cli create <host IP> <image Tag> "<frequency>"
# kronos_cli delete <host IP> <image Tag> <job id>
```
