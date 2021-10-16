FROM golang:1.13 AS build

WORKDIR /compose/
COPY main.go main.go
RUN CGO_ENABLED=0 go build -o frontend main.go

FROM scratch
ARG STRING_LENGTH 
ARG FRONTEND_PORT 
ARG BACKEND_PORT 
COPY --from=build /compose/frontend /usr/local/bin/frontend
CMD ["/usr/local/bin/frontend"]
