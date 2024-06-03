import Joi from "@hapi/joi"
import fs from "fs"

const todoSchema = Joi.object({
    'id': Joi.number(),
    'mission': Joi.string().required(),
    "isCompleted": Joi.boolean().required(),
})

const todos = JSON.parse(fs.readFileSync("./db/todo-database.json"))

export { todoSchema, todos }

