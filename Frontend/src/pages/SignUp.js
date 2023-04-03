import React, { Fragment } from "react";
import styled from "styled-components";
import WaveFooter from "../components/global/WaveFooter.js";
import { IsGuest } from "../components/IsAuth.js";
import FormSignUp from "../components/form/FormSignUp.js";

export default function SignUp() {
  return (
    <Fragment>
      <IsGuest>
        <ContainerWithHeight margintop="0px">
          <SeperateSection>
            <BlackDiv>
              <SignUpInfo>
                Getting <br />
                Started With <br />
                Matching system
              </SignUpInfo>
              <Image src="/images/signupDecorator.svg" alt="decorator" />
            </BlackDiv>
            <FormWrapper>
              <FormSignUp />
            </FormWrapper>
          </SeperateSection>
          <FooterWrapper>
            <WaveFooter />
          </FooterWrapper>
        </ContainerWithHeight>
      </IsGuest>
    </Fragment>
  );
}

// styled-components
const BlackDiv = styled.div`
  display: flex;
  flex-direction: column;
  position: relative;
  min-height: 140vh;
  width: 50%;
  padding: 60px 100px 0px 100px;
  background: #45424a;
`;

const SignUpInfo = styled.p`
  align-self: flex-start;
  display: grid;
  text-align: left;
  color: white;
  font-size: 50px;
`;

const Image = styled.img`
  align-self: flex-end;
  position: absolute;
  height: 550px;
  bottom: 0px;
  right: -120px;

  @media only screen and (max-width: 992px) {
    display: none;
  }
`;

const ContainerWithHeight = styled.div`
  position: relative;
  // padding: 0px 30px;
  margin-top: ${(props) => {
    return props.margintop;
  }};
  display: flex;
  flex-direction: column;
`;

const SeperateSection = styled.div`
  display: flex;
  width: 100%;
`;

const FormWrapper = styled.div`
  width: 50%;
`;

const FooterWrapper = styled.div`
    width: 100%;
    position: absolute;
    bottom: -70px;
`;