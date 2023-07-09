# ------------- BUILD --------------- #
FROM golang:1.20 as build

RUN mkdir -p /src/build
WORKDIR /src/build

COPY . .

RUN make build

# -------------- RUN ---------------- #
FROM alpine

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY --from=build /src/build/dist/sur ./
COPY tmpl/ tmpl/

EXPOSE 8080

ENTRYPOINT ["./sur"]
