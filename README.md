# get-my-id-max-bot

Бот для мессенджера [Max](https://max.ru), который отправляет пользователю его `user_id` в ответ на любое сообщение и при нажатии кнопки «Начать».

## Стек

- Python 3.13+
- [maxapi](https://github.com/max-messenger/maxapi) — фреймворк для Max Bot API
- pydantic-settings — управление конфигурацией

## Установка

```bash
# Клонировать репозиторий
git clone https://github.com/kurochki-n/get-my-id-max-bot.git
cd get-my-id-max-bot

# Установить зависимости через uv
uv sync
```

## Конфигурация

Скопируй `.env.example` в `.env` и укажи токен бота:

```bash
cp .env.example .env
```

`.env`:
```
BOT_TOKEN=<Max Bot API токен>
```

## Запуск

```bash
uv run main.py
```

## Структура

```
.
├── main.py           # Точка входа, обработчики событий
├── config_reader.py  # Чтение конфигурации из .env
├── .env.example      # Пример файла окружения
└── pyproject.toml    # Зависимости проекта
```
