import { categoriesJSON } from "./json-reader.js"

const COLLECTION_NAME = "categories"

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
}

const checkCategoriesCollection = async (db) => {
    try {
        const collectionsResult = await db.listCollections({}, { nameOnly: true }).toArray()

        if (collectionsResult.find(collection => collection.name === "categories")) {
            console.log("Collection categories exists")
        } else {
            await db.createCollection("categories")
            console.log("Collection categories created")
        }
    } catch (error) {
        console.trace("Error checking collections", error)
        process.exit(1)
    }
}

const applySchemaValidation = async (db, collectionName, schema) => {
    try {
        await db.command({
            collMod: collectionName,
            validator: schema
        })

        console.log("Schema created successfully")
    } catch (error) {
        console.trace("Error creating schema", error)
        process.exit(1)
    }

}

const insertCategories = async (db, categories) => {
    try {
        const result = await db.collection(COLLECTION_NAME).insertMany(categories)
        console.log(`${result.insertedCount} categories inserted successfully`)
    } catch (error) {
        console.trace("Error inserting categories", error)
        process.exit(1)
    }
}



const seedCategories = async (db) => {
    await checkCategoriesCollection(db)
    await applySchemaValidation(db, COLLECTION_NAME, schema)
    await insertCategories(db, categoriesJSON)
}

export { seedCategories }