CREATE TABLE users
(
    id         BIGSERIAL PRIMARY KEY,
    username   TEXT                     NOT NULL,
    type       INTEGER                  NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX ON users (username);

CREATE TABLE user_passwords
(
    id              BIGSERIAL PRIMARY KEY,
    user_id         BIGINT                   NOT NULL REFERENCES users (id),
    hashed_password TEXT                     NOT NULL,
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    expired_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX ON user_passwords (user_id);

CREATE TABLE user_login_attempts
(
    id         BIGSERIAL PRIMARY KEY,
    user_id    BIGINT                   NOT NULL REFERENCES users (id),
    succeeded  BOOLEAN                  NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX ON user_login_attempts (user_id);