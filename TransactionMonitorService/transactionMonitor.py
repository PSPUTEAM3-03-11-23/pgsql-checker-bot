import json
import win32pipe, win32file

# Здесь подготавливаем данные для отправки
database_info = {
    "DB_HOST": "192.168.0.18",
    "DB_USER": "postgres",
    "DB_PASS": "pspu!Team",
    "DB_NAME": "hackathon",
    "DB_PORT": 5432
}
json_data = json.dumps(database_info)

# Именованный канал должен совпадать с тем, который указан в C# приложении
pipe_name = r'\\.\pipe\dbConnectToWriteProcess'

# Подключаемся к именованному каналу
pipe_handle = win32file.CreateFile(
    pipe_name,
    win32file.GENERIC_READ | win32file.GENERIC_WRITE,
    0,
    None,
    win32file.OPEN_EXISTING,
    0,
    None
)

# Отправляем данные
win32file.WriteFile(pipe_handle, json_data.encode())

# Закрываем соединение
win32file.CloseHandle(pipe_handle)
