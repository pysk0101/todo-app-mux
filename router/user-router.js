import express from "express";
const userRouter = express.Router();


import {getUsersById, getUsers, createUser, deleteUser} from "../controllers/user-controller.js"
import {createTodo} from "../controllers/todo-controller.js"


userRouter.post("/add", createTodo); //
userRouter.delete("/:id", deleteUser);
userRouter.post("/", createUser);

userRouter.get("/:id", getUsersById);
userRouter.get("/", getUsers);

export default userRouter;