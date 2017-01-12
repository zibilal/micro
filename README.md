# micro-api
service wrapper for API.

### Requirements
  * [Golang](https://golang.org/doc/install)
  * [Glide](https://github.com/Masterminds/glide)

### Installation
  * `git clone https://github.com/maps90/micro-api.git path/to/src/github.com/mataharimall/micro-api`
  * `cd path/to/src/github.com/mataharimall/micro-api`
  * `glide install`
  * `cp config.yaml.dist config.yaml`
  * `go build`
  * `./micro-api`

### Running Test
  * `cd path/to/src/github.com/mataharimall/micro-api`
  * `go test -v $(go list ./... | grep -v "/vendor/")`
