export interface User {
    id: number;
    dateCreated: Date;
    username: string;
    password: string;
}

export interface Session {
    id: number;
    dateCreated: Date;
    username: string;
    iat: number;
    exp: number;
}

export type PartialSession = Omit<Session, "iat" | "exp">;

export interface EncodeResult {
    token: string;
    exp: number;
    iat: number;
}

export type DecodeResult =
    | { type: "valid"; session: Session; }
    | { type: "integrity-error"; }
    | { type: "invalid-token"; };

export type ExpirationStatus = "expired" | "active" | "grace";