import React from "react";
import { useNavigate } from "react-router-dom";
import {
  InfoContainter,
  InfoTitle,
  InfoContent,
  InfoWrapper,
} from "./ProfileStyle.js";
import styled from "styled-components";

export default function TutorInfo({ user, isOwner, othersUsername }) {
  const navigate = useNavigate();

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
          <InfoContent>
            {user.birthdate ? user.birthdate.split("T")[0] : ""}
          </InfoContent>
        </InfoWrapper>
        {isOwner ? (
          <InfoWrapper>
            <InfoTitle>Citizen ID</InfoTitle>
            <InfoContent>{user.citizen_id}</InfoContent>
          </InfoWrapper>
        ) : null}
        <InfoWrapper isSpan={true}>
          <InfoTitle>Description</InfoTitle>
          <InfoContent>{user.description}</InfoContent>
        </InfoWrapper>
      </InfoContainter>
      {isOwner ? <Button onClick={()=>navigate('reviews')} >My Reviews</Button> : <Button onClick={()=>navigate(`/tutors/${othersUsername}/reviews`)} >Read Reviews</Button>}
    </Container>
  );
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
  background-color: #ffffff;
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
    color: #ffffff;
  }
`;
