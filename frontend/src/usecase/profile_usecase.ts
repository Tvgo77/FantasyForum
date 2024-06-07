import * as domain from '@/domain';
import { ReadonlyRequestCookies } from 'next/dist/server/web/spec-extension/adapters/request-cookies';

class profileUsecase {
    public Token: string

    constructor (t: string) {
        this.Token = t
    }

    public GetToken(): string {
        return this.Token
    }
}

export function NewProfileUsecase(t: string): domain.ProfileUsecase {
    return new profileUsecase(t)
}