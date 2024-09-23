export const inventorySchema = {
    "$jsonSchema": {
        "bsonType": "object",
        "required": ["classUnitKey", "itemKey"],
        "properties": {
            "classUnitKey": {
                "bsonType": "object",
                "required": ["classUnitCode", "classUnitType"],
                "properties": {
                    "classUnitCode": {
                        "bsonType": "string",
                        "description": "must be a string and is required"
                    },
                    "classUnitType": {
                        "bsonType": "string",
                        "description": "must be a string and is required"
                    }
                }
            },
            "itemKey": {
                "bsonType": "object",
                "required": ["itemNo", "itemType"],
                "properties": {
                    "itemNo": {
                        "bsonType": "string",
                        "description": "must be a string and is required"
                    },
                    "itemType": {
                        "bsonType": "string",
                        "description": "must be a string and is required"
                    }
                }
            }
        }
    }
};
