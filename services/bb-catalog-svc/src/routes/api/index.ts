import express from "express";
import users from "../../controllers/users-controller";

var router = express.Router();

// TODO I'm not sure about this anymore
const GET = router.get.bind(router);
const POST = router.post.bind(router);

/* Endpoint definitions */
GET("/users", users.getUsers);

export default router;
