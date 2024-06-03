import {todoSchema,todos} from "../models/todo-model.js";

const getTodosById = (req,res) =>{
    const todo = todo.find(todo => todo.id === parseInt(req.params.id))
    if (!todo) return res.status(404).send("The todo with the given ID was not found.")
    res.send(todo)
}

const getTodos = (req, res) => {
    res.send(todos)
}


export {getTodosById, getTodos}