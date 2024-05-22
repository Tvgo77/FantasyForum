import * as domain from '@/domain';

class loginUsecase {
    //env: undefined,
    public StoreToken(token: string) {
        localStorage.setItem("token", token)
    } 
}

export function NewLoginUsecase(/*env: undefined*/): domain.LoginUsecase {
    //loginUsecase.env = env
    const lu: loginUsecase = new loginUsecase()
    return lu
}
