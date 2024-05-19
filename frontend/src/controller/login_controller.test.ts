import * as helper from "@/helper"
import * as usecase from "@/usecase"
import { NewLoginController } from '@/controller';
import { NewLoginUsecase } from '@/usecase';
import {describe, it, expect, vi, beforeEach} from "vitest"
import { mock, when, anyString } from 'ts-mockito';
import * as domain from '@/domain'


describe("TestLogin", function() {
    // Shared variable throw all test
    const mlu: domain.LoginUsecase = mock<domain.LoginUsecase>()  // Mocked LoginUsecase interface
    when(mlu.storeToken(anyString())).thenReturn()
    const lc = NewLoginController(mlu)

    beforeEach(function() {
        global.alert = vi.fn((message: string) => {console.log("Alert: %s", message)})
    })

    it("Login Successful", async function() {
        // Setup
        const mockResponseBody = {token: "token string"}
        const mockResponse = {ok: true, status: 200, json: async () => mockResponseBody} as Response
        const mockFetch = vi.spyOn(helper, "fetchWithRetry").mockResolvedValue(mockResponse)

        // Test
        const result = await lc.login({email: "test@gmail.com", password: "password"})

        // Assert
        expect(result).toBe(true)
    })

    it("Login fail"), async function() {
        // Setup
        const mockResponseBody = {message: "Email or Password wrong"}
        const mockRespone = {ok: false, status: 401, json: async () => mockResponseBody} as Response
        const mockFetch = vi.spyOn(helper, "fetchWithRetry").mockResolvedValue(mockRespone)

        // Test
        const result = await lc.login({email: "test@gmail.com", password: "password"})

        // Assert
        expect(global.alert).toBeCalledWith("Email address or Password is wrong")
        expect(result).toBe(false)
    }
})