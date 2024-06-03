import {todoSchema,todos} from "../models/todo-model.js";


const changeTodo = (req,res) => {
    const todo = todos.find(todo => todo.id === parseInt(req.params.id))
    if (!todo) return res.status(404).send("The todo with the given ID was not found.")
    const {error} = todoSchema.validate(req.body)
    if (error) return res.status(400).send(error.details[0].message)
    let oldTodo = Object.assign({}, todo)
    todo.mission = req.body.mission
    res.send(oldTodo, todo)
}

const deleteTodo= (req,res) => {
    const todo = todos.find(todo => todo.id === parseInt(req.params.id))
    if (!todo) return res.status(404).send("The todo with the given ID was not found.")
    const index = todos.indexOf(todo)
    todos.splice(index, 1)
    res.send(`The todo with the ID ${req.params.id} was deleted.`)
}

const createTodo = (req,res) => {
    const {error} = todoSchema.validate(req.body)
    if (error) return res.status(400).send(error.details[0].message)
    const todo = {
        id: todos.length + 1,
        mission: req.body.mission,
        isCompleted: req.body.isCompleted
    }
    todos.push(todo)
    res.send(todo)

}

const getTodosById = (req,res) =>{
    const todo = todo.find(todo => todo.id === parseInt(req.params.id))
    if (!todo) return res.status(404).send("The todo with the given ID was not found.")
    res.send(todo)
}

const getTodos = (req, res) => {
    res.send(todos)
}


export {getTodosById, getTodos, createTodo, deleteTodo, changeTodo}