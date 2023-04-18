# Rusprofile scraper gRPC-Gateway 

gRPC & gRPC-gateway обертка над сервисом Rusprofile

# Технологии

- Go 1.19
- protobuf
- gRPC
- gRPC-Gateway
- Docker
    
# Getting Started
Для начала работы:

Клонировать git

Установить зависимости `go mod download`

Запусить с помощью `go run main.go`

Получить доступ к swagger-ui по http://localhost:8080/swaggerui/
# Использование

Пример использования с помощью cURL:

```bash
$ curl http://localhost:8080/v1/find/12345
{"code":5,"message":"No companies with provided INN","details":[]}

$ curl http://localhost:8080/v1/find/7751012274
{"INN":"7751012274","KPP":"775101001","NAME":"ООО \"Магнолия\"","FIO":"Амирджанов Шамай Рафаилович"}
```

# Deployment

Для развертки контейнера:
```bash
$ make imagebuild && make dockerrun
```
**OR**
```bash
$ docker build -t imagename:latest .

$ docker run -it -p 8080:8080 -p 8888:8888 --rm imagename:latest
```
