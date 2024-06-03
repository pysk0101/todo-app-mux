import {todoSchema,todos} from "../models/todo-model.js";


const changeTodo =  async (req,res) => {
    const todo = todos.find(todo => todo.id === parseInt(req.params.id))
    if (!todo) return res.status(404).send("The todo with the given ID was not found.")
    const {error} = todoSchema.validate(req.body)
    if (error) return res.status(400).send(error.details[0].message)
    // let oldTodo = Object.assign({}, todo)
    todo.mission = req.body.mission
    res.send( todo)
}

const deleteTodo= async (req,res) => {
    const todo = todos.find(todo => todo.id === parseInt(req.params.id))
    if (!todo) return res.status(404).send("The todo with the given ID was not found.")
    const index = todos.indexOf(todo)
    todos.splice(index, 1)
    res.send(`The todo with the ID ${req.params.id} was deleted.`)
}

const completeTodo = async (req,res) => {
    const todo = todos.find(todo => todo.id === parseInt(req.params.id))
    if (!todo) return res.status(404).send("The todo with the given ID was not found.")
    todo.isCompleted = !todo.isCompleted
    res.send(todo)
}

const createTodo =  async (req,res) => {
    const {error} = todoSchema.validate(req.body)
    if (error) return res.status(400).send(error.details[0].message)
    const todo = {
        id: todos[todos.length-1].id + 1,
        mission: req.body.mission,
        isCompleted: false
    }
    todos.push(todo)
    res.send(todo)

}

const getTodosById = async (req,res) =>{
    const todo = todos.find(todo => todo.id === parseInt(req.params.id))
    if (!todo) return res.status(404).send("The todo with the given ID was not found.")
    res.send(todo)
}

const getTodos = (req, res) => {
    res.send(todos)
}


export {getTodosById, getTodos, createTodo, deleteTodo, changeTodo, completeTodo}