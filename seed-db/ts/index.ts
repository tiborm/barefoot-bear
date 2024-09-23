import { dbConnect, dbDisconnect, db } from "./src/connect";
import { seedCategories } from "./src/categories";

const forceSeed = process.env.FORCED_SEED === "true" || process.argv.slice(2).includes("--forced-seed") || false;

(async () => {
    await dbConnect();
    await seedCategories(db, forceSeed);
    await dbDisconnect();
})();
