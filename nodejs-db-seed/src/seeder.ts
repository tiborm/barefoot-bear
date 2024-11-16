import { Collection, Db } from "mongodb";

const checkCollection = async (db: Db, collectionName: string) => {
    try {
        const collectionsResult = await db.listCollections({}, { nameOnly: true }).toArray();

        if (collectionsResult.find(collection => collection.name === collectionName)) {
            console.log(`Collection ${ collectionName } exists`);
        } else {
            await db.createCollection(collectionName);
            console.log(`Collection ${ collectionName } created`);
        }
    } catch (error) {
        console.trace("Error checking collections", error);
        process.exit(1);
    }
};

const checkIfDataExist = async (collection: Collection) => {
    try {
        if (await collection.countDocuments({}) > 0) {
            return true;
        } else {
            return false;
        }
    } catch (error) {
        console.trace("Error fetching data", error);
        process.exit(1);
    }
};

const prepareForInsertion = async (db: Db, collectionName: string, isForced: boolean) => {
    const collection = db.collection(collectionName);

    if (await checkIfDataExist(collection)) {
        if (isForced) {
            console.warn("Forced seeding");
            collection.deleteMany({});

            return;
        }

        console.error(
            `\nCategories data already exists.
            If you want to force seeding, please run the command with --forced-seed flag.
            or use FORCED_SEED=true npm start.`.replace(/^ +/gm, '')
        );
        process.exit(1);
    }
};

const applySchemaValidation = async (db: Db, collectionName: string, schema: {}) => {
    try {
        await db.command({
            collMod: collectionName,
            validator: schema
        });

        console.log("Validation schema applied successfully");
    } catch (error) {
        console.trace("Error creating schema", error);
        process.exit(1);
    }
};

const insertJSONDocs = async (db: Db, collectionName: string, data: {}[]) => {
    try {
        const result = await db.collection(collectionName).insertMany(data);
        console.log(`${result.insertedCount} ${ collectionName } inserted successfully`);
    } catch (error) {
        console.trace(`Error inserting ${ collectionName }`, error);
        process.exit(1);
    }
};

export { 
    checkCollection,
    applySchemaValidation,
    prepareForInsertion,
    insertJSONDocs,
};