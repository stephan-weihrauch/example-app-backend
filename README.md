# example-app-backend

This is a simple backend service written in GoLang. It provides an HTTP Rest API. The API specification can be found here: [api/openapi.yaml](api/openapi.yaml).

## Build on Docker

This folder contains a [Dockerfile](./Dockerfile). This can be used to build the application and bake it into an image.

## Run on Docker

Here is an example command to run the backend on Docker:

```sh
docker container run --name backend -p 8080:8080-d backend:v1.0.0
```

## Verify deployment

```sh
curl localhost:8080/api/version
```

This command should return a response like this:

```json
{ "version": "1.0.0" }
```
