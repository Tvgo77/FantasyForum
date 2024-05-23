import {describe, it, expect, vi, beforeEach} from "vitest"
import { mock, when, anyString } from 'ts-mockito';
import * as domain from '@/domain'
import { NewSignupController } from ".";
import * as helper from "@/helper"


describe("TestSignup", function() {
    // Shared variable among test case
    const msu: domain.SignupUsecase = mock<domain.SignupUsecase>()
    when(msu.StoreToken(anyString())).thenReturn()
    const sc = NewSignupController(msu)

    beforeEach(function() {
        global.alert = vi.fn((message: string) => {console.log("Alert: %s", message)})  // simulate alert() in browser
    })

    it("Signup Success", async function() {
        // Setup Fetch() simulation behaviour
        const mockResponseBody = {message: "Success", token: "token string"}
        const mockResponse = {ok: true, status: 200, json: async () => mockResponseBody} as Response
        vi.spyOn(helper, "fetchWithRetry").mockResolvedValue(mockResponse)
        const methodMonitor = vi.spyOn(msu, "StoreToken")

        // Test
        const result = await sc.Signup({email: "test@gmail.com", password: "password"})
        
        // Assert
        expect(result).toBe(true)
        expect(methodMonitor).toBeCalledWith(mockResponseBody.token)
    })

    it("Signup Failed", async function() {
        // Setup Fetch() simulation behaviour
        const mockResponseBody = {message: "Error happened"}
        const mockResponse = {ok: false, status: 409, json: async () => mockResponseBody} as Response
        vi.spyOn(helper, "fetchWithRetry").mockResolvedValue(mockResponse)
        const methodMonitor = vi.spyOn(msu, "StoreToken")

        // Test
        const result = await sc.Signup({email: "test@gmail.com", password: "password"})
        expect(result).toBe(false)
        expect(global.alert).toBeCalledWith(mockResponseBody.message)
    })
})