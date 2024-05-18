
export type loginUIform = {
    email: string
    password: string
}

export interface loginUsecase {
    storeToken(token: string): void
}