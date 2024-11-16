export const categorySchema = {
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
