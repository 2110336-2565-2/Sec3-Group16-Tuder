import React, { useContext } from "react";
import styled from "styled-components";
import Footer from "../components/global/Footer.js";
import { useParams } from "react-router-dom";
import { useQuery } from "react-query";
import { IsUser } from "../components/IsAuth";
import ClassDetails from "../components/ClassDetail";

export default function ClassDetail() {

  const { id } = useParams();
  

  return (
    <IsUser>
      <Container>
        <ClassDetails classDetail={id} />
      </Container>
      <Footer />
    </IsUser>
  );
}

const Container = styled.div`
  margin-top: 3rem;
  margin-bottom: 3rem;
  padding-left: 7%;
  padding-right: 7%;
`;
