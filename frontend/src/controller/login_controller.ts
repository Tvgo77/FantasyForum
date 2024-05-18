import * as domain from '@/domain';
import { NewLoginUsecase } from '@/usecase';


export async function login(formData: domain.loginUIform) {
    const lu: domain.loginUsecase = NewLoginUsecase()

    // Send login http request

    // Handle response

    // Store token
    lu.storeToken("")
}