import axios from "axios";
import {API_BASE} from "../common/environment";
import {closeCentrifugoConnection, connectToCentrifugo, waitEvents, waitForDisconnect} from "../common/websockets";

describe("users controller", () => {
    test("publish message to users successfully", async () => {

        const publishTimes = 1;

        const userId = 1;

        let connection = await connectToCentrifugo(userId)

        const waitEventsPromise = waitEvents(connection, publishTimes)
    
        await axios.post(`${API_BASE}/api/users/publish`, { userIds: [userId], message: { "1": 2 }});
    
        const publishedTimes = await waitEventsPromise;

        expect(publishedTimes).toBe(publishTimes)

        await closeCentrifugoConnection(connection);
    })
})

// required with this small example
// to make the isolatedModules config happy
export {}

