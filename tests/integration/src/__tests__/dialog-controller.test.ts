import axios, {AxiosError} from "axios";
import {API} from "../common/environment";
import {publishAndTrackEvents} from "../common/websockets";
import {
    createDialogSubscriptions,
    getError,
    getRandomId, publishToChannel, publishToInterlocutors,
} from "../common/utils";
import _ from "lodash";

describe("dialog controller", () => {
    test("events aren't received after clearing subscriptions by user initiator id", async () => {
        const userId = getRandomId();
        const initiatorId = getRandomId();

        const createSubscriptionResp = await createDialogSubscriptions([userId], initiatorId);

        expect(createSubscriptionResp.status).toBe(204);

        const clearResponse = await axios.delete(API.DIALOG_SUBSCRIPTIONS.CLEAR_BY_INITIATOR_ID(initiatorId));

        expect(clearResponse.status).toBe(204);

        const publishFn = (index: number) => publishToInterlocutors(initiatorId, index);

        const publishTimes = 0;

        await publishAndTrackEvents(userId, publishTimes, publishFn)
    })

    test("events aren't received after subscriptions deletion by initiatorId and user id", async () => {
        const userId = getRandomId();
        const initiatorId = getRandomId();

        const createSubscriptionResp = await createDialogSubscriptions([userId], initiatorId);

        expect(createSubscriptionResp.status).toBe(204);

        const clearResponse = await axios.post(API.DIALOG_SUBSCRIPTIONS.BATCH_REMOVE, { userIds: [userId], initiatorId: initiatorId });

        expect(clearResponse.status).toBe(204);

        const publishFn = (index: number) => publishToInterlocutors(initiatorId, index);

        const publishTimes = 0;

        await publishAndTrackEvents(userId, publishTimes, publishFn)
    })

    test("throws error if try to add two the same subscriptions", async () => {
        const userId = getRandomId();
        const initiatorId = getRandomId();

        const addSubscriptionFn = () => createDialogSubscriptions([userId], initiatorId);

        const addSubscriptionResp = await addSubscriptionFn();

        expect(addSubscriptionResp.status).toBe(204)

        const error = await getError(async () => addSubscriptionFn());

        expect(error).toBeInstanceOf(AxiosError);
    })

    test("publish message to users successfully", async () => {
        const publishTimes = _.random(20, 100);

        const userId = getRandomId();

        const initiatorId = getRandomId();

        const createSubscriptionsResp = await createDialogSubscriptions([userId], initiatorId);

        expect(createSubscriptionsResp.status).toBe(204)

        const publishFn = (index: number) => publishToInterlocutors(initiatorId, index);

        await publishAndTrackEvents(userId, publishTimes, publishFn)
    })
})

// required with this small example
// to make the isolatedModules config happy
export {}
