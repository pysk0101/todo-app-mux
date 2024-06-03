import express from "express";
const todoRouter = express.Router();
import {getTodosById, getTodos, createTodo, deleteTodo, changeTodo} from "../controllers/todo-controller.js";


todoRouter.put("/:id", changeTodo);
todoRouter.delete("/:id", deleteTodo);
todoRouter.post("/", createTodo);

todoRouter.get("/:id", getTodosById);
todoRouter.get("/", getTodos);

export default todoRouter;