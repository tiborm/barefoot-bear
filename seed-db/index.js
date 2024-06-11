import { dbConnect, dbDisconnect, db } from "./connect.js";
import { seedCategories } from "./categories.js";

const forceSeed = process.env.FORCED_SEED === "true" || process.argv.slice(2).includes("--forced-seed") || false;

(async () => {
    await dbConnect();
    await seedCategories(db, forceSeed);
    await dbDisconnect();
})();
