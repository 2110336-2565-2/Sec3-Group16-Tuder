import React from "react";
import styled from "styled-components";
import MyClass from "./MyClass";

export default function ClassComponentList() {

    return (
    <Container>
        <h1>My Classes</h1>
        <ClassListPage>
            {/* use map for generate each class */}
            <MyClass/>
        </ClassListPage>
    </Container>
    );
}

const Container = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 5px;
  margin-bottom: 25px;
`;

const ClassListPage = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: center;
  padding: 10px;
  margin-top: 10px;
`;
