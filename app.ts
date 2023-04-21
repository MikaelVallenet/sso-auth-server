import express from 'express';
import { Request, Response, NextFunction } from "express";
import jwt from "jsonwebtoken";

require('dotenv').config();

const app = express();
const port = 3000;

const SECRET_KEY = process.env.SECRET_KEY

const getToken = (req: Request, res: Response, next: NextFunction) => {
    const token = jwt.sign({
        preferred_username: "mikatech",
        email: "mikael@berty.tech",
        email_verified: true,
        given_name: "Mikael",
        family_name: "Vallenet",
    }, SECRET_KEY, {expiresIn: "1h"})
    res.json({token})
}

app.listen(port, () => {
    console.log(`Server running on port ${port}`);
});

app.get('/token', getToken);
