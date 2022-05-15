import axios from "axios";
import {API_BASE} from "../common/environment";

describe("users controller", () => {
    test("publish message to users successfully", async () => {
        const response = await axios.post(`${API_BASE}/api/users/publish`, { userIds: [2], message: {test: 1} });
        expect(response.status).toBe(204)
    })
})

// required with this small example
// to make the isolatedModules config happy
export {}