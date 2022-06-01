export const API_BASE = process.env.WEBSOCKET_MANAGER  as string;
export const CENTRIFUGO_URL = process.env.CENTRIFUGO_URL as string;
export const CENTRIFUGO_HMAC_SECRET = process.env.CENTRIFUGO_HMAC_SECRET as string;
export const CENTRIFUGO_WAIT_EVENTS_TIMEOUT = process.env.CENTRIFUGO_WAIT_EVENTS_TIMEOUT ? +process.env.CENTRIFUGO_WAIT_EVENTS_TIMEOUT : 0;

export const DB_HOST = process.env.DB_HOST as string;
export const DB_PORT = process.env.DB_PORT as string;
export const DB_USER = process.env.DB_USER as string;
export const DB_PASSWORD = process.env.DB_PASSWORD as string;
export const DB_NAME = process.env.DB_NAME as string;

export const API = {
    USERS: {
        PUBLISH: `${API_BASE}/users/publish`,
    },
    GROUP_SUBSCRIPTIONS: {
        PUBLISH: `${API_BASE}/group-subscriptions/publish`,
        CREATE: `${API_BASE}/group-subscriptions`,
        BATCH_REMOVE: `${API_BASE}/group-subscriptions/batch-remove`,
        CLEAR_BY_GROUP_ID: (groupId: number) => `${API_BASE}/group-subscriptions/groups/${groupId}`,
        CLEAR_BY_USER_ID: (userId: number) => `${API_BASE}/users/${userId}/group-subscriptions`,
    },
    CHANNEL_SUBSCRIPTIONS: {
        PUBLISH: `${API_BASE}/channel-subscriptions/publish`,
        CREATE: `${API_BASE}/channel-subscriptions`,
        BATCH_REMOVE: `${API_BASE}/channel-subscriptions/batch-remove`,
        CLEAR_BY_CHANNEL_ID: (channelId: number) => `${API_BASE}/channel-subscriptions/channels/${channelId}`,
        CLEAR_BY_USER_ID: (userId: number) => `${API_BASE}/users/${userId}/channel-subscriptions`,
    },
    DIALOG_SUBSCRIPTIONS: {
        PUBLISH: `${API_BASE}/dialog-subscriptions/publish`,
        CREATE: `${API_BASE}/dialog-subscriptions`,
        BATCH_REMOVE: `${API_BASE}/dialog-subscriptions/batch-remove`,
        CLEAR_BY_INITIATOR_ID: (userId: number) => `${API_BASE}/users/${userId}/dialog-subscriptions`,
    }
}
