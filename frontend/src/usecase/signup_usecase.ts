import * as domain from '@/domain';

class signupUsecase {
    public StoreToken(token: string) {
        localStorage.setItem("token", token)
    }
}

export function NewSignupUsecase(): domain.SignupUsecase {
    const su = new signupUsecase()
    return su
}