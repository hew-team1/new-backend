CREATE TABLE users
(
    id         int(100) AUTO_INCREMENT NOT NULL,
    uid        varchar(255)            NOT NULL,
    name       varchar(255)            NOT NULL,
    email      varchar(255)            NOT NULL,
    created_at timestamp               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE recruits
(
    id               int(100) AUTO_INCREMENT NOT NULL,
    user_id          int(100)                NOT NULL,
    title            varchar(255)            NOT NULL,
    organizer        varchar(255)            NOT NULL,
    commit_frequency varchar(5)              NOT NULL,
    message          varchar(255)            NOT NULL,
    join_number      int(10)                 NOT NULL,
    is_beginner      int(1)                  NOT NULL,
    is_award         int(1)                  NOT NULL,
    slack_url        varchar(255)            NOT NULL,
    start_date       date                    NOT NULL,
    end_date         date                    NOT NULL,
    created_at       timestamp               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       timestamp               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE members
(
    user_id    int(100)     NOT NULL,
    recruit_id int(100)     NOT NULL,
    position   varchar(255) NOT NULL,
    created_at timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (recruit_id) REFERENCES recruits (id)
);