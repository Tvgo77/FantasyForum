import * as domain from '@/domain';
import { errorBadResponse, fetchWithRetry } from '@/helper';
import { NewLoginUsecase } from '@/usecase';
import { error } from 'console';



export async function login(formData: domain.loginUIform) {
    const lu: domain.loginUsecase = NewLoginUsecase()

    // Send login http request
    const request: domain.loginRequest = {email: formData.email, password: formData.password}
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
        return
    }

    // Handle response
    // Check Response format

    if (response.status != 200) {
        try{domain.errorRespSchema.parse(data)} catch (e) {alert(errorBadResponse); return}
        alert((data as domain.errorResponse).message)
        return
    }

    try{domain.loginRespSchema.parse(data)} catch (e) {alert(errorBadResponse); return}
    const token: string = (data as domain.loginResponse).token

    // Store token
    lu.storeToken(token)
}