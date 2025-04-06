# Todo App GO

Простое backend приложение для управления задачами (todo list), написанное на Go с использованием PostgreSQL в качестве базы данных.

## Технологии

- Go
- PostgreSQL
- Docker & Docker Compose
- Migrate (для управления миграциями базы данных)

## Предварительные требования

- Docker
- Docker Compose

## Установка и запуск

### С использованием Docker Compose

1. Клонируйте репозиторий:
```bash
git clone <repository-url>
cd todo-app-GO
```

2. Создайте файл `.env` в корневой директории проекта со следующими переменными:
```env
DB_PASSWORD=your_db_password
SIGNING_KEY=your_jwt_signing_key
SALT=your_password_salt
```

3. Запустите приложение с помощью Docker Compose:
```bash
docker-compose up -d
```

Приложение будет доступно по адресу `http://localhost:8000`
Swagger документация доступна по адресу `http://localhost:8000/swagger/`

### Переменные окружения

| Переменная    | Описание                                     |
|---------------|----------------------------------------------|
| DB_PASSWORD   | Пароль для базы данных PostgreSQL            |
| SIGNING_KEY   | Ключ для подписи JWT токенов                 |
| SALT          | Соль для хеширования паролей                 |

(Такие вещи как порт, имя пользователя, имя БД в postgres настраиваются в /configs/config.yml)
### Порты

- 8000: API сервер
- 5436: PostgreSQL

## Структура проекта

- `/cmd` - точка входа в приложение
- `/pkg` - основной код приложения
- `/configs` - конфигурационные файлы
- `/schema` - миграции базы данных
- `/docs` - документация

## База данных

PostgreSQL база данных автоматически инициализируется при первом запуске. Миграции применяются автоматически с помощью сервиса `migrate` в Docker Compose.
