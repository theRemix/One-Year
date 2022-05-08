FROM golang:1.15-alpine as server

RUN apk --update add ca-certificates
RUN apk --update add gcc musl-dev

COPY . /go/src/app/
WORKDIR /go/src/app/

ENV APP_BUILD_NAME="one-year"
ENV CGO_ENABLED=1
RUN go mod download
RUN go build -ldflags="-linkmode external -extldflags '-static' -s -w" -o ${APP_BUILD_NAME} .
RUN chmod +x ${APP_BUILD_NAME}


FROM node:lts-alpine as client

COPY client/ /srv/
WORKDIR /srv/

RUN apk --update add git
RUN npx yarn install
RUN npm run build


FROM scratch

WORKDIR /
COPY --from=server /go/src/app/one-year /one-year
COPY --from=server /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=client  /srv/public/ /client/public/

EXPOSE 4000
CMD [ "/one-year" ]
