import axios from "axios";
import {API, API_BASE} from "../common/environment";
import {
    publishAndTrackEvents,
} from "../common/websockets";
import _ from "lodash";
import clearDb from "../common/clear-db";

describe("users controller", () => {
    test("ensure all published events are received by a user", async () => {
        const publishTimes = _.random(10, 30);

        const userId = _.random(1, 100_000);

        const publishFn = (data: number) => axios.post(API.USERS.PUBLISH, { userIds: [userId], message: data});

        await publishAndTrackEvents(userId, publishTimes, publishFn)
    })
})

// required with this small example
// to make the isolatedModules config happy
export {}

