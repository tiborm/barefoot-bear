import { dbConnect, dbDisconnect, db } from "./connect.js";
import { checkCollection, applySchemaValidation, prepareForInsertion, insertJSONDocs } from "./seeder.js";
import { readFilesInDirectory, readSingleFile } from "./json-reader.js";
import { categorySchema } from "./schema-validation/category-validation.js";
import { productSchema } from "./schema-validation/product-validation.js";
import { inventorySchema } from "./schema-validation/inventory-validation.js";
import { extractInventory, extractProducts } from "./data-extractors.js";

const isForced = process.env.FORCED_SEED === "true" || process.argv.slice(2).includes("--forced-seed") || false;

(async () => {
    const categoriesJSON: {}[] = await readSingleFile("../json/categories.json");

    const extractedProducts = extractProducts(await readFilesInDirectory("../json/products"));
    const extractedInventory = extractInventory(await readFilesInDirectory("../json/inventory"));

    await dbConnect();

    let collectionName = "categories";
    await checkCollection(db, collectionName);
    await applySchemaValidation(db, collectionName, categorySchema);
    await prepareForInsertion(db, collectionName, isForced);
    await insertJSONDocs(db, collectionName, categoriesJSON);

    collectionName = "products";
    await checkCollection(db, collectionName);
    await applySchemaValidation(db, collectionName, productSchema);
    await prepareForInsertion(db, collectionName, isForced);
    await insertJSONDocs(db, collectionName, extractedProducts);

    collectionName = "inventories";
    await checkCollection(db, collectionName);
    await applySchemaValidation(db, collectionName, inventorySchema);
    await prepareForInsertion(db, collectionName, isForced);
    await insertJSONDocs(db, collectionName, extractedInventory);

    await dbDisconnect();
})();
