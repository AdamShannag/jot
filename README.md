# Jot
<img alt="jot" src="assets/jot.png" />
Introducing a tool for effortlessly creating projects following the micro-services architecture. This tool streamlines project generation, offering modular components and automated dependency management. It simplifies the complexities of micro-services development, empowering the delivery of high-quality, scalable projects with ease.

## Installation

```
go install -v github.com/AdamShannag/jot/v2/cli/cmd/jot@latest
```

## How to use this tool

Leverage `jot` for effortless project generation. Execute the necessary commands to create services, endpoints, and middlewares, while enjoying the simplicity and time-saving benefits.

### Creating a new project

Easily initiate a new project using `jot`. Just execute the following command:

```
jot init project-path project-name
```

You can also use a shorter version of the command if you don't need to specify the project name:

```
jot init project-path
```

This will create a new project directory at the specified path, and also generates a "jot.yaml" file, used by the tool itself to manage micro-services and their components.

### Creating a New Service with Endpoints and Middlewares

To create a new service within your project, follow these commands executed from the project directory that contains the "jot.yaml" file:

1. Create a simple service without any endpoints or middlewares:

```
jot add --service service-name --port service-port
```

2. Create or add endpoints to an existing service:

```
jot add --service service-name --port service-port --rest --endpoints first-endpoint,second-endpoint
```

3. Create or add middlewares to an existing service:

```
jot add --service service-name --port service-port --rest --middlewares first-middleware,second-middleware
```

4. Create or add both endpoints and middlewares simultaneously:

```
jot add --service service-name --port service-port --rest --endpoints first-endpoint,second-endpoint --middlewares first-middleware,second-middleware
```

A service can contain both endpoints and middlewares. Currently, you can generate RESTful endpoints or middlewares by utilizing the "--rest" flag in combination with the "--endpoints" or "--middlewares" flag.

Furthermore, you have the option to include the "--crud" flag to generate simple stub functions for REST operations on an endpoint or more.

```
jot add --service service-name --port service-port --rest --crud --endpoints first-endpoint,second-endpoint
```

Certain flags in the command have shorter versions for convenience:

- Instead of `--service`, you can use `--srv`.
- Instead of `--port`, you can use `--p`.
- Instead of `--endpoints`, you can use `--end`.
- Instead of `--middlewares`, you can use `--mid`.

For example:

```
jot add --srv service-name --p service-port --rest --end first-endpoint,second-endpoint --mid first-middleware,second-middleware
```

To explore additional commands and get more detailed information, you can utilize the help command:

```
jot help
```

This command provides a comprehensive overview of available commands, their usage, and any additional details you may need while working with Jot.
