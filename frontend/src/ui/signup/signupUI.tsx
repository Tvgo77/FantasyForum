"use client"

import React, { useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
// import 'bootstrap/dist/js/bootstrap.bundle';

import { NewLoginController, NewSignupController } from '@/controller';
import { NewLoginUsecase } from '@/usecase';
import * as domain from '@/domain';
import { set } from 'zod';
import { NewSignupUsecase } from '@/usecase/signup_usecase';

export const SignupUI = () => {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [isFocusedOnce, setFocusedOnce] = useState(false)
  const [isPasswordValid, setPasswordValid] = useState(false)

  const [loading, setLoading] = useState(false)

  async function handleSubmit(event: React.FormEvent): Promise<void> {
    event.preventDefault()
    setLoading(true)

    const formData: domain.SignupUIform = {email: email, password: password}
    const su = NewSignupUsecase()
    const sc = NewSignupController(su)

    try {
      await sc.Signup(formData)  // Controller function
    } catch (error) {
      alert(error)
    }
    

    setLoading(false)
  }

  function checkPassword(event: React.ChangeEvent<HTMLInputElement>) {
    const password2: string = event.target.value as string
    if (password != password2 || !password2) {
      setPasswordValid(false)
    } else {
      setPasswordValid(true)
    }
  }

  return (
    <>
      <button type="button" className="btn btn-primary" data-bs-toggle="modal" data-bs-target="#signupModal">
        Signup
      </button>

      <div className="modal fade" id="signupModal" tabIndex={-1} aria-labelledby="exampleModalLabel" aria-hidden="true">
        <div className="modal-dialog">
          <div className="modal-content">
            <div className="modal-body">
              <button type="button" className="btn-close position-absolute top-0 mt-1 end-0 me-1" data-bs-dismiss="modal" aria-label="Close" />
              <form onSubmit={handleSubmit}>
                <div className="mb-3">
                  <label htmlFor="exampleInputEmail1" className="form-label">Email address</label>
                  <input
                    type="email"
                    className="form-control"
                    id="exampleInputEmail1"
                    aria-describedby="emailHelp"
                    onChange={(e) => {setEmail(e.target.value) }}
                    required
                  />
                </div>

                <div className="mb-3">
                  <label htmlFor="exampleInputPassword1" className="form-label">Password</label>
                  <input
                    type="password"
                    className="form-control"
                    id="exampleInputPassword1"
                    onChange={(e) => {setPassword(e.target.value)}}
                    required
                  />
                </div>

                <div className="mb-3">
                  <label htmlFor="exampleInputPassword2" className="form-label">Confirm Password</label>
                  <input
                    type="password"
                    className= {"form-control" + ((isFocusedOnce && !isPasswordValid) ? (" is-invalid"): (""))}
                    id="exampleInputPassword2"
                    onFocus={() => {setFocusedOnce(true)}}
                    onChange={checkPassword}
                    required
                  />
                  <div className="invalid-feedback"> 2 Password Not Same </div>
                </div>

                <button 
                  type="submit" 
                  className="btn btn-primary" 
                  data-bs-dismiss="modal"
                  disabled={!isPasswordValid}
                >
                  Signup
                </button>
              </form>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};