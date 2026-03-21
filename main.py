import asyncio
import logging

from maxapi import Bot, Dispatcher
from maxapi.types import BotStarted, Command, MessageCreated
from maxapi.enums.parse_mode import ParseMode

from config_reader import config

logging.basicConfig(level=logging.INFO)

bot = Bot(
    token=config.BOT_TOKEN.get_secret_value(),
    parse_mode=ParseMode.HTML
)
dp = Dispatcher()


@dp.bot_started()
async def bot_started(event: BotStarted):
    """Ответ бота при нажатии на кнопку `Начать`."""
    await event.bot.send_message(
        chat_id=event.chat_id,
        text=(
            f"user_id: <code>{event.user.user_id}</code>\n"
            f"chat_id: <code>{event.chat.chat_id}</code>"
        )
    )


@dp.message_created()
async def message_handler(event: MessageCreated):
    """Отвтет бота на любое сообщение."""
    await event.message.answer(
        text=(
            f"user_id: <code>{event.from_user.user_id}</code>\n"
            f"chat_id: <code>{event.chat.chat_id}</code>"
        )
    )


async def main():
    await bot.delete_webhook()
    await dp.start_polling(bot, skip_updates=True)


if __name__ == '__main__':
    asyncio.run(main())