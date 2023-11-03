import datetime
import logging
import os
import re
from datetime import time, datetime
from functools import wraps
from time import sleep

import pytz
from sqlalchemy import select
from telegram.ext import ApplicationBuilder, CommandHandler, CallbackContext, ContextTypes, ConversationHandler, \
    MessageHandler

from telegram import ForceReply, Update
from telegram.ext.filters import TEXT, COMMAND

FIO, SECRET_CODE, INVITE_CODE = range(3)

# from telegram import Update, InlineKeyboardButton, InlineKeyboardMarkup, Bot, \
#     InputMediaPhoto
# from telegram.ext.filters import TEXT, COMMAND, PHOTO


from sql_alchemy.models import Session, User, OrganizationInviteCode, UserSecret


async def start(update: Update, context: ContextTypes.DEFAULT_TYPE):
    await context.bot.send_message(chat_id=update.effective_chat.id,
                                   text="Введите ФИО")
    return FIO


async def fioInput(update: Update, context: ContextTypes.DEFAULT_TYPE):
    await context.bot.send_message(chat_id=update.effective_chat.id,
                                   text="Введите Secret")
    context.user_data['user_fio'] = update.message.text
    return SECRET_CODE


async def secretInput(update: Update, context: ContextTypes.DEFAULT_TYPE):
    await context.bot.send_message(chat_id=update.effective_chat.id,
                                   text="Введите Invite")
    context.user_data['user_secret'] = update.message.text
    return INVITE_CODE


async def inviteInput(update: Update, context: ContextTypes.DEFAULT_TYPE):
    with Session() as session:
        invite = session.query(OrganizationInviteCode).filter(
            OrganizationInviteCode.tg_username == update.message.from_user.username
            and OrganizationInviteCode.invitation_code == update.message.text
            and not OrganizationInviteCode.is_activated).first()

        if invite.expiration_date < datetime.now():
            await context.bot.send_message(chat_id=update.effective_chat.id,
                                           text="Приглашение просрочено")
            return ConversationHandler.END

        user = User()
        user.name = context.user_data['user_fio']
        user_secret = UserSecret()
        user_secret.user = user
        user_secret.organization_id = invite.organization_id
        user_secret.set_secret_code(update.message.text.encode())
        session.add(user)
        session.add(user_secret)
        session.commit()

        await context.bot.send_message(chat_id=update.effective_chat.id,
                                       text=f"Регистрация успешна. Организация: {invite.organization.title}")

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