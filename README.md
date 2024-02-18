
# Expressionist

Проект для Лицея Академии Яндекс



### !!! Я еще дополняю документацию

### В случае ошибки или при любых других вопросах пишите на [@moolcoov](https://moolcoov.t.me/)


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