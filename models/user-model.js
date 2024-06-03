import Joi from "@hapi/joi"
import fs from "fs"
const userSchema = Joi.object({
    id : Joi.number(),
    name: Joi.string().required(),
    todos: Joi.array()
})

const users = JSON.parse(fs.readFileSync("./db/user-database.json"))

export { userSchema, users }