# ------------- BUILD --------------- #
FROM golang:1.15 as build

RUN mkdir -p /src/build
WORKDIR /src/build

COPY . .

RUN make build

# -------------- RUN ---------------- #
FROM alpine

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

RUN addgroup -g 21337 app
RUN adduser -D -H -u 21337 -G app app
USER app

COPY --from=build /src/build/dist/sur ./
COPY tmpl/ tmpl/

EXPOSE 8080

ENTRYPOINT ["./sur"]
