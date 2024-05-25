import { NewProfileController } from '@/controller/profile_controller';
import { LoginUI, ProfileUI } from '@/ui';
import { headers } from 'next/headers';
import React from 'react';

export default async function Page({params}: {params: {uid: string}}) {
  const headersList = headers()
  const hasNoAuth = headersList.has("no-auth")

  // Check authentication
  if (hasNoAuth) {
    return (
      <div className='container'>
        <div> Please Login </div>
        <LoginUI></LoginUI>
      </div>
    )
  }

  // Fetch profile data
  const pc = NewProfileController()
  try {
    var formData = await pc.FetchProfile(params.uid)
  } catch (e) {
    return <div> Fail to retrieve user profile </div>
  }

  return (
    <div className="container">
      <ProfileUI formData={formData} uid={params.uid} ></ProfileUI>   
    </div>
  )
}