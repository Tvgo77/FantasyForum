import * as domain from '@/domain';
import Cookies from 'js-cookie'

class loginUsecase {
    //env: undefined,
    public StoreToken(token: string) {
        // localStorage.setItem("token", token)
        Cookies.set("token", token)
    } 
}

export function NewLoginUsecase(/*env: undefined*/): domain.LoginUsecase {
    //loginUsecase.env = env
    const lu: loginUsecase = new loginUsecase()
    return lu
}
