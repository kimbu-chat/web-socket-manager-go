
-- +migrate Up
CREATE TABLE user_group_subscriptions
(
    user_id BIGINT,
    group_id BIGINT,
    PRIMARY KEY(user_id, group_id)
);

CREATE TABLE dialog_subscriptions
(
    initiatorId BIGINT,
    userId BIGINT,
    PRIMARY KEY(initiatorId, userId)
);

-- +migrate Down
DROP TABLE user_group_subscriptions;
DROP TABLE dialog_subscriptions;
