import { userSchema, users } from "../models/user-model.js";



const deleteUser = async (req, res) => {
    const user = users.find(user => user.id === parseInt(req.params.id))
    if (!user) return res.status(404).send("The user with the given ID was not found.")
    const index = users.indexOf(user)
    users.splice(index, 1)
    res.send(`The user with the ID ${req.params.id} was deleted.`)
}


const createUser = async (req, res) => {
    try {
   
        const { error } = userSchema.validate(req.body)
        if (error) return res.status(400).send(error.details[0].message)
        const user = {
            id: users[users.length - 1].id + 1,
            name: req.body.name,
            todos: []
        }
        users.push(user)
        res.send(user)
    }
    catch (err){
        res.status(500).send(err.message)
    } 

}

const getUsersById = async (req, res) => {
    const user = users.find(user => user.id === parseInt(req.params.id))
    if (!user) return res.status(404).send("The user with the given ID was not found.")
    res.send(user)
}

const getUsers = (req, res) => {
    res.send(users)
}


export { getUsersById, getUsers, createUser, deleteUser }