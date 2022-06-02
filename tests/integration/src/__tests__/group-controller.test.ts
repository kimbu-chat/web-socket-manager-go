import axios, {AxiosError} from "axios";
import {API_BASE, CENTRIFUGO_URL, CENTRIFUGO_HMAC_SECRET, API} from "../common/environment";
import {publishAndTrackEvents} from "../common/websockets";
import {createGroupSubscriptions, getError, getRandomId, publishToGroup} from "../common/utils";
import _ from "lodash";

describe("groups controller", () => {
    test("events aren't received after clearing subscriptions by user id", async () => {
        const userId = getRandomId();
        const groupId = getRandomId();

        const createSubscriptionResp = await createGroupSubscriptions([userId], groupId);

        expect(createSubscriptionResp.status).toBe(204);

        const clearResponse = await axios.delete(API.GROUP_SUBSCRIPTIONS.CLEAR_BY_USER_ID(userId));

        expect(clearResponse.status).toBe(204);

        const publishFn = (index: number) => publishToGroup(groupId, index);

        const publishTimes = 0;

        await publishAndTrackEvents(userId, publishTimes, publishFn)
    })

    test("events aren't received after clearing subscriptions by group id", async () => {
        const userId = getRandomId();
        const groupId = getRandomId();

        const createSubscriptionResp = await createGroupSubscriptions([userId], groupId);

        expect(createSubscriptionResp.status).toBe(204);

        const clearResponse = await axios.delete(API.GROUP_SUBSCRIPTIONS.CLEAR_BY_GROUP_ID(userId));

        expect(clearResponse.status).toBe(204);

        const publishFn = (index: number) => publishToGroup(groupId, index);

        const publishTimes = 0;

        await publishAndTrackEvents(userId, publishTimes, publishFn)
    })

    test("events aren't received after subscriptions deletion by group and user id", async () => {
        const userId = getRandomId();
        const groupId = getRandomId();

        const createSubscriptionResp = await createGroupSubscriptions([userId], groupId);

        expect(createSubscriptionResp.status).toBe(204);

        const clearResponse = await axios.post(API.GROUP_SUBSCRIPTIONS.BATCH_REMOVE, { userIds: [userId], groupId: groupId });

        expect(clearResponse.status).toBe(204);

        const publishFn = (index: number) => publishToGroup(groupId, index);

        const publishTimes = 0;

        await publishAndTrackEvents(userId, publishTimes, publishFn)
    })

    test("throws error if try to add two the same subscriptions", async () => {
        const userId = getRandomId();
        const groupId = getRandomId();

        const addSubscriptionFn = () => createGroupSubscriptions([userId], groupId);

        const addSubscriptionResp = await addSubscriptionFn();

        expect(addSubscriptionResp.status).toBe(204)

        const error = await getError(async () => addSubscriptionFn());

        expect(error).toBeInstanceOf(AxiosError);
    })

    test("publish message to group successfully", async () => {
        const publishTimes = _.random(20, 100);

        const userId = getRandomId();

        const groupId = getRandomId();

        const createSubscriptionsResp = await createGroupSubscriptions([userId], groupId);

        expect(createSubscriptionsResp.status).toBe(204)

        const publishFn = (index: number) => publishToGroup(groupId, index);

        await publishAndTrackEvents(userId, publishTimes, publishFn)
    })
})

// required with this small example
// to make the isolatedModules config happy
export {}
