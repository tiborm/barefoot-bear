import { dbConnect, dbDisconnect, db } from "./connect";
import { seedCategories } from "./categories";

const forceSeed = process.env.FORCED_SEED === "true" || process.argv.slice(2).includes("--forced-seed") || false;

(async () => {
    await dbConnect();
    await seedCategories(db, forceSeed);
    await dbDisconnect();
})();
