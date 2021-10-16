## Composer Sosivio application

Prerequisite: Ubuntu 18.04 / Docker / docker-compose

Project structure:
```
.
├── backend
│   ├── Dockerfile
│   └── main.go
├── docker-compose.yml
├── frontend
│   ├── Dockerfile
│   └── main.go
└── README.md
```

[_docker-compose.yaml_](docker-compose.yaml)
```
version: "3.7"
services:
  frontend:
    build: frontend
    args:
     STRING_LENGTH = ${STRING_LENGTH}
     FRONTEND_PORT = ${FRONTEND_PORT}
     BACKEND_PORT = ${BACKEND_PORT}
    ports:
      - ${FRONTEND_PORT}:${FRONTEND_PORT}
    depends_on:
      - backend
  backend:
    build: backend
    args:
     THREADS = ${THREADS}
     BACKEND_PORT = ${BACKEND_PORT}
    ports:
      - ${BACKEND_PORT}:${BACKEND_PORT}
```
The compose file defines an application with two services `frontend` and `backend`.
When deploying the application, docker-compose map TCP ports 8888 and 9999 for the frontend and backend services containers, to the same port of the host as specified in the .env file.
Make sure ports 8888 and 9999 on the host is not already in use.
The .env also define the number of threads the backend can run.

## Deploy with docker-compose

```
$ docker-compose up -d
Building backend
Sending build context to Docker daemon  7.469MB
Step 1/10 : FROM golang:1.13 AS build
1.13: Pulling from library/golang
d6ff36c9ec48: Pull complete
c958d65b3090: Pull complete
edaf0a6b092f: Pull complete
80931cf68816: Pull complete
813643441356: Pull complete
799f41bb59c9: Pull complete
16b5038bccc8: Pull complete
Digest: sha256:8ebb6d5a48deef738381b56b1d4cd33d99a5d608e0d03c5fe8dfa3f68d41a1f8
Status: Downloaded newer image for golang:1.13
 ---> d6f3656320fe
Step 2/10 : WORKDIR /compose/
 ---> Running in c9e57c44869d
Removing intermediate container c9e57c44869d
 ---> fd9a9e150b8e
Step 3/10 : COPY main.go main.go
 ---> 5662d6dffb3d
Step 4/10 : RUN CGO_ENABLED=0 go build -o backend main.go
 ---> Running in c07b9e736492
Removing intermediate container c07b9e736492
 ---> f70a093df447
Step 5/10 : FROM scratch
 --->
Step 6/10 : ARG STRING_LENGTH
 ---> Running in 7b691a77b1f9
Removing intermediate container 7b691a77b1f9
 ---> 878ef789ba34
Step 7/10 : ARG FRONTEND_PORT
 ---> Running in 3f1c154c6fd0
Removing intermediate container 3f1c154c6fd0
 ---> c587e893e66e
Step 8/10 : ARG BACKEND_PORT
 ---> Running in 977e8cd92478
Removing intermediate container 977e8cd92478
 ---> 6cf4de3ee717
Step 9/10 : COPY --from=build /compose/backend /usr/local/bin/backend
 ---> e38df0efdb4d
Step 10/10 : CMD ["/usr/local/bin/backend"]
 ---> Running in aa99cc5fb242
Removing intermediate container aa99cc5fb242
 ---> 91484cc90f49
Successfully built 91484cc90f49
Successfully tagged project_backend:latest
WARNING: Image for service backend was built because it did not already exist. To rebuild this image you must use `docker-compose build` or `docker-compose up --build`.
Building frontend
Sending build context to Docker daemon   7.57MB
Step 1/10 : FROM golang:1.13 AS build
 ---> d6f3656320fe
Step 2/10 : WORKDIR /compose/
 ---> Using cache
 ---> fd9a9e150b8e
Step 3/10 : COPY main.go main.go
 ---> 83b83c1ece4c
Step 4/10 : RUN CGO_ENABLED=0 go build -o frontend main.go
 ---> Running in bf6306b3d747
Removing intermediate container bf6306b3d747
 ---> ee21b52d42a0
Step 5/10 : FROM scratch
 --->
Step 6/10 : ARG STRING_LENGTH
 ---> Using cache
 ---> 878ef789ba34
Step 7/10 : ARG FRONTEND_PORT
 ---> Using cache
 ---> c587e893e66e
Step 8/10 : ARG BACKEND_PORT
 ---> Using cache
 ---> 6cf4de3ee717
Step 9/10 : COPY --from=build /compose/frontend /usr/local/bin/frontend
 ---> 128fcb73de4f
Step 10/10 : CMD ["/usr/local/bin/frontend"]
 ---> Running in e11ae2fd8aa7
Removing intermediate container e11ae2fd8aa7
 ---> 2e2ed4fe7131
Successfully built 2e2ed4fe7131
Successfully tagged project_frontend:latest
WARNING: Image for service frontend was built because it did not already exist. To rebuild this image you must use `docker-compose build` or `docker-compose up --build`.
Creating project_backend_1 ... done
Creating project_frontend_1 ... done
```

## Expected result

Listing containers must show two containers running and the port mapping as below:
```
? docker ps
CONTAINER ID   IMAGE              COMMAND                  CREATED          STATUS          PORTS     NAMES
3f9974ac645b   project_frontend   "/usr/local/bin/fron…"   33 seconds ago   Up 33 seconds             project_frontend_1
03163f174072   project_backend    "/usr/local/bin/back…"   33 seconds ago   Up 33 seconds             project_backend_1
```

After the application starts, pick a random number of string you wish to encrypt (for example 5) and navigate to `http://localhost:8888/?5` in your web browser or run:
```
$ curl "http://localhost:8888/?5"
47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFXjsMRCmPwcFJr79MiZb7kkJ65B5GSbk0yklZkbeFK4VeOwxEKY/BwUmvv0yJlvuSQnrkHkZJuTTKSVmRt4UrhV47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFXZyRTLFm0zLhhj1yqIeSROFInlnjkNUOjEdTEQ2snG1g==
```

Stop and remove the containers
```
$ docker-compose down
```
