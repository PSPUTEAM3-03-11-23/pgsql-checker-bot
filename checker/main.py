from sqlalchemy import text

from bot.sql_alchemy import Session


def show_backends(session):
    backends = (
        session.execute(text("SELECT pid, query FROM pg_stat_activity where state = 'active' OR state ='idle';"))
        .fetchall())
    print("found processes:")
    for backend in backends:
        print(f'{backend[0]}, {backend[1]}')


if __name__ == '__main__':
    with Session() as session:
        show_backends(session)
