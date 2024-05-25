import { Profile } from "./user_domain"

export type ProfileUIform = {
    name: string
    bio: string
    birthdayDate: string
}


export type FetchProfileRequest = {

}

export type FetchProfileResponse = {
    profile: Profile
}

export type UpdateProfileRequest = {
    profile: Profile
}

export type UpdateProfileResponse = {
    message: string
}

