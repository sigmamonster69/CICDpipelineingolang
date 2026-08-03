FROM golang:1.22-alpine AS build

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/cicd-dashboard ./cmd/cicd-dashboard

FROM alpine:3.20

WORKDIR /app
COPY --from=build /out/cicd-dashboard /app/cicd-dashboard

EXPOSE 8080
ENTRYPOINT ["/app/cicd-dashboard"]
