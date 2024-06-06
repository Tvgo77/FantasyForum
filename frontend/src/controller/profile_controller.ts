import * as domain from '@/domain';


class profileController {
    public profileUsecase: domain.ProfileUsecase

    constructor(pu: domain.ProfileUsecase) {
        this.profileUsecase = pu
    }

    public async FetchProfile(uid: string): Promise<domain.ProfileUIform> {
        // Fetch profile from backend

        // Check Response format

        // Convert Response to form data format and return

        let x: any
        return x
    }

    public async UpdateProfile(uid: string, formData: domain.ProfileUIform): Promise<boolean> {
        // Convert form data to request format

        // Send request

        // Handle response

        return false
    }
}

export function NewProfileController(pu: domain.ProfileUsecase ): profileController {
    return new profileController(pu)
}