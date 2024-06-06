"use client"

import React, { useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
// import 'bootstrap/dist/js/bootstrap.bundle';

import * as domain from '@/domain';

type ProfileUIprops = {
  formData: domain.ProfileUIform
  uid: string
  isMyself: boolean
}

export const ProfileUI: React.FC<ProfileUIprops> = ({formData, uid, isMyself}) => {
  return (
    <div>
      <form onSubmit={() => {}}>
        <div className="mb-3">
          <label htmlFor="profileInputName" className="form-label">Name</label>
          <input
            type="text"
            className="form-control"
            id="profileInputName"
            value={formData.name}
            onChange={() => {}}
          />
        </div>
        <div className="mb-3">
          <label htmlFor="profileInputBio" className="form-label">Bio</label>
          <input
            type="text"
            className="form-control"
            id="profileInputBio"
            value={formData.bio}
            onChange={() => {}}
          />
        </div>
        <div className="mb-3">
          <label htmlFor="profileInputBirthday" className="form-label">Birthday</label>
          <input
            type="text"
            className="form-control"
            id="profileInputBirthday"
            value={formData.birthdayDate}
            onChange={() => {}}
          />
        </div>
        {isMyself && <button type="submit" className="btn btn-primary" data-bs-dismiss="modal">Update</button>}
      </form>
    </div>
  )
}