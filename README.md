// TODO
# Задание 1

## Переменные окружения
- ```URL``` url, где нужно получать информацию о количестве NGL.
- ```PORT``` порт, по которому сервер будет слушать.

Запустить можно двумя способами:
- Локально
```
go run ./cmd/router -config .env
```
- Docker
```
docker build -t task1 .
sudo docker run \
    -p 8080:8080 \
    -e CONFIG_PATH=.env \
    -v ./.env:/router/.env:ro \
    --name task1 \
    task1
```