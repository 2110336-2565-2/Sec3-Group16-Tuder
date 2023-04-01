import React from 'react'
import { InfoContainter, InfoTitle, InfoContent, InfoWrapper } from './ProfileStyle.js'

export default function StudentInfo({user}) {

  return (
    <InfoContainter>
      <InfoWrapper>
        <InfoTitle>Address</InfoTitle>
        <InfoContent>{user.address}</InfoContent>
      </InfoWrapper>
      <InfoWrapper>
        <InfoTitle>Contact Number</InfoTitle>
        <InfoContent>{user.phone}</InfoContent>
      </InfoWrapper>
      <InfoWrapper>
        <InfoTitle>Gender</InfoTitle>
        <InfoContent>{user.gender}</InfoContent>
      </InfoWrapper>
      <InfoWrapper>
        <InfoTitle>Birth Date</InfoTitle>
        <InfoContent>{user.birthdate?user.birthdate.split("T")[0]:""}</InfoContent>
      </InfoWrapper>
    </InfoContainter>
  )
}
