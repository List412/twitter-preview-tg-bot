[![Go](https://github.com/List412/twitter-preview-tg-bot/actions/workflows/go.yml/badge.svg)](https://github.com/List412/twitter-preview-tg-bot/actions/workflows/go.yml)

[Бот](https://t.me/twitter_preview_tg_bot) умеет брать твиты по ссылке и скидывать их в телеграм)) 

# build

создать .env на основе default.env

для линуха:
```
GOOS=linux go build -o bot cmd/bot/main.go
```

# как использовать

* Добавьте бота в конфу
* Скидывайте ссылку на твит, например через "поделиться" на мобилках или просто ссылку ввида https://twitter.com/silooo89/status/1581660694612697088?s=46&t=B_1vnsxYE3TVczwwnubdoA
* Бот возвращает текст картинки и видео с основного твита

Так же бот умеет подтягивать продолжение твита из реплаев автора