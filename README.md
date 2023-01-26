# ShareBurn

## About

Sometimes you need a string from another machine and you don't want to log in to install or log in to another service.

### CURL

#### Post a string
`$ curl -X POST http://localhost:8092/s -H 'Content-Type: application/json' -d '{"string_value":"newstring"}'`
`Rovershadow`

#### Get a string
`$ curl localhost:8092/s/Rovershadow`
`newstring`

## Deploy

### Docker
 - docker build .
 - sudo docker run -d -p 0.0.0.0:8092:8092 local/shareburn

## Dev
### Local
 - go get .
 - go run .

## ROADMAP
 - Support encryption
 - Add files
 - Log in for larger file capability
 - cmd line app
 - delete strings after 1 hour