import express from 'express';
const server = express();

server.use(express.json());

server.listen(3000, () => { "Server is running on port 3000" })