import axios, {AxiosError} from "axios";
import {API_BASE, CENTRIFUGO_URL, CENTRIFUGO_HMAC_SECRET, API} from "../common/environment";
import {publishAndTrackEvents} from "../common/websockets";
import {
    createChannelSubscriptions,
    createGroupSubscriptions,
    getError,
    getRandomId, publishToChannel,
    publishToGroup
} from "../common/utils";
import _ from "lodash";

describe("channel controller", () => {
    test("events aren't received after clearing subscriptions by user id", async () => {
        const userId = getRandomId();
        const channelId = getRandomId();

        const createSubscriptionResp = await createChannelSubscriptions([userId], channelId);

        expect(createSubscriptionResp.status).toBe(204);

        const clearResponse = await axios.delete(API.CHANNEL_SUBSCRIPTIONS.CLEAR_BY_USER_ID(userId));

        expect(clearResponse.status).toBe(204);

        const publishFn = (index: number) => publishToChannel(channelId, index);

        const publishTimes = 0;

        await publishAndTrackEvents(userId, publishTimes, publishFn)
    })

    test("events aren't received after clearing subscriptions by channel id", async () => {
        const userId = getRandomId();
        const channelId = getRandomId();

        const createSubscriptionResp = await createChannelSubscriptions([userId], channelId);

        expect(createSubscriptionResp.status).toBe(204);

        const clearResponse = await axios.delete(API.CHANNEL_SUBSCRIPTIONS.CLEAR_BY_CHANNEL_ID(userId));

        expect(clearResponse.status).toBe(204);

        const publishFn = (index: number) => publishToChannel(channelId, index);

        const publishTimes = 0;

        await publishAndTrackEvents(userId, publishTimes, publishFn)
    })

    test("events aren't received after subscriptions deletion by channel and user id", async () => {
        const userId = getRandomId();
        const channelId = getRandomId();

        const createSubscriptionResp = await createChannelSubscriptions([userId], channelId);

        expect(createSubscriptionResp.status).toBe(204);

        const clearResponse = await axios.post(API.CHANNEL_SUBSCRIPTIONS.BATCH_REMOVE, { userIds: [userId], channelId: channelId });

        expect(clearResponse.status).toBe(204);

        const publishFn = (index: number) => publishToChannel(channelId, index);

        const publishTimes = 0;

        await publishAndTrackEvents(userId, publishTimes, publishFn)
    })

    test("throws error if try to add two the same subscriptions", async () => {
        const userId = getRandomId();
        const channelId = getRandomId();

        const addSubscriptionFn = () => createChannelSubscriptions([userId], channelId);

        const addSubscriptionResp = await addSubscriptionFn();

        expect(addSubscriptionResp.status).toBe(204)

        const error = await getError(async () => addSubscriptionFn());

        expect(error).toBeInstanceOf(AxiosError);
    })

    test("publish message to channel successfully", async () => {
        const publishTimes = _.random(20, 100);

        const userId = getRandomId();

        const channelId = getRandomId();

        const createSubscriptionsResp = await createChannelSubscriptions([userId], channelId);

        expect(createSubscriptionsResp.status).toBe(204)

        const publishFn = (index: number) => publishToChannel(channelId, index);

        await publishAndTrackEvents(userId, publishTimes, publishFn)
    })
})

// required with this small example
// to make the isolatedModules config happy
export {}
