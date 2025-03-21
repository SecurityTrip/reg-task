# Этап сборки
FROM golang:1.24.1-alpine AS builder

# Устанавливаем переменные окружения для статической сборки
ENV CGO_ENABLED=0

WORKDIR /app

# Копируем файлы зависимостей и загружаем их
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарник
RUN go build -o backend .

# Финальный этап
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Копируем бинарник из этапа сборки
COPY --from=builder /app/backend .

# Открываем порт, который слушает приложение
EXPOSE 8080

# Запускаем приложение
CMD ["./backend"]
