import { Request, Response, NextFunction } from "express";
import { encode, TAlgorithm } from "jwt-simple";
import {Session, PartialSession, EncodeResult} from "./utils";

const getToken = (req: Request, res: Response, next: NextFunction): EncodeResult => {
    const algorithm: TAlgorithm = "HS512";
    const issued = Date.now();
    const fifteenMinutes = 1000 * 60 * 15;
    const partialSession: PartialSession = {
        id: 1,
        dateCreated: new Date(),
        username: "mikatech",
    }
    const expires = issued + fifteenMinutes; // 15 minutes from now
    const session: Session = {
        ...partialSession,
        iat: issued,
        exp: expires
    }

    return {
        token: encode(session, "secret", algorithm),
        iat: issued,
        exp: expires
    };
}