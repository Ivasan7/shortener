# Intro to URL shorter

In this repo there is a GO(Golang) CLI application which can shorten a URL. Alternative to the Bit.ly 
or well known google services.

## Basic operating principle
* Get URL
⋅⋅* Possible to get a standard URL in text form or open in the browser
⋅⋅* Get a short URL, text or browser
* Lengthen URL
⋅⋅* Look up short URL original names
* Shorten URL
⋅⋅* Generate short URL and store it in the DB

## User instructions
Assumes docker is installed on the host computer
1. Clone the repo
2. cd repo
3. docker run . -t <image name>
4. docker run -it <image name> /bin/bash
5. (cd code/shortener) -> supposed to start up in this directory
5. shortener -h
6. On Linux systems the binary can be copied out and used natively so the open in browser functinality is available.
  (docker cp...)

## Developer mode
1. Repear steps 1-3
2. docker run -it -v <abs/path/to/repo>:/code <image name> /bin/bash 
  (pontentially volume mount can also be used, here binding mount due to simplicity of use)
3. Run the tests
⋅⋅* cd /code/shortener/dbman
⋅⋅* go test dbman_test.go dbman.go
4. Run the app: go run main.go <cmd> <flags>
5. Build the app: go build -o <binary name> main.go
  
## Improvement points
1. More extensive tests
2. Application level itegration/regressions tests
3. Remove code duplication in some places
4. Remove the necessity of storing the short URL-s
⋅⋅* Long URL's can be found buy decoding the URI after fake.it/URI which provides ID of the original URL to be fetched
5. Cross compilation does not work correctly
⋅⋅* for mac do env GOOS="darwin" go build -o <name> main.go, however the application is poroducing error, likely incompatibility with the SQLlite3 driver.
