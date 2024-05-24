import * as domain from '@/domain';
import Cookies from 'js-cookie'

class signupUsecase {
    public StoreToken(token: string) {
        // localStorage.setItem("token", token)
        Cookies.set("token", token)
    }
}

export function NewSignupUsecase(): domain.SignupUsecase {
    const su = new signupUsecase()
    return su
}