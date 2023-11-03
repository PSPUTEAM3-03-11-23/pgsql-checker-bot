import datetime
import logging
import os
import re
from datetime import time
from functools import wraps
from time import sleep

import pytz
from telegram.ext import ApplicationBuilder, CommandHandler, CallbackContext, ContextTypes

from telegram import ForceReply, Update

# from telegram import Update, InlineKeyboardButton, InlineKeyboardMarkup, Bot, \
#     InputMediaPhoto
# from telegram.ext.filters import TEXT, COMMAND, PHOTO


# from sql_alchemy.models import Account, Message, Session, Photo, MessageRecipient, Post, PostRecipient, Member

async def start(update: Update, context: ContextTypes.DEFAULT_TYPE):
    await context.bot.send_message(chat_id=update.effective_chat.id,
                                   text="Hello")


if __name__ == '__main__':
    application = ApplicationBuilder().token(
        os.environ.get('TG_TOKEN', '1239481186:AAGj2GoeUJHGVXYaYcXSUz4igo-4pT8As3M')).build()

    start_handler = CommandHandler('start', start)  # /start
    application.add_handler(start_handler)
    application.run_polling()
