# Используем официальный образ Golang в качестве базового образа
FROM golang:latest

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /src/server

# Копируем исходный код в рабочую директорию контейнера
COPY . .

# Собираем приложение
RUN go build -o server cmd/server/app.go

# Экспортируем порт 8080 для взаимодействия с внешним миром
EXPOSE 8080

# Команда для запуска приложения при запуске контейнера
CMD ["./server"]
