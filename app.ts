import express from 'express';
import { Request, Response, NextFunction } from "express";
import jwt from "jsonwebtoken";

require('dotenv').config();

const app = express();
const port = 3000;

const SECRET_KEY = process.env.SECRET_KEY
const CORS_FRONT_URL = process.env.CORS_FRONT_URL

const cors = require('cors');
app.use(cors({
    origin: CORS_FRONT_URL
}));

const getToken = (req: Request, res: Response, next: NextFunction) => {
    const token = jwt.sign({
        preferred_username: "mikalelito",
        email: "mikael@berty.tech",
        email_verified: true,
        given_name: "Mikael",
        family_name: "Vallenet",
        sub: "532cb4a4-7ad7-40a5-826a-c8272af2d9f3",
    }, SECRET_KEY, {expiresIn: "30s"})
    res.json({token})
}

app.listen(port, () => {
    console.log(`Server running on port ${port}`);
});

app.get('/token', getToken);
