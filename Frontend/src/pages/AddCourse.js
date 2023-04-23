import React from "react";
import styled from "styled-components";
import FormAddCourse from "../components/form/FormAddCourse";
import { IsAdmin } from "../components/IsAuth";

export default function Addcourse() {
  return (
    <Container>
      <IsAdmin>
        <FormAddCourse />
      </IsAdmin>
    </Container>
  );
}

const Container = styled.div``;
