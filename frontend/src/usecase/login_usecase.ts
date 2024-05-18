import * as domain from '@/domain';

export function NewLoginUsecase(): domain.loginUsecase {
    const lu: domain.loginUsecase = {
        storeToken: storeToken
    }
    return lu
}

function storeToken(token: string) {
    localStorage.setItem("token", token)
}