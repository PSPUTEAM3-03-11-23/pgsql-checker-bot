import os

import bcrypt
from sqlalchemy import Column, Integer, String, ForeignKey, create_engine, Boolean, DateTime, LargeBinary, JSON
from sqlalchemy.orm import declarative_base, relationship
from sqlalchemy.orm import sessionmaker

engine = create_engine(
    f"postgresql+psycopg2://postgres:pspu!Team@{os.environ.get('DB_HOST', '90.150.183.80:10100')}/hackathon")

Base = declarative_base()
Session = sessionmaker(bind=engine)


class User(Base):
    __tablename__ = 'users'

    id = Column(Integer, primary_key=True)
    tg_id = Column(Integer)
    name = Column(String(255))
    is_deactivated = Column(Boolean)

    user_secrets = relationship("UserSecret", back_populates='user', cascade="all, delete, delete-orphan")
    organization_admins = relationship("OrganizationAdmin", back_populates='user', cascade="all, delete, delete-orphan")
    user_db_subscribers = relationship("UserDBSubscriber", back_populates='user', cascade="all, delete, delete-orphan")
    alerts = relationship("Alert", back_populates='user', cascade="all, delete, delete-orphan")

    def __repr__(self):
        return f'{self.id} | {self.tg_id} | {self.name} | {self.is_deactivated}'


class Organization(Base):
    __tablename__ = 'organizations'

    id = Column(Integer, primary_key=True)
    title = Column(String(255))

    invite_codes = relationship("OrganizationInviteCode", back_populates='organization',
                                cascade="all, delete, delete-orphan")
    user_secrets = relationship("UserSecret", back_populates='organization', cascade="all, delete, delete-orphan")
    dbs = relationship("DB", back_populates='organization', cascade="all, delete, delete-orphan")
    organization_admins = relationship("OrganizationAdmin", back_populates='organization',
                                       cascade="all, delete, delete-orphan")

    def __repr__(self):
        return f'{self.id}|{self.title}'


class OrganizationInviteCode(Base):
    __tablename__ = 'organization_invite_codes'

    id = Column(Integer, primary_key=True)

    organization_id = Column(ForeignKey('organizations.id', ondelete='CASCADE'))
    invitation_code = Column(String(256), nullable=False)
    expiration_date = Column(DateTime, nullable=False)
    tg_username = Column(String(128), nullable=False)
    is_organization_admin = Column(Boolean, default=False)
    is_activated = Column(Boolean, default=False)

    organization = relationship("Organization", back_populates="invite_codes")

    def __repr__(self):
        return f'{self.id}|{self.organization.title}|{self.invitation_code}|{self.expiration_date}'


def get_hashed_password(plain_text_password, salt):
    # Hash a password for the first time
    #   (Using bcrypt, the salt is saved into the hash itself)
    return bcrypt.hashpw(plain_text_password, salt)


def check_password(plain_text_password, hashed_password):
    # Check hashed password. Using bcrypt, the salt is saved into the hash itself
    return bcrypt.checkpw(plain_text_password, hashed_password)


class UserSecret(Base):
    __tablename__ = 'user_secrets'
    id = Column(Integer, primary_key=True)
    organization_id = Column(ForeignKey('organizations.id', ondelete='CASCADE'))
    user_id = Column(ForeignKey('users.id', ondelete='CASCADE'))
    secret_code = Column(LargeBinary)
    salt = Column(LargeBinary)
    is_organization_admin = Column(Boolean, default=False)

    organization = relationship("Organization", back_populates="user_secrets")
    user = relationship("User", back_populates="user_secrets")

    def set_secret_code(self, new_pass):
        """Salt/Hash and save the user's new password."""
        self.salt = bcrypt.gensalt()
        new_password_hash = get_hashed_password(new_pass, self.salt)
        self.secret_code = new_password_hash

    def __repr__(self):
        return f'{self.id}|{self.recipient_id}|{self.recipient_first_name}|{self.recipient_last_name}'


class OrganizationAdmin(Base):
    __tablename__ = 'organization_admins'

    id = Column(Integer, primary_key=True)
    user_id = Column(ForeignKey('users.id', ondelete='CASCADE'))
    organization_id = Column(ForeignKey('organizations.id', ondelete='CASCADE'))

    organization = relationship("Organization", back_populates="organization_admins")
    user = relationship("User", back_populates="organization_admins")

    def __repr__(self):
        return f'{self.id}|{self.user.name}|{self.organization.title}'


class UserDBSubscriber(Base):
    __tablename__ = 'user_db_subscribers'

    id = Column(Integer, primary_key=True)
    user_id = Column(ForeignKey('users.id', ondelete='CASCADE'))
    db_id = Column(ForeignKey('dbs.id', ondelete='CASCADE'))

    db = relationship("DB", back_populates="user_db_subscribers")
    user = relationship("User", back_populates="user_db_subscribers")


class DB(Base):
    __tablename__ = 'dbs'

    id = Column(Integer, primary_key=True)
    host = Column(String(32))
    port = Column(Integer)
    username = Column(String(255))
    password = Column(String(255))
    db_name = Column(String(255))
    schema = Column(String(255))
    title = Column(String(255))
    organization_id = Column(ForeignKey('organizations.id', ondelete='CASCADE'))
    organization = relationship("Organization", back_populates="dbs")

    user_db_subscribers = relationship("UserDBSubscriber", back_populates='db', cascade="all, delete, delete-orphan")
    incidents = relationship("Incident", back_populates='db', cascade="all, delete, delete-orphan")

    def __repr__(self):
        return f'{self.id}|{self.title}'


class Incident(Base):
    __tablename__ = 'incidents'

    id = Column(Integer, primary_key=True)
    db_id = Column(ForeignKey('dbs.id', ondelete='CASCADE'))
    error = Column(JSON)
    date = Column(DateTime)

    alerts = relationship("Alert", back_populates='incident', cascade="all, delete, delete-orphan")
    db = relationship("DB", back_populates="incidents")

    def __repr__(self):
        return f'{self.id}|{self.error}'


class Alert(Base):
    __tablename__ = 'alerts'

    id = Column(Integer, primary_key=True)
    user_id = Column(ForeignKey('users.id', ondelete='CASCADE'))
    incident_id = Column(ForeignKey('incidents.id', ondelete='CASCADE'))
    is_sent = Column(Boolean, default=False)

    incident = relationship("Incident", back_populates="alerts")
    user = relationship("User", back_populates="alerts")

    def __repr__(self):
        return f'{self.id}|{self.user_id}| {self.is_sent}'


if __name__ == '__main__':
    Base.metadata.create_all(engine)
    with Session() as session:
        org = session.query(Organization).filter(Organization.title == 'main').first()
        if org is None:
            org = Organization(title='main')
            session.add(org)
            session.commit()
