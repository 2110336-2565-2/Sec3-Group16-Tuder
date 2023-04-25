import React from "react";
import styled from "styled-components";
import FormAddCourse from "../components/form/FormAddCourse";
import { IsTutor } from "../components/IsAuth";

export default function Addcourse() {
  return (
    <Container>
      <IsTutor>
        <FormAddCourse />
      </IsTutor>
    </Container>
  );
}

const Container = styled.div``;
