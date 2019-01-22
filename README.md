# exampleservice

Exampel service shows how a todo rest api can be implemented in golang, with an idiomatic and modular codebase.

### To Run

From the exampleservice directory, use the make file.

Build (creates bin directory with `exampleservice` binary):

```
make
```

Run all tests:

```
make test
```

### To Dockerize:

Build docker image:

```
docker build . -t <image-name>
```

Run docker image and map port 8080:

```
docker run -P <image-name>
```

- Test out the basic routes in your browser at: localhost:8080 {eg http://localhost:8080/v1/api/user}
