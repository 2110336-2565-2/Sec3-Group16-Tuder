
import React from "react";
import Footer from "../components/global/Footer.js";

import styled from "styled-components";
import { IsUser } from "../components/IsAuth.js";
import CancelRequest from "../components/form/FormCancelRequest.js";

export default function UserCancelRequest() {
  
  
    return (
      <Container>
        <IsUser>
        <CancelRequest/>
           
        <Footer />
        </IsUser>
      </Container>
      
    );
  }

const Container = styled.div`
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    background-color: #fdedeb;
`;


