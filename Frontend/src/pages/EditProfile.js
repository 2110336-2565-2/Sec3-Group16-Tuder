import React from "react";
import styled from "styled-components";
import useRole from "../hooks/useRole";
import { dummyStudent, dummyTutor } from "../datas/Profile.role";

import FormEditProfile from "../components/profile/FormEditProfile";
import { IsUser } from "../components/IsAuth";

export default function EditProfile() {
  // CHANGE THIS TO GET USER FROM BACKEND
  const [role, handleRole] = useRole();
  const user = role === "student" ? dummyStudent : dummyTutor;

  return (
    <IsUser>
      <Container>
        <FormEditProfile user={user} />
      </Container>
    </IsUser>
  );
}

const Container = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 50px 0px;
  background-color: #fdedeb;
`;
