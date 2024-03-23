import express from "express";
import users from "../../controllers/users-controller";

var router = express.Router();

interface EndpointDefinition {
    endpoint: string | RegExp; // path params not captured if endpoint is a RegExp
    method: Function;
    handler: Function;
}

const GET = router.get.bind(router);
const POST = router.post.bind(router);
const PUT = router.put.bind(router);
const DELETE = router.delete.bind(router);

/* Endpoint definitions */
const endpoints: EndpointDefinition[] = [
    { method: GET, endpoint: '/users', handler: users.getUsers, },
    { method: GET, endpoint: '/users/:id', handler: users.getUsersByID, },
];

/* Register endpoints */
endpoints.forEach((endpoint: EndpointDefinition) => {
    endpoint.method(endpoint.endpoint, endpoint.handler);
});

export default router;
