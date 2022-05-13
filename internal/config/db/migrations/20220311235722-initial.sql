
-- +migrate Up
CREATE TABLE group_subscriptions
(
    user_id BIGINT,
    group_id BIGINT,
    PRIMARY KEY(user_id, group_id)
);

CREATE TABLE channel_subscriptions
(
    user_id BIGINT,
    channel_id BIGINT,
    PRIMARY KEY(user_id, channel_id)
);

CREATE TABLE dialog_subscriptions
(
    initiatorId BIGINT,
    userId BIGINT,
    PRIMARY KEY(initiatorId, userId)
);

-- +migrate Down
DROP TABLE group_subscriptions;
DROP TABLE dialog_subscriptions;
DROP TABLE channel_subscriptions;
