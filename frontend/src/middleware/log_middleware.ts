import { NextResponse } from "next/server";
import { MiddlewareContext } from "./middleware_chainer";

export async function LogMiddleware(c: MiddlewareContext): Promise<NextResponse> {
    let logString: string = ""
    logString += (c.request.ip as string) + " | "
    logString += c.request.method + " | "
    logString += c.request.url + " | "

    const response = await c.Next()

    
    logString += response.status.toString()
    console.log(logString)
    return response 
}