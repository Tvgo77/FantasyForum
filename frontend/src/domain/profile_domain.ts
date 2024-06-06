import { z } from "zod"
import { Profile } from "./user_domain"
import { BioRhyme } from "next/font/google"

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

export const ProfileRespSchema = z.object({
    profile: z.object({
        name: z.string(),
        bio: z.string(),
        birthday: z.string()
    })
})

export type UpdateProfileRequest = {
    profile: Profile
}

export type UpdateProfileResponse = {
    message: string
}

export interface ProfileUsecase {
    GetToken: () => string
}

