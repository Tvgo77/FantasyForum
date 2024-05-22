"use client"

import React, { useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
// import 'bootstrap/dist/js/bootstrap.bundle';

import { NewLoginController } from '@/controller';
import { NewLoginUsecase } from '@/usecase';
import * as domain from '@/domain';

export const LoginUI = () => {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")

  const [loading, setLoading] = useState(false)

  async function handleSubmit(event: React.FormEvent): Promise<void> {
    event.preventDefault()
    setLoading(true)

    const formData: domain.LoginUIform = { email, password }
    const lu = NewLoginUsecase()
    const lc = NewLoginController(lu)
    
    try {
      await lc.login(formData)  // Controller function
    } catch (error) {
      alert(error)
    }

    setLoading(false)
  }

  return (
    <>
      <button type="button" className="btn btn-primary" data-bs-toggle="modal" data-bs-target="#loginModal">
        Login
      </button>

      <div className="modal fade" id="loginModal" tabIndex={-1} aria-labelledby="exampleModalLabel" aria-hidden="true">
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
                    onChange={(e) => { setEmail(e.target.value) }}
                    required
                  />
                </div>
                <div className="mb-3">
                  <label htmlFor="exampleInputPassword1" className="form-label">Password</label>
                  <input
                    type="password"
                    className="form-control"
                    id="exampleInputPassword1"
                    onChange={(e) => { setPassword(e.target.value) }}
                    required
                  />
                </div>
                <div className="mb-3 form-check">
                  <input type="checkbox" className="form-check-input" id="exampleCheck1" />
                  <label className="form-check-label" htmlFor="exampleCheck1">Remeber me</label>
                </div>
                <button type="submit" className="btn btn-primary" data-bs-dismiss="modal">Login</button>
                {/* <button type="button" className="btn btn-outline-primary mx-4">Signup</button> */}
              </form>
            </div>
            {/* <div className="modal-footer">
            <button type="button" className="btn btn-secondary" data-bs-dismiss="modal">Close</button>
            <button type="button" className="btn btn-primary">Save changes</button>
          </div> */}
          </div>
        </div>
      </div>
    </>
  );
};