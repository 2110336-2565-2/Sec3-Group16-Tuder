import React from 'react'
import { InfoContainter, InfoTitle, InfoContent, InfoWrapper } from './ProfileStyle.js'
import styled from 'styled-components'

export default function TutorInfo({user}){
  return (
    <Container>
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
          <InfoContent>{user.birth_date}</InfoContent>
        </InfoWrapper>
        <InfoWrapper>
          <InfoTitle>Citizen ID</InfoTitle>
          <InfoContent>{user.citizen_id}</InfoContent>
        </InfoWrapper>
        <InfoWrapper isSpan={true}>
          <InfoTitle>Description</InfoTitle>
          <InfoContent>{user.description}</InfoContent>
        </InfoWrapper>
      </InfoContainter>
      <Button>My Reviews</Button>
    </Container>
  )
}

const Container = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
`;

const Button = styled.button`
  width: 300px;
  height: 50px;
  margin-top: 55px;
  background-color: #FFFFFF;
  border: 2px solid #000000;
  border-radius: 10px;
  color: #000000;
  text-align: center;
  font-size: 16px;
  font-weight: medium;
  text-decoration: none;
  cursor: pointer;
  transition: 0.3s ease all;

  &:hover {
    background-color: #000000;
    color: #FFFFFF;
  }
`;
