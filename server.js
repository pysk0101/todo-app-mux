import express from 'express';
import path from "path"
import { fileURLToPath } from 'url'


const server = express();

server.use(express.static(path.join("./src")))
const __filename = fileURLToPath(import.meta.url)
const __dirname = path.dirname(__filename)

server.use(express.json());

import todoRouter from "./router/todo-router.js";
import userRouter from "./router/user-router.js";



server.get("/deneme", (req, res) => {
    res.sendFile(path.join(__dirname, "./src/index.html"));
});






server.use("/users", userRouter);
server.use("/todos", todoRouter);
server.get("/", (req, res) => {
    res.send("Welcome to the todo app")
})

server.listen(3000, () => { console.log("Server is running at http://localhost:3000")
 })