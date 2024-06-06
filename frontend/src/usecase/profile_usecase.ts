import * as domain from '@/domain';
import Cookies from 'js-cookie'

class profileUsecase {
    public GetToken(): string {
        const token = Cookies.get("token")
        if (token) {
            return token
        }
        return ""
    }
}

export function NewProfileUsecase(): domain.ProfileUsecase {
    return new profileUsecase()
}