import Centrifuge from "centrifuge";
import {AxiosResponse} from "axios";
import _ from "lodash";
import jwt from "jsonwebtoken";
import {CENTRIFUGO_HMAC_SECRET, CENTRIFUGO_URL} from "./environment";
import WebSocket from "ws";

export const connect = async (userId: number): Promise<Centrifuge> => {
    const token = jwt.sign({sub: userId.toString()}, CENTRIFUGO_HMAC_SECRET);

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

const waitEvents = async (connection: Centrifuge, times: number): Promise<unknown[]> => {

    let counter = 0;
    const receivedObjects: unknown[] = [];

    return new Promise(function (fulfilled, rejected) {
        connection.on('publish', (eventData) => {
            receivedObjects.push(eventData);
            counter++;
            if(counter === times){
                fulfilled(receivedObjects);
            }
        })

        connection.on('disconnect', () => {
            rejected();
        });
    })
}

const waitForDisconnect = async (connection: Centrifuge): Promise<void> => {
    return new Promise(function (fulfilled, rejected) {
        connection.on('disconnect', () => {
            fulfilled();
        });
    })
}

const closeConnection = async (connection: Centrifuge): Promise<void> => {
    const disconnectPromise = waitForDisconnect(connection);

    connection.disconnect();

    await disconnectPromise;
}

export const publishAndTrackEvents = async (userId: number,
                                            publishTimes: number,
                                            publishFn : (data: number) => Promise<AxiosResponse<void>>): Promise<void> => {

    const connection = await connect(userId);

    const waitEventsPromise = waitEvents(connection, publishTimes)

    const publishedMessages: number[] = [];

    for (let i = 0; i < publishTimes; i++){
        publishedMessages.push(i);

        const publishMessageResp = await publishFn(i);

        expect(publishMessageResp.status).toBe(204)
    }

    const receivedMessages = await waitEventsPromise;

    const messagesAreEqual = _.isEqual(publishedMessages.sort(), receivedMessages.sort());

    expect(messagesAreEqual).toBe(true)

    await closeConnection(connection);
}



