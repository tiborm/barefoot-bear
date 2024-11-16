import { dbConnect, dbDisconnect, db } from "./connect.js";
import { checkCollection, applySchemaValidation, prepareForInsertion, insertJSONDocs } from "./seeder.js";
import { readFilesInDirectory, readSingleFile } from "./json-reader.js";
import { categorySchema } from "./schema-validation/category-validation.js";
import { productSchema } from "./schema-validation/product-validation.js";
import { inventorySchema } from "./schema-validation/inventory-validation.js";
import { extractInventory, extractProducts } from "./data-extractors.js";

const isForced = process.env.FORCED_SEED === "true" || process.argv.slice(2).includes("--forced-seed") || false;

(async () => {
    // FIXME extracting file paths to .env would be nice
    const categoriesJSON: {}[] = await readSingleFile("../../seed-data/categories.json");

    const extractedProducts = extractProducts(await readFilesInDirectory("../../seed-data/products"));
    const extractedInventory = extractInventory(await readFilesInDirectory("../../seed-data/inventory"));

    await dbConnect();

    let collectionName = "categories";
    await checkCollection(db, collectionName);
    await applySchemaValidation(db, collectionName, categorySchema);
    await prepareForInsertion(db, collectionName, isForced);
    await insertJSONDocs(db, collectionName, categoriesJSON);

    // FIXME products are duplicated in jsons, therefore no unique index can be applied
    // solved in the golang seeder
    collectionName = "products";
    await checkCollection(db, collectionName);
    await applySchemaValidation(db, collectionName, productSchema);
    await prepareForInsertion(db, collectionName, isForced);
    await insertJSONDocs(db, collectionName, extractedProducts);

    collectionName = "inventory";
    await checkCollection(db, collectionName);
    await applySchemaValidation(db, collectionName, inventorySchema);
    await prepareForInsertion(db, collectionName, isForced);
    await insertJSONDocs(db, collectionName, extractedInventory);

    await dbDisconnect();
})();
