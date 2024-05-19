import * as domain from '@/domain';

let loginUsecase = {
    //env: undefined,
    storeToken: storeToken 
}

export function NewLoginUsecase(/*env: undefined*/): domain.LoginUsecase {
    //loginUsecase.env = env
    return loginUsecase
}

function storeToken(token: string) {
    localStorage.setItem("token", token)
}