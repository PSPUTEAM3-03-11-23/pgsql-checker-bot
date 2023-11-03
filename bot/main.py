import datetime
import logging
import os
import re
from datetime import time
from functools import wraps
from time import sleep

import pytz
from telegram.ext import ApplicationBuilder, CommandHandler, CallbackContext, ContextTypes, ConversationHandler, \
    MessageHandler

from telegram import ForceReply, Update
from telegram.ext.filters import TEXT, COMMAND

FIO, SECRET_CODE, INVITE_CODE = range(3)

# from telegram import Update, InlineKeyboardButton, InlineKeyboardMarkup, Bot, \
#     InputMediaPhoto
# from telegram.ext.filters import TEXT, COMMAND, PHOTO


# from sql_alchemy.models import Account, Message, Session, Photo, MessageRecipient, Post, PostRecipient, Member

async def start(update: Update, context: ContextTypes.DEFAULT_TYPE):
    await context.bot.send_message(chat_id=update.effective_chat.id,
                                   text=update.message.text)
    return FIO

async def fioInput(update: Update, context: ContextTypes.DEFAULT_TYPE):
    await context.bot.send_message(chat_id=update.effective_chat.id,
                                   text="Введите Secret")
    return SECRET_CODE
async def secretInput(update: Update, context: ContextTypes.DEFAULT_TYPE):
    await context.bot.send_message(chat_id=update.effective_chat.id,
                                   text="Введите Invite")
    return INVITE_CODE
async def inviteInput(update: Update, context: ContextTypes.DEFAULT_TYPE):
    await context.bot.send_message(chat_id=update.effective_chat.id,
                                   text="Super!")

    return ConversationHandler.END


if __name__ == '__main__':
    application = ApplicationBuilder().token(
        os.environ.get('TG_TOKEN', '1239481186:AAGj2GoeUJHGVXYaYcXSUz4igo-4pT8As3M')).build()

    # start_handler = CommandHandler('start', start)  # /start


    start_handler = ConversationHandler(
        entry_points=[CommandHandler('start', start)], # /new_account
        states={
            FIO: [MessageHandler(TEXT & (~COMMAND), fioInput)],
            INVITE_CODE: [MessageHandler(TEXT & (~COMMAND), inviteInput)],
            SECRET_CODE: [MessageHandler(TEXT & (~COMMAND), secretInput)],
        },
        fallbacks=[],
        #TODO!: conversation_timeout=CONVERSATION_TIMEOUT
    )

    application.add_handler(start_handler)
    print('Started')
    application.run_polling()