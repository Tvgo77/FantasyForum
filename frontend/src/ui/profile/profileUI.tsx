"use client"

import React, { useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
// import 'bootstrap/dist/js/bootstrap.bundle';

import * as domain from '@/domain';
import { NewProfileController, ProfileController } from '@/controller/profile_controller';
import { NewProfileUsecase } from '@/usecase/profile_usecase';

type ProfileUIprops = {
  formData: domain.ProfileUIform
  uid: string
  isMyself: boolean
  token: string
}

export const ProfileUI: React.FC<ProfileUIprops> = ({formData, uid, isMyself, token}) => {
  const [name, setName] = useState(formData.name)
  const [bio, setBio] = useState(formData.bio)
  const [birthdayDate, setBirthday] = useState(formData.birthdayDate)

  async function handleSubmit(event: React.FormEvent): Promise<void> {  
    event.preventDefault()
    
    const formData: domain.ProfileUIform = {name: name, bio: bio, birthdayDate: birthdayDate}
    const pu = NewProfileUsecase(token)
    const pc = NewProfileController(pu)

    let result: boolean
    try {
      result = await pc.UpdateProfile(uid, formData)
    } catch (e) {
      result = false
      alert((e as Error).message)
    }

    if (result) {
      alert("Update Success")
    }
  }

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <div className="mb-3">
          <label htmlFor="profileInputName" className="form-label">Name</label>
          <input
            type="text"
            className="form-control"
            id="profileInputName"
            defaultValue={formData.name}
            onChange={(e) => {setName(e.target.value)}}
          />
        </div>
        <div className="mb-3">
          <label htmlFor="profileInputBio" className="form-label">Bio</label>
          <input
            type="text"
            className="form-control"
            id="profileInputBio"
            defaultValue={formData.bio}
            onChange={(e) => {setBio(e.target.value)}}
          />
        </div>
        <div className="mb-3">
          <label htmlFor="profileInputBirthday" className="form-label">Birthday</label>
          <input
            type="text"
            className="form-control"
            id="profileInputBirthday"
            defaultValue={formData.birthdayDate}
            onChange={(e) => {setBirthday(e.target.value)}}
          />
        </div>
        {isMyself && <button type="submit" className="btn btn-primary" data-bs-dismiss="modal">Update</button>}
      </form>
    </div>
  )
}