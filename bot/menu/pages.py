from telegram import InlineKeyboardButton

class Pages:
    admin_main_page = [
        [
            InlineKeyboardButton('Генерация кода приглашения', callback_data='admin_generate_invite_code_page')
        ],
        [
            InlineKeyboardButton('Базы данных', callback_data='generate_invite_code_page')
        ],
        [
            InlineKeyboardButton('Выдача доступа к БД', callback_data='bd_access_page')
        ],
        [
            InlineKeyboardButton('Настройки оповещений', callback_data='alerts_settings_page'),
            InlineKeyboardButton('Уведомления', callback_data='alerts_page')
        ],
    ]
    code_generator_menu = [
        [
            InlineKeyboardButton('Генерация кода приглашения', callback_data='admin_generate_invite_code')
        ],
        [
            InlineKeyboardButton('Назад', callback_data='go_admin_main_menu')
        ],
    ]