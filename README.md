# TODO App
## Запуск
```bash
# С переменными окружения
docker build -t todo-app .
docker run -p 7540:7540 -e TODO_PASSWORD=secret todo-app

# Локально
go run main.go