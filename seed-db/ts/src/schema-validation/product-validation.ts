export const productSchema = {
    $jsonSchema: {
        "bsonType": "object",
        "required": ["name", "typeName", "id"],
        "properties": {
            "name": {
                "bsonType": "string",
                "description": "must be a string and is required"
            },
            "typeName": {
                "bsonType": "string",
                "description": "must be a string and is required"
            },
            "id": {
                "bsonType": "string",
                "description": "must be a string and is required"
            },
        }
    }
};