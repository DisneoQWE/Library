FROM golang:1.19-alpine AS builder

#Задает рабочую директорию внутри самого образа
WORKDIR /source
#Копируем в контейнер файлы
#COPY go.mod /restApiLibrary
#COPY go.sum /restApiLibrary
COPY . /source
#переменные среды proxy
ENV HTTP_PROXY="http://proxy.halykbank.nb:8080"
ENV HTTPS_PROXY="http://proxy.halykbank.nb:8080"
ENV NO_PROXY="localhost, 127.0.0.0/8, ::1, 172.*, *.halykbank.nb, *.consul, *kkb.kz, websd, servicedesk, myhalyk, ala620b03i06, om, gitlab.cloud.halykbank.nb, jira.halykbank.kz, confluence.halykbank.kz, *.homebank.kz, 172.26.60.81, *.cloud.halykbank.nb, *.service.test-dc.consul"
#выполняем команду, скачиваем зависимости из mod

RUN go mod download
#копируем в контейнеры осталньые файлы

#выполняе команду билда
RUN CGO_ENABLED=0 GOOS=linux go build -o /app ./cmd/.

FROM alpine:3.9

#киоруем в контейнер билд builder
COPY --from=builder /app .
#открываем на необходимость порта
EXPOSE "1234"
#описываем команду с аргументами, которую нужно выполнять когда конт. запущен
ENTRYPOINT ["/app"]

#FROM golang:1.19-alpine AS builder
#WORKDIR /source
#COPY . /source
#RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -o app ./cmd/.
#
#FROM registry.query.consul:5000/alpine:3.9
#COPY --from=builder /source/app /usr/local/bin
#RUN chmod a+x /usr/local/bin/app
#ENTRYPOINT [ "app" ]