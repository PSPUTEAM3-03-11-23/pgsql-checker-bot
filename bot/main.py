import os
from telegram.ext import ApplicationBuilder, CommandHandler, ContextTypes, ConversationHandler, MessageHandler
from telegram import Update, InlineKeyboardButton, InlineKeyboardMarkup
from telegram.ext.filters import TEXT, COMMAND
import requests  # для отправки запросов на сервер

# Состояния разговора
START, BUTTON_PRESSED = range(2)

# Клавиатура для кнопки "Сделать запрос"
keyboard = [[InlineKeyboardButton("Сделать запрос", callback_data="make_request")]]
reply_markup = InlineKeyboardMarkup(keyboard)

# Основная функция обработки команды /start
def start(update: Update, context: ContextTypes.DEFAULT_TYPE):
    update.message.reply_text("Добро пожаловать! Нажмите на кнопку, чтобы сделать запрос на сервер.", reply_markup=reply_markup)
    return START

# Функция для обработки кнопки "Сделать запрос"
def button_pressed(update: Update, context: ContextTypes.DEFAULT_TYPE):
    query = update.callback_query
    query.answer()  # Ответить на запрос
    if query.data == "make_request":
        response = requests.get("https://example.com/api")  # Здесь вы должны указать URL вашего сервера
        query.edit_message_text(f"Ответ от сервера: {response.text}")

def main():
    app = ApplicationBuilder().token("YOUR_BOT_TOKEN")
    dp = app.dp

    # Добавление обработчика команды /start
    dp.add_handler(CommandHandler("start", start))

    # Добавление обработчика кнопки "Сделать запрос"
    dp.add_handler(ConversationHandler(
        entry_points=[],
        states={
            START: [MessageHandler(TEXT, lambda update, context: None)],
            BUTTON_PRESSED: [CallbackQueryHandler(button_pressed, pattern="make_request")],
        },
        fallbacks=[]
    ))

    app.run_polling()

if __name__ == '__main__':
    main()

import os


from telegram.ext import ApplicationBuilder, CommandHandler, ContextTypes, ConversationHandler, \
    MessageHandler
from telegram import Update, InlineKeyboardButton, InlineKeyboardMarkup
from telegram.ext.filters import TEXT, COMMAND
from bot.menu.pages import Pages

FIO, SECRET_CODE, INVITE_CODE, MAIN_PAGE = range(4)

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
    if update.message.text == 'Da':
        #TODO: Вынести в отдельный блок
        await context.bot.send_message(chat_id=update.effective_chat.id,
                                       text="Выберите пункт меню:",
                                       reply_markup=Pages.admin_main_page)



        return MAIN_PAGE
    else:
        await context.bot.send_message(chat_id=update.effective_chat.id, text="Введите Invite")
        return INVITE_CODE
    # with Session() as session:
    #     invite = session.query(OrganizationInviteCode).filter(
    #         OrganizationInviteCode.tg_username == update.message.from_user.username
    #         and OrganizationInviteCode.invitation_code == update.message.text
    #         and not OrganizationInviteCode.is_activated).first()
    #
    #     if invite.expiration_date < datetime.now():
    #         await context.bot.send_message(chat_id=update.effective_chat.id,
    #                                        text="Приглашение просрочено")
    #         return ConversationHandler.END
    #
    #     user = User()
    #     user.name = context.user_data['user_fio']
    #     user_secret = UserSecret()
    #     user_secret.user = user
    #     user_secret.organization_id = invite.organization_id
    #     user_secret.set_secret_code(update.message.text.encode())
    #     user_secret.is_organization_admin = invite.is_organization_admin
    #     invite.is_activated = True
    #     session.add(user)
    #     session.add(user_secret)
    #     session.commit()

        # await context.bot.send_message(chat_id=update.effective_chat.id,
        #                                text=f"Регистрация успешна. Организация: {invite.organization.title}")


async def mainInput(update: Update, context: ContextTypes.DEFAULT_TYPE):

    return SECRET_CODE


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
            MAIN_PAGE: [MessageHandler(TEXT & (~COMMAND), mainInput)],
        },
        fallbacks=[],
        #TODO!: conversation_timeout=CONVERSATION_TIMEOUT
    )
    application.add_handler(start_handler)
    print('Started')
    application.run_polling()