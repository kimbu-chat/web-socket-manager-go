import jwt from "jsonwebtoken";
import {CENTRIFUGO_HMAC_SECRET, CENTRIFUGO_URL} from "./environment";
import Centrifuge from "centrifuge";
import WebSocket from "ws";

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
