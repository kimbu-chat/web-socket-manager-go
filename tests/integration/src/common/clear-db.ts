import {DB_HOST, DB_NAME, DB_PASSWORD, DB_PORT, DB_USER} from "./environment";

const cleaner = require('postgres-cleaner')
import {Client} from "pg";

const options = {
    type: 'truncate',
    skipTables: ['gorp_migrations']
}

const cn = {
    host: DB_HOST,
    port: +DB_PORT,
    user: DB_USER,
    password: DB_PASSWORD,
    database: DB_NAME
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
