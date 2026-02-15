# Get My ID Max Bot

Простой бот для Max Messenger, который отправляет пользователю его User ID.
*Оригинальный бот:* https://max.ru/id631822786049_bot

## Описание

Бот автоматически отвечает на любые сообщения, отправляя пользователю его уникальный идентификатор в Max.

## Требования

- Go 1.25.6 или выше
- Токен бота Max

## Установка

1. Клонируйте репозиторий:
```bash
git clone https://github.com/kurochki-n/get-my-id-max-bot.git
cd get-my-id-max-bot
```

2. Установите зависимости:
```bash
go mod download
```

3. Создайте файл `.env` на основе `.env.example`:
```bash
cp .env.example .env
```

4. Укажите токен вашего бота в файле `.env`:
```
BOT_TOKEN=ваш_токен_бота
```

## Запуск

```bash
go run main.go
```

## Использование

1. Запустите бота
2. Откройте чат с ботом в Max
3. Отправьте любое сообщение
4. Бот ответит вам с вашим User ID

## Зависимости

- [max-bot-api-client-go](https://github.com/max-messenger/max-bot-api-client-go) - клиент для работы с Max Bot API
- [godotenv](https://github.com/joho/godotenv) - загрузка переменных окружения из .env файла

## Лицензия

MIT
