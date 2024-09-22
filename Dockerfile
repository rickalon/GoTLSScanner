FROM golang:1.23.1-bookworm AS build-stage
WORKDIR /app
COPY . .
RUN go mod download   
RUN CGO_ENABLED=0 GOOS=linux go build -o urltls ./main.go

FROM scratch AS build-release-stage
WORKDIR /app
COPY --from=build-stage /app/urltls .
ENV DB_PORT=5432
EXPOSE 8080
ENTRYPOINT ["./urltls"]