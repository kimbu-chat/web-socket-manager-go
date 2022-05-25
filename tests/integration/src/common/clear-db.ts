import {DB_PORT} from "./environment";

const cleaner = require('postgres-cleaner')
import {Client} from "pg";

const options = {
    type: 'truncate',
    skipTables: ['gorp_migrations']
}

const cn = {
    host: 'localhost',
    port: +DB_PORT,
    user: 'postgres',
    password: 'postgres',
    database: 'websocketmanager'
};

export default async () : Promise<void> => {

    const client = new Client(cn);

    try {
        await client.connect();
        await cleaner(options, client);
    } finally {
        // https://github.com/brianc/node-postgres/issues/2648
        if (client) {
            await client.end();
        }
    }
}
