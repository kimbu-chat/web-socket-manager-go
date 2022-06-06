
-- +migrate Up
CREATE TABLE group_subscriptions
(
    user_id BIGINT,
    group_id BIGINT,
    created_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY(user_id, group_id)
);

CREATE TABLE channel_subscriptions
(
    user_id BIGINT,
    channel_id BIGINT,
    created_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY(user_id, channel_id)
);

CREATE TABLE dialog_subscriptions
(
    initiator_id BIGINT,
    user_id BIGINT,
    created_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY(initiator_id, user_id)
);

-- +migrate Down
DROP TABLE group_subscriptions;
DROP TABLE dialog_subscriptions;
DROP TABLE channel_subscriptions;
