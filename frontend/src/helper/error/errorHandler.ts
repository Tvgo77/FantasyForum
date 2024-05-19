export const errorBadResponse: string = "Server Error: Bad response"

export async function fetchWithRetry(
    url: string | URL | globalThis.Request,
    init?: RequestInit,
    retries: number = 3,
    delay: number = 1000
): Promise<Response> {
    let response: Response
    for (let i = 0; i < retries; i++) {
        try {
            response = await fetch(url, init)
            return response
        } catch (error) {
            if (i === retries -1) throw error
            await new Promise(res => setTimeout(res, delay))
        }
    }
    throw new Error("Make typescript happy")  // Make typescript happy
}