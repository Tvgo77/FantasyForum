import * as domain from '@/domain';

class signupController {
    public signupUsecase: domain.SignupUsecase

    constructor(su: domain.SignupUsecase) {
        this.signupUsecase = su
    }

    public async Signup(formData: domain.SignupUIform): Promise<boolean> {
        return false
    }
}

export function NewSignupController(su: domain.SignupUsecase): signupController {
    const sc: signupController = new signupController(su)
    return sc
}
