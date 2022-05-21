import axios from "axios";
import {API_BASE} from "../common/environment";
import {
    publishAndTrackEvents,
} from "../common/websockets";
import _ from "lodash";

describe("users controller", () => {
    test("ensure published events are received", async () => {

        const publishTimes = _.random(10, 30);

        const userId = _.random(1, 100_000);

        const publishFn = (data: number) => axios.post(`${API_BASE}/api/users/publish`, { userIds: [userId], message: data});

        await publishAndTrackEvents(userId, publishTimes, publishFn)
    })
})

// required with this small example
// to make the isolatedModules config happy
export {}

