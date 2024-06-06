import * as domain from '@/domain';
import { errorBadResponse } from '@/helper';


class profileController {
    public profileUsecase: domain.ProfileUsecase

    constructor(pu: domain.ProfileUsecase) {
        this.profileUsecase = pu
    }

    public async FetchProfile(uid: string): Promise<domain.ProfileUIform> {
        // Get token in cookies
        const token = this.profileUsecase.GetToken()
        
        // Fetch profile from backend
        let resp: Response
        let data: any

        resp = await fetch("http://localhost:8080/profile/" + uid, {
            method: "GET",
            mode: "cors",
            headers: {
                "Authorization": "bear " + token
            }
        })
        data = await resp.json()

        // Handle response 
        // Check response format
        if (!resp.ok) {
            try {domain.ErrorRespSchema.parse(data)} catch (e) {alert(errorBadResponse); throw new Error()}
            alert((data as domain.ErrorResponse).message)
            throw new Error()
        } 
        
        try {domain.ProfileRespSchema.parse(data)} catch(e) {alert(errorBadResponse); throw new Error()}
        
        // Convert Response to form data format and return
        const formData: domain.ProfileUIform = {
            name: (data as domain.FetchProfileResponse).profile.name,
            bio: (data as domain.FetchProfileResponse).profile.bio,
            birthdayDate: (data as domain.FetchProfileResponse).profile.birthday.substring(0, 10)  // "YYYY-MM-DD"
        }

        return formData
    }

    public async UpdateProfile(uid: string, formData: domain.ProfileUIform): Promise<boolean> {
        // Convert form data to request format
        const req: domain.UpdateProfileRequest = {
            profile: {
                name: formData.name,
                bio: formData.bio,
                birthday: formData.birthdayDate
            }
        }

        // Send request
        const token = this.profileUsecase.GetToken()
        let resp: Response
        let data: any

        try {
            resp = await fetch("http://localhost:8080/profile/" + uid, {
                method: "POST",
                mode: "cors",
                headers: {
                    "Authorization": "bear " + token
                }
            })
            data = await resp.json()
        } catch (e) {
            alert((e as Error).message)
            return false
        }

        // Handle response
        // Check response format
        if (!resp.ok) {
            try {domain.ErrorRespSchema.parse(data)} catch (e) {alert(errorBadResponse); return false}
            alert((data as domain.ErrorResponse).message)
            return false
        } 
        return true
    }
}

export function NewProfileController(pu: domain.ProfileUsecase ): profileController {
    return new profileController(pu)
}