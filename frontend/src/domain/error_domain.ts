import { z } from 'zod';

export type errorResponse = {
    message: string
}

export const errorRespSchema = z.object({
    message: z.string(),
});