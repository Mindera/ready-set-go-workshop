![workshop logo](https://scontent.fopo2-1.fna.fbcdn.net/v/t1.0-9/46185393_10155995013962358_3040697147461730304_o.jpg?_nc_cat=103&_nc_ht=scontent.fopo2-1.fna&oh=28ec97a0d2ef8fd6a5d0d393d41dd5f9&oe=5C7F40C0)
# ready, set, go!
This workshop is aimed to provide a fast integration with Go.

## Goals
- Install and configure Go;
- Go basics;
- Develop a web application that access a datastore.

## Workshop resources
* Chat room: https://tlk.io/mindera
* Download Go: https://golang.org/dl/
* Presentation:

## Requirements
* Git
* Text Editor or IDE (VSCode with Go extension recommended)
* Docker (recommended)

## Running the project
### Starting datastore
On the root folder you will need to start redis container with docker-compose
```bash
$ docker-compose up -d
```
Redis will be available on `localhost:6379`.

### Executing source code
On the root folder you need to run the execute script
```bash
$ ./scripts/execute.sh
```

### Building application
On the root folder you need to execute the build script
```bash
$ ./scripts/build.sh
```
Executable will be available on `app/build/webapp`.

### Executing application
First you need to execute the `build` script. Then, on the root folder you need to execute the run script
```bash
$ ./scripts/run.sh
```

### Building a docker image with Executable
On the root folder you need to execute docker build command
```bash
$ docker build -t webapp .
# after it has created the image, execute it
$ docker run -it --rm -p 8080:8080 webapp
```
Browse to http://localhost:8080/health and you should receive `health` endpoint response.

## Workshop timeline summary
* Presentation (40/45min)
* Setup (10/15min)
* Development of web application (45min)
* Questions (10/15min)

## Workshop challenge
### Setup
* Install Go
* Make sure Go is working by testing `go` command in your terminal
* Git clone this project
* Compile and execute application
* Pause and check if everyone is one the same step
* Go to http://localhost:8080/health. You should get a `Connection refused` error.

### Development
* Complete the code following steps presented below. Shout any questions, if you don't understand some/any parts.

#### Step a)
On `app/cmd/webapp/main.go` you have to initialise the redis client on the init function. If you've used `docker-compose.yml` to start redis service, use `localhost:6379` (address:port) to connect to redis.
#### Step b)
On `app/cmd/web/main.go` you need to finish `http.Server` implementation. Head to [http docs](https://golang.org/pkg/net/http/) and search for a function that allows you to listen and serve content. You should also check for errors.
#### Step c.1)
On `app/internal/webapp/routes.go` you need to complete `RegisterRoutes()` and add one http route (PUT) for any given path.
#### Step c.2)
On `app/internal/webapp/handler.go` you need to implement a function that given a key (request URL path) and a value (request body), it should add it to the datastore. If it's successful return the request body, if not `http.StatusBadRequest` (header).
##### Step c.3)
On `app/internal/webapp/routes.go` you need to update `RegisterRoutes()` and add the implemented function on c.2 as a handler of `PUT`.
#### Step d)
On `app/internal/webapp/handler.go` you need to complete `health` function: return `{"status": "OK"}` if connection with datastore if ok; return `{"status": "FAIL", "error": errorMessage}`.
#### Step e.1)
On `app/internal/webapp/handler.go` you should implement a logger handler and it should send to stdout:
- `{request end date} | {request Method} | {time to execute in msec} | {request path}`

As an example:
`2018/11/20 18:50:12 | GET |    0.01344ms /health`
#### Step e.2)
On `app/internal/webapp/routes.go` you should add the logger handler implemented on e.1) to each mux handlers.

---
## Expected results
* `GET /health` endpoint should return `OK` if datastore is up and running; and should return `FAIL` if datastore is down.
* `GET /{key}` endpoint should return the value of the `{key}` on the datastore.
* `PUT /{key}` endpoint should accept a `{body}`, save on the datastore and on success should return it.
### BONUS
* All `http` methods (`GET` and `PUT`) should be logged to `stdout`.
