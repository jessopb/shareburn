# ShareBurn

## What
App for quickly moving a string (or eventually, a file) to a remote machine.

## Why
Sometimes you need a key and you don't want to log in to your email.

## How

### CURL

#### Post a string
`$ curl -X POST http://localhost:8091/s -H 'Content-Type: application/json' -d '{"string_value":"newstring"}'`
`Rovershadow`

#### Get a string
`$ curl localhost:8091/s/Rovershadow`
`newstring`

# ROADMAP
 - Simple post / get
 - Env for port etc
 - Add files
 - Log in for larger file capability
 - Frontend
 - cmd line app
 - delete strings after 1 hour