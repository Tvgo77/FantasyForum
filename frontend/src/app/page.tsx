import { headers } from "next/headers";
import {LoginUI, SignupUI} from "../ui"
import React from 'react';
import Link from "next/link";

export default function Page() {
  const headersList = headers()
  const hasNoAuth = headersList.has("no-auth")
  const uid = headersList.get("uid")

  return (
    <div className="container">
      <LoginUI/>
      <SignupUI/>
      <div className="container">
        {hasNoAuth? "Guest": (<Link href={`/profile/${uid}`}>{"User: " + uid}</Link>)}
      </div>
    </div>
  );
}

