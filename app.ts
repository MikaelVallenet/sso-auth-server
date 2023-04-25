import express from 'express';
import { Request, Response, NextFunction } from "express";
import jwt from "jsonwebtoken";
import fs from "fs";
import { createPublicKey } from 'crypto'

require('dotenv').config();

const app = express();
const port = 3000;

const SECRET_KEY = process.env.SECRET_KEY
const CORS_FRONT_URL = process.env.CORS_FRONT_URL

const privateKEY = fs.readFileSync('private.key', 'utf8');
const publicKEY = createPublicKey(fs.readFileSync('public.key', 'utf8'));

const cors = require('cors');
app.use(cors({
    origin: CORS_FRONT_URL
}));

const getToken = (req: Request, res: Response, next: NextFunction) => {
    const payload = {
        preferred_username: "mikalelito",
        email: "mikael@berty.tech",
        email_verified: true,
        given_name: "Mikael",
        family_name: "Vallenet",
    }

    const signOptions: jwt.SignOptions = {
        issuer: "http://localhost:3000",
        subject: "532cb4a4-7ad7-40a5-826a-c8272af2d9f3",
        expiresIn: "30s",
        algorithm: "RS256"
    }


    const token = jwt.sign( payload, privateKEY, signOptions)
    res.json({token})
}

const getPublicKey = (req: Request, res: Response, next: NextFunction) => {
    res.json(publicKEY.export({ format: 'jwk' }));
}

app.listen(port, () => {
    console.log(`Server running on port ${port}`);
});

app.get('/token', getToken);
app.get('/public-key', getPublicKey);
