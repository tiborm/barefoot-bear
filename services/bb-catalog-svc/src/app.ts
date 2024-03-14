import config from "./config";
import express from "express";
import cookieParser from "cookie-parser";
import logger from "morgan";

import apiRouter from "./routes/api";

const app = express();

app.use(logger(config.logs.format));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());

// Mount the API router at /api/v1
app.use("/api/v1", apiRouter);

module.exports = app;
