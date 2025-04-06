# Banking App API

REST API сервис для управления банковскими операциями, написанный на Go с использованием MongoDB в качестве основной базы данных и Redis для кэширования.

## Технологии

- Go
- Echo (веб-фреймворк)
- MongoDB
- Redis
- Docker & Docker Compose
- JWT (аутентификация)
- Swagger (документация API)
- Viper (конфигурация)
- Logrus (логирование)

## Предварительные требования

- Docker
- Docker Compose
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
git clone https://github.com/MDmitryM/banking-app-go.git
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
docker-compose up --d
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

## Структура проекта

- `/cmd` - точка входа в приложение
- `/pkg` - основной код приложения
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
