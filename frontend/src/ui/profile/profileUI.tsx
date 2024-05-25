"use client"

import React, { useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
// import 'bootstrap/dist/js/bootstrap.bundle';

import * as domain from '@/domain';

type ProfileUIprops = {
  formData: domain.ProfileUIform
  uid: string
}

export const ProfileUI: React.FC<ProfileUIprops> = ({formData, uid}) => {
  return (
    <div>
            
    </div>
  )
}