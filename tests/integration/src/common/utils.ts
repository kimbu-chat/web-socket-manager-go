import _ from "lodash";
import axios from "axios";
import {API} from "./environment";

class NoErrorThrownError extends Error {}
export const getError = async <TError>(call: () => unknown): Promise<TError> => {
    try {
        await call();

        throw new NoErrorThrownError();
    } catch (error: unknown) {
        return error as TError;
    }
};

export const getRandomId = () => _.random(1, 100_000);

export const createDialogSubscriptions = (userIds: number[], initiatorId: number) => {
    return axios.post(API.DIALOG_SUBSCRIPTIONS.CREATE, { userIds: userIds, initiatorId: initiatorId });
}

export const publishToInterlocutors = (initiatorId: number, message: any) => {
    return axios.post(API.DIALOG_SUBSCRIPTIONS.PUBLISH, { initiatorId: initiatorId, message: message });
}

export const createChannelSubscriptions = (userIds: number[], channelId: number) => {
    return axios.post(API.CHANNEL_SUBSCRIPTIONS.CREATE, { userIds: userIds, channelId: channelId });
}

export const publishToChannel = (channelId: number, message: any) => {
    return axios.post(API.CHANNEL_SUBSCRIPTIONS.PUBLISH, { channelId: channelId, message: message });
}

export const createGroupSubscriptions = (userIds: number[], groupId: number) => {
    return axios.post(API.GROUP_SUBSCRIPTIONS.CREATE, { userIds: userIds, groupId: groupId });
}

export const publishToGroup = (groupId: number, message: any) => {
    return axios.post(API.GROUP_SUBSCRIPTIONS.PUBLISH, { groupId: groupId, message: message });
}

export const publishToUsers = (userIds: number[], message: any) => {
    return axios.post(API.USERS.PUBLISH, { userIds: userIds, message: message});
}
