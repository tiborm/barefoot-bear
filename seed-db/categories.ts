import { Collection, Db } from "mongodb";
import { categoriesJSON } from "./json-reader";

const collectionName = "categories";

const schema = {
    $jsonSchema: {
        bsonType: "object",
        required: ["id", "name", "subs"],
        properties: {
            id: {
                bsonType: "string",
                uniqueItems: true,
                description: "must be a string and is required"
            },
            name: {
                bsonType: "string",
                description: "must be a string and is required"
            },
            subs: {
                bsonType: "array",
                description: "must be an array and is required"
            }
        }
    }
};

const checkCategoriesCollection = async (db: Db) => {
    try {
        const collectionsResult = await db.listCollections({}, { nameOnly: true }).toArray();

        if (collectionsResult.find(collection => collection.name === "categories")) {
            console.log("Collection categories exists");
        } else {
            await db.createCollection("categories");
            console.log("Collection categories created");
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

const prepareForInsertion = async (db: Db, isForced: boolean) => {
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

const insertCategories = async (db: Db, categories: {}[]) => {
    try {
        const result = await db.collection(collectionName).insertMany(categories);
        console.log(`${result.insertedCount} categories inserted successfully`);
    } catch (error) {
        console.trace("Error inserting categories", error);
        process.exit(1);
    }
};

const seedCategories = async (db: Db, isForced = false) => {
    await checkCategoriesCollection(db);
    await applySchemaValidation(db, collectionName, schema);
    await prepareForInsertion(db, isForced);
    await insertCategories(db, categoriesJSON);
};

export { seedCategories };