import * as domain from '@/domain';
import { errorBadResponse, fetchWithRetry } from '@/helper';

class signupController {
    public signupUsecase: domain.SignupUsecase

    constructor(su: domain.SignupUsecase) {
        this.signupUsecase = su
    }

    public async Signup(formData: domain.SignupUIform): Promise<boolean> {
        // Send signup http request
        const req: domain.SignupRequest = {email: formData.email, password: formData.password}
        let resp: Response
        let data: any
        
        try {
            resp = await fetchWithRetry("http://localhost:8080/login", {
                method: "POST",
                mode: "cors",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(req)
            })
            data = await resp.json()
        } catch (error) {
            alert((error as Error).message)
            return false
        }

        // Handle http response
        if (resp.status != 200) {
            try {domain.errorRespSchema.parse(data)} catch (e) {alert(errorBadResponse); return false}
            alert((data as domain.errorResponse).message)
            return false
        }

        try {domain.SignupRespSchema.parse(data)} catch (e) {alert(errorBadResponse); return false}
        const token = (data as domain.SignupResponse).token

        // Store token
        this.signupUsecase.StoreToken(token)
        return true
    }
}

export function NewSignupController(su: domain.SignupUsecase): signupController {
    const sc: signupController = new signupController(su)
    return sc
}
