
-- +migrate Up
CREATE TABLE user_group_subscriptions
(
    user_id BIGINT,
    group_id BIGINT,
    PRIMARY KEY(user_id, group_id)
);

CREATE TABLE user_interlocutor_subscriptions
(
    user_id BIGINT,
    interlocutor_id BIGINT,
    PRIMARY KEY(user_id, interlocutor_id)
);

-- +migrate Down
DROP TABLE user_group_subscriptions;
DROP TABLE user_interlocutor_subscriptions;
