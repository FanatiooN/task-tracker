# task-tracker

### Быстрый старт

```bash
# 1. Скопировать примеры env-файлов и заполнить значения
cp .env.example .env
cp gateway/.env.example      gateway/.env
cp auth-service/.env.example auth-service/.env
cp user-service/.env.example user-service/.env
cp task-service/.env.example task-service/.env

# 2. Поднять всё
docker compose up --build
```


### Auth

| Метод | Путь                     | Что делает                    |
|-------|--------------------------|-------------------------------|
| POST  | `/register`              | регистрация email + пароль    |
| POST  | `/login`                 | вход по email + паролю        |
| POST  | `/refresh`               | обмен refresh на новые токены |
| POST  | `/logout`                | отзыв refresh-токена          |
| POST  | `/login/telegram`        | вход по Telegram id_token     |
| POST  | `/login/google`          | редирект на Google OAuth      |
| GET   | `/login/google/callback` | Google OAuth callback         |

### Tasks

| Метод  | Путь          | Что делает                                                                    |
|--------|---------------|-------------------------------------------------------------------------------|
| POST   | `/tasks`      | создать                                                                       |
| GET    | `/tasks`      | список заданий с фильтром + пагинацией (`?taskStatus=&pageSize=&pageToken=`)  |
| GET    | `/tasks/{id}` | получение задания                                                             |
| PUT    | `/tasks/{id}` | обновление задания                                                            |
| DELETE | `/tasks`      | удаление заданий с id `{ "ids": [...] }`                                      |

### Users

| Метод  | Путь          | Что делает |
|--------|---------------|------------|
| GET    | `/users/{id}` | прочитать  |
| PUT    | `/users/{id}` | обновить   |
| DELETE | `/users/{id}` | удалить    |


### Миграции внутри каджого сервиса

```bash
make migrate-up         # накатить
make migrate-down       # откатить одну
make migrate-status     # статус
make migrate-create name=add_something
```
