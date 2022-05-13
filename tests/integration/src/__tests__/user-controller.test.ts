import axios from "axios";

describe("users controller", () => {
    test("publish message to users successfully", async () => {
        const response = await axios.post('http://localhost:8080/api/users/publish', { userIds: [2], message: {test: 1} });
        expect(response.status).toBe(204)
    })
})

// required with this small example
// to make the isolatedModules config happy
export {}