import React from 'react'
import { InfoContainter, InfoTitle, InfoContent, InfoWrapper } from './profileStyle.js'

export default function StudentInfo({user}) {
  return (
    <InfoContainter>
      <InfoWrapper>
        <InfoTitle>Address</InfoTitle>
        <InfoContent>{user.address}</InfoContent>
      </InfoWrapper>
      <InfoWrapper>
        <InfoTitle>Contact Number</InfoTitle>
        <InfoContent>{user.contactNumber}</InfoContent>
      </InfoWrapper>
      <InfoWrapper>
        <InfoTitle>Gender</InfoTitle>
        <InfoContent>{user.gender}</InfoContent>
      </InfoWrapper>
      <InfoWrapper>
        <InfoTitle>Birth Date</InfoTitle>
        <InfoContent>{user.birthdate}</InfoContent>
      </InfoWrapper>
      <InfoWrapper>
        <InfoTitle>School</InfoTitle>
        <InfoContent>{user.school}</InfoContent>
      </InfoWrapper>
    </InfoContainter>
  )
}
