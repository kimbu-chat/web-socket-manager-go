import axios, {AxiosError} from "axios";
import {API_BASE, CENTRIFUGO_URL, CENTRIFUGO_HMAC_SECRET} from "../common/environment";
import {closeCentrifugoConnection, connectToCentrifugo, waitEvents} from "../common/websockets";
import {getError} from "../common/utils";

describe("groups controller", () => {
    test("events aren't received after subscription deletion", async () => {
        const userId = 1;

        const clearResponse = await axios.post(`${API_BASE}/api/groups/subscriptions/clear-by-user-id`, { userId });

        expect(clearResponse.status).toBe(204);

        const response = await axios.post(`${API_BASE}/api/groups/subscriptions`, { userIds: [userId], groupId: 1 });

        expect(response.status).toBe(204);


    })

    test("throws error if try to add two the same subscriptions", async () => {
        const userId = 1;

        const clearResponse = await axios.post(`${API_BASE}/api/groups/subscriptions/clear-by-user-id`, { userId });

        expect(clearResponse.status).toBe(204);

        const request = { userIds: [userId], groupId: 1 };

        const addSubscriptionFn = () => axios.post(`${API_BASE}/api/groups/subscriptions`, request);

        const addSubscriptionResp = await addSubscriptionFn();

        expect(addSubscriptionResp.status).toBe(204)

        const error = await getError(async () => addSubscriptionFn());

        expect(error).toBeInstanceOf(AxiosError);
    })

    

    test("publish message to group successfully", async () => {
        const userId = 1;

        const clearResponse = await axios.post(`${API_BASE}/api/groups/subscriptions/clear-by-user-id`, { userId: userId });

        expect(clearResponse.status).toBe(204)

        const createSubscriptionsResp = await axios.post(`${API_BASE}/api/groups/subscriptions`, { userIds: [userId], groupId: 1 });

        expect(createSubscriptionsResp.status).toBe(204)

        const connection = await connectToCentrifugo(userId);

        const publishTimes = 1;

        const waitEventsPromise = waitEvents(connection, publishTimes)

        const publishMessageResp = await axios.post(`${API_BASE}/api/groups/publish`, { groupId: 1, message: '123' });

        expect(publishMessageResp.status).toBe(204)

        const publishedTimes = await waitEventsPromise;

        expect(publishedTimes).toBe(publishTimes)

        await closeCentrifugoConnection(connection);
    })
})

// required with this small example
// to make the isolatedModules config happy
export {}
