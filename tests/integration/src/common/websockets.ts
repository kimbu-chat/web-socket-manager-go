import jwt from "jsonwebtoken";
import {API_BASE, CENTRIFUGO_HMAC_SECRET, CENTRIFUGO_URL} from "./environment";
import Centrifuge from "centrifuge";
import WebSocket from "ws";
import axios, {AxiosResponse} from "axios";

export const connectToCentrifugo = async (userId: number): Promise<Centrifuge> => {
    const token = jwt.sign({ sub: userId.toString() }, CENTRIFUGO_HMAC_SECRET);

    const connection = new Centrifuge(CENTRIFUGO_URL, {
        debug: false,
        websocket: WebSocket
    });

    connection.setToken(token);

    connection.connect();

    return new Promise(function (fulfilled, rejected) {
        connection.on('connect', (ctx) => {
            fulfilled(connection);
        });

        connection.on('disconnect', () => {
            rejected();
        });
    })
}

export const waitEvents = async (connection: Centrifuge, times: number): Promise<number> => {

    let counter = 0;

    return new Promise(function (fulfilled, rejected) {
        connection.on('publish', () => {
            counter++;
            if(counter === times){
                fulfilled(counter);
            }
        })

        connection.on('disconnect', () => {
            rejected();
        });
    })
}

export const waitForDisconnect = async (connection: Centrifuge): Promise<void> => {
    return new Promise(function (fulfilled, rejected) {
        connection.on('disconnect', () => {
            fulfilled();
        });
    })
}

export const closeCentrifugoConnection = async (connection: Centrifuge): Promise<void> => {
    const disconnectPromise = waitForDisconnect(connection);

    connection.disconnect();

    await disconnectPromise;
}

export const publishAndTrackEvents = async (userId: number, publishFn : () => Promise<AxiosResponse<void>>): Promise<void> => {
    const connection = await connectToCentrifugo(userId);

    const publishTimes = 1;

    const waitEventsPromise = waitEvents(connection, publishTimes)

    const publishMessageResp = await publishFn;

    expect(publishMessageResp.status).toBe(204)

    const publishedTimes = await waitEventsPromise;

    expect(publishedTimes).toBe(publishTimes)

    await closeCentrifugoConnection(connection);
}



