import { z } from 'zod';

export type LoginUIform = {
    email: string
    password: string
}

export type LoginRequest = {
    email: string
    password: string
}

export type LoginResponse = {
    token: string
}

export const LoginRespSchema = z.object({
    token: z.string(),
});

export interface LoginUsecase {
    storeToken: (token: string) => void
}