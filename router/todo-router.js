import express from "express";
const todoRouter = express.Router();
import {getTodosById, getTodos} from "../controllers/todo-controller.js";

todoRouter.get("/:id", getTodosById);
todoRouter.get("/", getTodos);

export default todoRouter;