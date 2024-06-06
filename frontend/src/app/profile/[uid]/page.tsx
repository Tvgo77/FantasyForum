import { NewProfileController } from '@/controller/profile_controller';
import { LoginUI, ProfileUI } from '@/ui';
import * as domain from '@/domain';
import { headers } from 'next/headers';
import React from 'react';
import { NewProfileUsecase } from '@/usecase/profile_usecase';

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
  const pu: domain.ProfileUsecase = NewProfileUsecase()
  const pc = NewProfileController(pu)
  try {
    var formData = await pc.FetchProfile(params.uid)
  } catch (e) {
    return <div> Fail to retrieve user profile </div>
  }

  const uid_jwt = headersList.get("uid")
  const isMyself = (uid_jwt == params.uid)

  return (
    <div className="container">
      <ProfileUI formData={formData} uid={params.uid} isMyself={isMyself} ></ProfileUI>   
    </div>
  )
}