# ShareBurn

## What
App for quickly moving a string (or eventually, a file) to a remote machine.

## Why
Sometimes you need a key and you don't want to log in to your email.

## How
### CURL

#### Post a string
`$ curl -X POST http://localhost:8092/s -H 'Content-Type: application/json' -d '{"string_value":"newstring"}'`
`Rovershadow`

#### Get a string
`$ curl localhost:8092/s/Rovershadow`
`newstring`

# Deploy
## Docker
 - docker build .
 - sudo docker run -d -p 0.0.0.0:8092:8092 local/shareburn

# ROADMAP
 - Support encryption
 - Add files
 - Log in for larger file capability
 - cmd line app
 - delete strings after 1 hour