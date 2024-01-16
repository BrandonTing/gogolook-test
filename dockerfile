FROM golang:latest
WORKDIR /project_name
COPY go.mod ./
RUN go mod download
COPY . .
RUN make build
EXPOSE 8080
CMD ["./main"]
