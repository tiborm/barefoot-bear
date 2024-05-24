import { dbConnect, dbDisconnect, db } from "./connect.js"
import { seedCategories } from "./categories.js"

(async () => {
    await dbConnect()
    await seedCategories(db)
    await dbDisconnect()
})()




// TODO CHECK if data exists
// TODO exit with error if ENV var not suggest force seeding
// TODO write file to mark DB seeded (warn if file exists and suggest force seeding or delete file)
// TODO create database
// TODO seed data