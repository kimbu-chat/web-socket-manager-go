import {DB_PORT} from "./environment";

const cleaner = require('postgres-cleaner')
const pgp = require('pg-promise')({
    // Initialization Options
})

const options = {
    type: 'truncate',
    skipTables: ['SequelizeMeta'],
}

const cn = {
    host: 'localhost',
    port: +DB_PORT,
    user: 'postgres',
    password: 'postgres',
};

export default (): Promise<void> => {
    const db = pgp(cn)

    return cleaner(options, db);
}
