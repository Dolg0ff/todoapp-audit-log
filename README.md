# Сервис логгирования для приложения создания TODO Списков (ветка gRPC)

## Использовал следующие концепции:
- gRPC и protocol buffers
- Работа с фреймворком gin-gonic/gin.
- Работа с NoSQL(MongoDB). Запуск из Docker.
- Graceful Shutdown
- Логгирование Logrus

# Запуск

## Для запуска необходимо указать переменные окружения, например в файле .env

```Bash
export DB_URI=localhost:27017
export DB_USERNAME=admin
export DB_PASSWORD=g0langn1nja
export DB_DATABASE=audit
export DB_COLLECTION=logs

export SERVER_PORT=9000
```

## Сброрка и запуск

```Bash
source .env && go build -o app cmd/main.go && ./app
```

## Для mongo можно использовать Docker

```Bash
docker run --rm -d --name audit-log-mongo \
> -e MONGO_INITDB_ROOT_USERNAME=admin \
> -e MONGO_INITDB_ROOT_PASSWORD=g0langn1nja \
> -p 27017:27017 mongo:latest
```