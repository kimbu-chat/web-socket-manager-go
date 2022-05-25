import axios, {AxiosError} from "axios";
import {API_BASE, CENTRIFUGO_URL, CENTRIFUGO_HMAC_SECRET, API} from "../common/environment";
import {publishAndTrackEvents} from "../common/websockets";
import {getError} from "../common/utils";
import _ from "lodash";

describe("groups controller", () => {
    test("events aren't received after subscription deletion", async () => {
        const userId = _.random(1, 100_000);
        const groupId = _.random(1, 100_000);

        const createSubscriptionResp = await axios.post(API.GROUP_SUBSCRIPTIONS.CREATE, { userIds: [userId], groupId: groupId });

        expect(createSubscriptionResp.status).toBe(204);

        const clearResponse = await axios.delete(API.GROUP_SUBSCRIPTIONS.CLEAR_BY_USER_ID(userId));

        expect(clearResponse.status).toBe(204);

        const publishFn = (index: number) => axios.post(API.GROUP_SUBSCRIPTIONS.PUBLISH, { groupId: groupId, message: index });

        const publishTimes = 0;

        await publishAndTrackEvents(userId, publishTimes, publishFn)
    })

    test("throws error if try to add two the same subscriptions", async () => {
        const userId = _.random(1, 100_000);
        const groupId = _.random(1, 100_000);

        const request = { userIds: [userId], groupId: groupId };

        const addSubscriptionFn = () => axios.post(API.GROUP_SUBSCRIPTIONS.CREATE, request);

        const addSubscriptionResp = await addSubscriptionFn();

        expect(addSubscriptionResp.status).toBe(204)

        const error = await getError(async () => addSubscriptionFn());

        expect(error).toBeInstanceOf(AxiosError);
    })

    test("publish message to group successfully", async () => {
        const publishTimes = _.random(20, 100);

        const userId = _.random(1, 100_000);

        const groupId = _.random(1, 100_000);

        const createSubscriptionsResp = await axios.post(API.GROUP_SUBSCRIPTIONS.CREATE, { userIds: [userId], groupId: groupId });

        expect(createSubscriptionsResp.status).toBe(204)

        const publishFn = (index: number) => axios.post(API.GROUP_SUBSCRIPTIONS.PUBLISH, { groupId: groupId, message: index });

        await publishAndTrackEvents(userId, publishTimes, publishFn)
    })
})

// required with this small example
// to make the isolatedModules config happy
export {}
