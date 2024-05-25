
export type RFC3339string = string

export type Profile = {
    name: string
    bio: string
    birthdaty: RFC3339string
}


export type User = {
    uid: number
    email: string
    profile: Profile
}
