<<<<<<< HEAD
# Banking App API

REST API сервис для управления банковскими операциями, написанный на Go с использованием MongoDB в качестве основной базы данных и Redis для кэширования.
=======
# Todo App GO

Простое backend приложение для управления задачами (todo list), написанное на Go с использованием PostgreSQL в качестве базы данных.
>>>>>>> 7f4c1f4c839b7830929c9bbc74bba3f63c6b913f

## Технологии

- Go
<<<<<<< HEAD
- Echo (веб-фреймворк)
- MongoDB
- Redis
- Docker & Docker Compose
- JWT (аутентификация)
- Swagger (документация API)
- Viper (конфигурация)
- Logrus (логирование)
=======
- PostgreSQL
- Docker & Docker Compose
- Migrate (для управления миграциями базы данных)
>>>>>>> 7f4c1f4c839b7830929c9bbc74bba3f63c6b913f

## Предварительные требования

- Docker
- Docker Compose
<<<<<<< HEAD
- Go 1.21 или выше

## API Documentation

Полная API документация доступна через Swagger UI по адресу: `http://localhost:8080/swagger/`

### Доступные эндпоинты
#### Аутентификация

| Метод | Эндпоинт      | Описание                    | Требует авторизации |
|-------|---------------|----------------------------|-------------------|
| POST  | /auth/sign-up | Регистрация нового пользователя | Нет |
| POST  | /auth/sign-in | Вход в систему (получение JWT токена) | Нет |

#### Банковские операции

| Метод | Эндпоинт      | Описание                    | Требует авторизации |
|-------|---------------|----------------------------|-------------------|
| GET   | /api/accounts    | Получить все счета пользователя | Да |
| POST  | /api/accounts    | Создать новый счет        | Да |
| GET   | /api/accounts/:id | Получить информацию о счете | Да |
| POST  | /api/transactions | Создать новую транзакцию   | Да |
| GET   | /api/transactions | Получить историю транзакций | Да |

## Установка и запуск

### Использование Docker Compose

1. Клонируйте репозиторий:
```bash
git clone https://github.com/your-username/banking-app-go.git
cd banking-app-go
```

2. Создайте файл `.env` в корневой директории проекта:
```env
# MongoDB
MONGO_USER=mongo
MONGO_PASSWORD=your_password
MONGO_DB_NAME=budget_app

# Redis
REDIS_PASSWORD=your_redis_password

# Application
JWT_SECRET=your_jwt_secret
ENV=production
```

3. Запустите приложение:
```bash
docker-compose up --build
```

Приложение будет доступно по адресу `http://localhost:8080`

### Переменные окружения

| Переменная    | Описание                           |
|---------------|------------------------------------|
| MONGO_USER    | Имя пользователя MongoDB           |
| MONGO_PASSWORD| Пароль для MongoDB                 |
| MONGO_DB_NAME | Название базы данных               |
| REDIS_PASSWORD| Пароль для Redis                   |
| JWT_SECRET    | Ключ для подписи JWT токенов       |
| ENV           | Окружение (development/production) |

### Порты

- 8080: API сервер
- 27017: MongoDB (доступен на хосте)
- 6379: Redis (доступен на хосте)
=======

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
>>>>>>> 7f4c1f4c839b7830929c9bbc74bba3f63c6b913f

## Структура проекта

- `/cmd` - точка входа в приложение
- `/pkg` - основной код приложения
<<<<<<< HEAD
  - `/handler` - обработчики HTTP запросов
  - `/repository` - работа с базами данных
  - `/service` - бизнес-логика
- `/configs` - конфигурационные файлы
- `/docs` - документация API (Swagger)

## База данных

MongoDB база данных автоматически инициализируется при первом запуске с указанными учетными данными.

### Конфигурация

Основные настройки приложения находятся в:
- `.env` - переменные окружения (учетные данные БД, ключи)
- `configs/config.yml` - конфигурация приложения (порты, хосты)
=======
- `/configs` - конфигурационные файлы
- `/schema` - миграции базы данных
- `/docs` - документация

## База данных

PostgreSQL база данных автоматически инициализируется при первом запуске. Миграции применяются автоматически с помощью сервиса `migrate` в Docker Compose.
>>>>>>> 7f4c1f4c839b7830929c9bbc74bba3f63c6b913f
