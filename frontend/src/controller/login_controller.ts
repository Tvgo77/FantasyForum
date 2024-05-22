import * as domain from '@/domain';
import { errorBadResponse, fetchWithRetry } from '@/helper';
import { NewLoginUsecase } from '@/usecase';
import { error } from 'console';

class loginController {
    public loginUsecase: domain.LoginUsecase

    constructor(lu: domain.LoginUsecase) {
        this.loginUsecase = lu
    }

    public async Login(formData: domain.LoginUIform): Promise<boolean> {
        // Send login http request
        const request: domain.LoginRequest = {email: formData.email, password: formData.password}
        let response: Response 
        let data
    
        try {
            response = await fetchWithRetry("http://localhost:8080/login", {
                method: "POST",
                mode: "cors",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(request)
            })
    
            data = await response.json()
        } catch (error) {
            alert((error as Error).message)
            return false
        }
    
        // Handle response
        // Check Response format
    
        if (response.status != 200) {
            try{domain.errorRespSchema.parse(data)} catch (e) {alert(errorBadResponse); return false}
            alert((data as domain.errorResponse).message)
            return false
        }
    
        try{domain.LoginRespSchema.parse(data)} catch (e) {alert(errorBadResponse); return false}
        const token: string = (data as domain.LoginResponse).token
    
        // Store token
        this.loginUsecase.StoreToken(token)
        return true
    }
} 

export function NewLoginController(lu: domain.LoginUsecase): loginController {
    const lc: loginController = new loginController(lu)
    return lc
}
