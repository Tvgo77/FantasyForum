import { z } from 'zod';

export type SignupUIform = {
    email: string
    password: string
}

export type SignupRequest = {
    email: string
    password: string
}

export type SignupResponse = {
    message: string
    token: string
}

// Response format checker
export const SignupRespSchema = z.object({
    message: z.string(),
    token: z.string(),
});

export interface SignupUsecase {
    StoreToken: (token: string) => void
}
