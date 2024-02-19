# Expressionist

Проект для Лицея Академии Яндекс

### !!! Я еще дополняю документацию

### В случае ошибки или при любых других вопросах пишите сюды [@moolcoov](https://moolcoov.t.me/)

## Установка

#### Docker

Для начала на вашем компьютере должен быть установлен [Docker](https://docker.com). Как его установить написано в [документации](https://docs.docker.com/get-docker/).

#### Клонирование репозитория

Далее необходимо склонировать репозиторий с кодом

```bash
git clone https://github.com/moolcoov/expressionist.git
```

#### Build

После этого нужно создать билд с помощью Docker Compose

```bash
docker-compose build
```

P.S Этот процесс может занять довольно много времени ☕

## Запуск

После этого можно запускать проект

Для того чтобы запустить все сервисы:

```bash
docker-compose up
```

Для того чтобы запустить еще агентов:

```bash
docker-compose scale agent=3
```

Вместо 3 можно подставить любое число, сколько агентов вы хотите.

## На чем оно написано

### 1. Оркестратор

Оркестратор находится в директории [`/orchestra`](https://github.com/moolcoov/expressionist/tree/main/orchestra). Написан на Go.

#### Роутинг

Для роутинга изпользуется [gorilla/mux](https://github.com/gorilla/mux). Все эндпоинты в субдиректории [`/routes`](https://github.com/moolcoov/expressionist/tree/main/orchestra/routes).

#### База данных

База данных, в которую сохраняются выражения: [PostgreSQL](https://www.postgresql.org/). Для доступа используется [sqlx](https://github.com/jmoiron/sqlx) с драйвером [pgx](https://github.com/jackc/pgx). Файл, в котором происходит подключение к бд [`/lib/postgres.go`](https://github.com/moolcoov/expressionist/blob/main/orchestra/lib/postgres.go).

#### Кэш

Для кэширования результатов используется key-value база данных [Redis](https://redis.io/). Для подключения применяется пакет [`go-redis`](github.com/redis/go-redis). Файл, в котором происходит подключение [`/lib/redis.go`](https://github.com/moolcoov/expressionist/blob/main/orchestra/lib/redis.go)

#### Связь с агентами

Для передачи выражений агентам используется [RabbitMQ](https://rabbitmq.com/). Для подключения применяется [`amqp091-go`](https://github.com/rabbitmq/amqp091-go). Файл [`/lib/redis.go`](https://github.com/moolcoov/expressionist/blob/main/orchestra/lib/rabbitmq.go).

#### Проверка активности агентов

Активность агентов проверяется в отдельной горутине с бесконечным циклом, которая каждые 10 секунд проверяет, не истекло ли время пинга у агентов.

```go
go func() {
    for {
        agent.Agents.CheckAgents()
        time.Sleep(10 * time.Second)
    }
}()
```

Схема оркестратора в конце концов выглядит так:

![s](./.github/media/orchestra.png)

### 2. Агенты

Агент находится в директории [`/agent`](https://github.com/moolcoov/expressionist/tree/main/agent). Написан на Go.

#### Получение выражений

Агент получает выражения с помощью того же [RabbitMQ](https://rabbitmq.com/). Схема подключения такая же, как и у оркестратора. [`/lib/redis.go`](https://github.com/moolcoov/expressionist/blob/main/agent/lib/rabbitmq.go).

#### Вычисление

Для вычисления используется измененная версия библиотеки [go-shunting-yard](https://github.com/mgenware/go-shunting-yard).

Схема:

![](./.github/media/agent.png)

### 3. Клиент

Фронтенд написан на [Typescript](https://www.typescriptlang.org/), с использованием фреймворка [Next.js](https://nextjs.org). Находится в директории [`/client`](https://github.com/moolcoov/expressionist/tree/main/client).

#### Запросы к оркестратору

На серверной части запросы делаются с помощью официального API (добавление нового выражения, обновление настроек).

```ts
const res = await fetch(... ,{})
```

На клиентской части используется библиотека [SWR](https://swr.vercel.app/) (получение списка выражений, агентов, настроек).

```ts
const { data, error, isLoading } = useSWR(..., fetcher)
```

Схема:

![](./.github/media/client.png)
