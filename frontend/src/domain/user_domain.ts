
export type RFC3339string = string

export type Profile = {
    name: string
    bio: string
    birthday: RFC3339string
}


export type User = {
    uid: number
    email: string
    profile: Profile
}
