import { MongoClient } from 'mongodb';

const dbName = 'barefoot-bear';  // TODO workname of the project -> process.env.DB_NAME
const mongoUri = "mongodb://root:example@localhost:27017"; // TODO -> process.env.MONGO_URI;

const mongoClient = new MongoClient(mongoUri);

const dbConnect = async () => {
    try {
        await mongoClient.connect();
        console.log('Connected to database');
    } catch (error) {
        console.error('Error connecting to database', error);
        process.exit(1);
    }
};

const dbDisconnect = async () => {
    try {
        await mongoClient.close();
        console.log('Disconnected from database');
    } catch (error) {
        console.error('Error disconnecting from database', error);
    }
};

const db = mongoClient.db(dbName);

export { dbConnect, dbDisconnect, db, dbName };