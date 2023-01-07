# Poofbin

## What
App for quickly moving a string (or eventually, a file) to a remote machine.

## Why
Sometimes you need a key and you don't want to log in to your email.

## HOW

### CURL

#### Post
`$ curl -X POST http://localhost:8091/poof -H 'Content-Type: application/json' -d '{"string_value":"newstring"}'`
`Rovershadow`

#### Get
`$ curl localhost:8091/poofs/Rovershadow`
`newstring`

