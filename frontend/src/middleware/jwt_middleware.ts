import { NextResponse } from "next/server";
import { MiddlewareContext } from "./middleware_chainer";
import { jwtVerify, type JWTPayload } from 'jose';

export async function JWTmiddleware(c: MiddlewareContext): Promise<NextResponse> {
    // Extract token from cookie
    const token = c.request.cookies.get("token")
    if (!token) {
        // Set no-auth true in request header
        c.request.headers.set("no-auth", "true")
        return await c.Next()
    }
    
    // Verify token

    try {
        const secret = new TextEncoder().encode("secret");
        var decoded = await jwtVerify(token.value, secret)
    } catch (e) {
        c.request.headers.set("no-auth", "true")
        return await c.Next()
    }

    // Add uid in request cookie
    const uid = decoded.payload.sub as string
    c.request.headers.set("uid", uid)
    return await c.Next()
}