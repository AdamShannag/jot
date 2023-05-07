# Jot
Quickly generate microservices and related components

## Installation
```
go install -v github.com/AdamShannag/jot/cmd/jot@latest
```

## Commands
* `jot init`
    * > `jot init`
    * > `jot init ./myproject`
    * > `jot init ./myproject projectName`
* `jot add`
    * > `jot add --service user --port 8080`
    * > `jot add --service user --port 8080 --rest --endpoints users,posts`
