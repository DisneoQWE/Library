FROM golang:1.18-alpine AS builder

#Задает рабочую директорию внутри самого образа
WORKDIR /FirstTask
#Копируем в контейнер файлы
COPY go.mod .
COPY go.sum .
#переменные среды proxy
ENV HTTP_PROXY="http://proxy.halykbank.nb:8080"
ENV HTTPS_PROXY="http://proxy.halykbank.nb:8080"
ENV NO_PROXY="localhost, 127.0.0.0/8, ::1, 172.*, *.halykbank.nb, *.consul, *kkb.kz, websd, servicedesk, myhalyk, ala620b03i06, om, gitlab.cloud.halykbank.nb, jira.halykbank.kz, confluence.halykbank.kz, *.homebank.kz, 172.26.60.81, *.cloud.halykbank.nb, *.service.test-dc.consul"
#выполняем команду, скачиваем зависимости из mod
RUN go mod download
#копируем в контейнеры осталньые файлы
COPY . /FirstTask
#выполняе команду билда
RUN go build -o /rest-api_library

FROM alpine:3

#киоруем в контейнер билд builder
COPY --from=builder /rest-api_library .
#открываем на необходимость порта
EXPOSE "1234"
#описываем команду с аргументами, которую нужно выполнять когда конт. запущен
CMD ["/rest-api_library"]