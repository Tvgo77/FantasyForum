import { z } from 'zod';

export type loginUIform = {
    email: string
    password: string
}

export type loginRequest = {
    email: string
    password: string
}

export type loginResponse = {
    token: string
}

export const loginRespSchema = z.object({
    token: z.string(),
});

export interface loginUsecase {
    storeToken(token: string): void
}