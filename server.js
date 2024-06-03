import express from 'express';
const server = express();

server.use(express.json());

import todoRouter from "./router/todo-router.js";

server.use("/todos", todoRouter);


server.listen(3000, () => { "Server is running on port 3000" })