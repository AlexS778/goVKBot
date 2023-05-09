# Простой чат бот на базе ВКонтакте API

Бота можно попробовать тут - <https://vk.com/club220394114>, написав сообществу в личные сообщения "Начать"

## Функционал бота

У бота доступны 4 кнопки:
- Получить погоду: Бот отслыает сообщение с двумя inline кнопками, которыми можно выбирать погоду для города.
- Go to google.com: Нажатием на кнопку бот открывает cсылку <https://google.com>
- Получить фото кота!: Бот отсылает сообщение с фотографией кота.
- Забронировать столик: Бот отсылает сообщение с выбором времени для бронирования столика с 4 inline кнопками. После выбора времени высылается сообщение - подтверждение с 2 inline кнопками.

Картинки можно посмотреть снизу страницы

## Структра бота

```
|-goVkBot
| |-.gitignore
| |-go.mod
| |-README.md
| |-main.go
| |-go.sum
| |-folders.py
| |-internal
| | |-models
| | | |-models.go | все модели json необходимые для отправки и получения запросов
| | |-server
| | | |-server.go | модуль с функциями для "слушания" LongPollServer
| | |-bot
| | | |-bot.go | модуль с "обертками" для VKApi
| | |-utils
| | | |-utils.go | модуль с функциями "помощниками"
| |-.env
| |-Dockerfile
```

## Технические требования

Использован 1 сторонний пакет: <https://github.com/joho/godotenv> для загрузки переменных окружения из .env файла

## Работа с API ВКонтакте

Были "обернуты" следующие методы:

- groups.getLongPollServer
- messages.send
- messages.edit
- messages.sendMessageEventAnswer

## Запуск бота

1. Создайте файл .env cо следющими значениями:

```env
    TOKEN=`ваш access_token`
    GROUPID=`id вашего сообщества`
```

2. Создайте образ докера:

```shell
docker build -t go-vk-bot .
```

3. Запустите докер контейнер:

```
docker run -v /путь/к/вашему/.env:/app/.env go-vk-bot
```

## Картинки
<img src="https://github.com/AlexS778/goVKBot/blob/master/pics/bookatable.png" alt="book a table screenshot" style="height: 500px; width:667px;"/>
<img src="https://github.com/AlexS778/goVKBot/blob/master/pics/cat.png" alt="cat screenshot" style="height: 500px; width:667px;"/>
<img src="https://github.com/AlexS778/goVKBot/blob/master/pics/weather.png" alt="weather screenshot" style="height: 500px; width:667px;"/>
