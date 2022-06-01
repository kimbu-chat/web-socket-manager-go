import {
    publishAndTrackEvents,
} from "../common/websockets";
import _ from "lodash";
import clearDb from "../common/clear-db";
import {getRandomId, publishToUsers} from "../common/utils";

describe("users controller", () => {
    test("ensure all published events are received by a user", async () => {
        const publishTimes = _.random(10, 30);

        const userId = getRandomId();

        const publishFn = (data: number) => publishToUsers([userId], data);

        await publishAndTrackEvents(userId, publishTimes, publishFn)
    })
})

// required with this small example
// to make the isolatedModules config happy
export {}

