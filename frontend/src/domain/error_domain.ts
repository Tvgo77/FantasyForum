import { z } from 'zod';

export type ErrorResponse = {
    message: string
}

export const ErrorRespSchema = z.object({
    message: z.string(),
});