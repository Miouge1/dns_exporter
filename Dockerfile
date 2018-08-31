# compiler container
FROM golang as build
WORKDIR /go
ADD main.go .
RUN CGO_ENABLED=0 go build -o /main

# runtime container
FROM scratch
COPY --from=build /main /main
EXPOSE 9100
ENTRYPOINT ["/main"]
