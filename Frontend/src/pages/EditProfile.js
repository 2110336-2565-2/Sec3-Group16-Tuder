import React from "react";
import styled from "styled-components";
import { dummyStudent, dummyTutor } from "../datas/Profile.role";

import TutorFormEditProfile from "../components/profile/TutorFormEditProfile";

export default function EditProfile() {
  const user = dummyTutor;

  return (
    <Container>
      <TutorFormEditProfile user={user} />
    </Container>
  );
}

const Container = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 50px 0px;
  background-color: #fdedeb;
`;
