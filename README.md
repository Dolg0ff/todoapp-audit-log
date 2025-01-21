# Сервис логгирования для приложения создания TODO Списков (ветка gRPC)

## Описание
Сервис аудита, реализующий логирование действий пользователей в TODO приложении с использованием MongoDB и gRPC.

## Технологический стек и концепции
- gRPC и Protocol Buffers для межсервисного взаимодействия
- MongoDB в качестве NoSQL БД для хранения логов
- gin-gonic/gin для HTTP endpoints
- Docker для развертывания MongoDB
- Graceful Shutdown
- Логирование с Logrus

## Структура проекта
```
├── cmd/              # Точка входа приложения
├── internal/         # Внутренний код приложения
│   ├── config/      # Конфигурация приложения
│   ├── repository/  # Слой работы с MongoDB
│   ├── server/      # gRPC сервер
│   └── service/     # Бизнес-логика
├── pkg/             # Переиспользуемые пакеты
│   └── domain/      # Модели и интерфейсы
└── proto/           # Protocol Buffers определения
```

## Требования
- Go (последняя стабильная версия)
- Docker
- Protocol Buffers compiler (protoc)
- Свободные порты:
  - 27017 (MongoDB)
  - 9000 (gRPC сервер)

## Быстрый старт

### 1. Запуск MongoDB в Docker

#### Windows (PowerShell):
```powershell
docker run --rm -d --name audit-log-mongo -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=g0langn1nja -p 27017:27017 mongo:latest
```

### 2. Установка переменных окружения

#### Windows (PowerShell):
```powershell
$env:DB_URI="localhost:27017"
$env:DB_USERNAME="admin"
$env:DB_PASSWORD="g0langn1nja"
$env:DB_DATABASE="audit"
$env:DB_COLLECTION="logs"
$env:SERVER_PORT="9000"
```

Или создайте файл `.env` и загрузите переменные:
```bash
source .env
```

### 3. Сборка и запуск приложения

#### С предварительной сборкой:
```bash
go build -o app cmd/main.go && ./app
```

#### Прямой запуск:
```bash
go run cmd/main.go
```

## Проверка работоспособности

### Проверка статуса MongoDB:
```bash
docker ps
```

### Подключение к MongoDB:
```bash
docker exec -it audit-log-mongo mongosh -u admin -p g0langn1nja
```

## Решение проблем

### Типичные проблемы

1. Ошибка подключения к MongoDB:
- Проверьте, что контейнер MongoDB запущен
- Убедитесь в правильности учетных данных
- Проверьте доступность порта 27017

2. Ошибки при запуске сервиса:
- Проверьте, что все переменные окружения установлены
- Убедитесь, что порт 9000 свободен
- Проверьте логи на наличие ошибок