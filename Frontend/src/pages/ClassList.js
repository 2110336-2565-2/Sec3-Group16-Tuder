import React from "react";
import { IsUser } from "../components/IsAuth";
import styled from "styled-components";
import ClassComponentList from "../components/ClassComponentList";

export default function ClassList() {
  return (
    <>
      <IsUser />
      <ContainerWithHeight margintop="25px">
        <ClassComponentList />
      </ContainerWithHeight>
    </>
  );
}

const ContainerWithHeight = styled.div`
  display: flex;
  flex-direction: column;
  padding: 0px 40px;
  margin-top: ${(props) => {
    return props.margintop;
  }};
  justify-content: center;
`;
