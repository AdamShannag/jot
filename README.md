# Jot
Quickly generate microservices and related components

## Installation
```
go install github.com/AdamShannag/jot
```

## Commands
* `jot init`
    * > `jot init . myproject`
* `jot add`
    * > `jot add --service user --port 8080`
    * > `jot add --service user --port 8080 --rest --endpoints users,posts`