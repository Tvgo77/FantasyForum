import { NextRequest, NextResponse } from "next/server"

type middlewareFunc = (c: MiddlewareContext) => Promise<NextResponse>

export class MiddlewareContext  {
    public request: NextRequest
    // response: NextResponse
    public handlers: middlewareFunc[] = []
    index: number = -1

    constructor(req: NextRequest) {
        this.request = req
    }

    public async Next(): Promise<NextResponse> {
        this.index++
        let response: NextResponse
        if (this.index < this.handlers.length) {
            response = await this.handlers[this.index](this)
            return response
	    } else {
            return NextResponse.next({
                request: {
                    headers: new Headers(this.request.headers)
                }
            })
        }
    }

    public async Abort(): Promise<NextResponse> {
        return NextResponse.next({
            request: {
                headers: new Headers(this.request.headers)
            }
        })
    }
}
