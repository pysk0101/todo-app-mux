import express from "express";
const todoRouter = express.Router();
import {getTodosById, getTodos, deleteTodo, changeTodo, completeTodo} from "../controllers/todo-controller.js";


//! Crate Todo is on the "../controllers/user-controller.js" file
todoRouter.patch("/:id", completeTodo);
todoRouter.put("/:id", changeTodo);
todoRouter.delete("/:id", deleteTodo);
todoRouter.get("/:id", getTodosById);
todoRouter.get("/", getTodos);

export default todoRouter;