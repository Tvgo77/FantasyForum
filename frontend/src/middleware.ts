import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'
import { MiddlewareContext } from './middleware/middleware_chainer'
import { LogMiddleware } from './middleware/log_middleware'
import { JWTmiddleware } from './middleware/jwt_middleware'

export async function middleware(req: NextRequest) {
    let c = new MiddlewareContext(req)
    const globalMiddleware = [LogMiddleware, JWTmiddleware]
    c.handlers = globalMiddleware

    const response = await c.Next()
    return response
}

export const config = {
    matcher: [
      /*
       * Match all request paths except for the ones starting with:
       * - api (API routes)
       * - _next/static (static files)
       * - _next/image (image optimization files)
       * - favicon.ico (favicon file)
       */
      '/((?!api|_next/static|_next/image|favicon.ico).*)',
    ],
}