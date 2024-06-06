import {describe, it, expect, vi, beforeEach} from "vitest"
import { mock, when, anyString } from 'ts-mockito';
import { NewProfileController } from "./profile_controller";
import * as domain from '@/domain'
import * as helper from "@/helper"
import { resourceLimits } from "worker_threads";
import { error } from "console";


describe("TestProfile", function() {
    // Mock Profile Usecase
    const mpu: domain.ProfileUsecase = mock<domain.ProfileUsecase>()
    when(mpu.GetToken()).thenReturn("test-cookies")
    const pc = NewProfileController(mpu)

    beforeEach(function() {
        global.alert = vi.fn((message: string) => {console.log("Alert: %s", message)})  // simulate alert() in browser
    })

    it("FetchProfileSuccess", async function() {
        // Simulate Fetch()
        const mockResponseBody: domain.FetchProfileResponse = {
            profile: {
                name: "test-user",
                bio: "I'm a test user",
                birthday: "2000-05-06T12:00:00Z"
            } 
        }
        const mockResponse = {ok: true, status: 200, json: async () => mockResponseBody} as Response
        vi.spyOn(global, "fetch").mockResolvedValue(mockResponse)

        // Test
        const result = await pc.FetchProfile("6")

        // Assert
        expect(result.name).toBe("test-user")
        expect(result.birthdayDate).toBe("2000-05-06")
    })

    it("FetchProfileFail", async function() {
        // Simulate Fetch()
        const mockResponseBody: domain.ErrorResponse = {message: "Error in test"}
        const mockResponse = {ok: false, status: 404, json: async () => mockResponseBody} as Response
        vi.spyOn(global, "fetch").mockResolvedValue(mockResponse)

        // test
        try {
            const result = await pc.FetchProfile("404")
        } catch (e) {
            expect(e).toBeInstanceOf(Error)
        }
        // Assert
        expect(global.alert).toBeCalledWith(mockResponseBody.message)
    })

    it("UpdateProfileSuccess", async function() {
        // Simulate Fetch()
        const mockResponseBody: domain.UpdateProfileResponse = {message: "Update success"}
        const mockResponse = {ok: true, status: 200, json: async () => mockResponseBody} as Response
        vi.spyOn(global, "fetch").mockResolvedValue(mockResponse)

        // Test
        const formData: domain.ProfileUIform = {name: "test user", bio: "I'm a test user", birthdayDate: "2000-05-06"}
        const result = await pc.UpdateProfile("6", formData)

        // Assert
        expect(result).toBe(true)
    })

    it("UpdateProfileFail", async function() {
        // Simulate Fetch()
        const mockResponseBody: domain.ErrorResponse = {message: "Error in test"}
        const mockResponse = {ok: false, status: 401, json: async () => mockResponseBody} as Response
        vi.spyOn(global, "fetch").mockResolvedValue(mockResponse)

        // Test
        const formData: domain.ProfileUIform = {name: "test user", bio: "I'm a test user", birthdayDate: "2000-05-06"}
        const result = await pc.UpdateProfile("1", formData)

        // Assert
        expect(result).toBe(false)
        expect(global.alert).toBeCalledWith(mockResponseBody.message)
    })
})