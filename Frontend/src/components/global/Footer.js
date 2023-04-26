import React from "react";
import styled from "styled-components";
import { NavLink } from "react-router-dom";

// icon
import {
  EnvironmentFilled,
  PhoneFilled,
  PrinterFilled,
  FacebookFilled,
  TwitterCircleFilled,
  LinkedinFilled,
  YoutubeFilled,
  InstagramFilled,
  GooglePlusCircleFilled,
} from "@ant-design/icons";

export default function Footer() {
  return (
    <Container>
      <Wrapper>
        <HR borderWidth="1.3px" color="#EB7B42" />
        <MainSection>
          <AppLogo>Tuder</AppLogo>
          <InfoContainter>
            <InfoWrapper gridColumn="1 / 3">
              <EnvironmentFilled style={iconStyle} />
              <Address>
                345 Faulconer Drive, Suite 4 • Charlottesville, CA, 12345
              </Address>
            </InfoWrapper>
            <InfoWrapper gridRow="2 / 3">
              <PhoneFilled style={iconStyle} />
              <Info>(123) 456-7890</Info>
            </InfoWrapper>
            <InfoWrapper gridRow="2 / 3">
              <PrinterFilled style={iconStyle} />
              <Info>(123) 456-7890</Info>
            </InfoWrapper>
            <InfoWrapper marginTop="15px" gridColumn="1 / 3">
              <Topic>Social Media</Topic>
              <FacebookFilled style={iconStyle} />
              <TwitterCircleFilled style={iconStyle} />
              <LinkedinFilled style={iconStyle} />
              <YoutubeFilled style={iconStyle} />
              <InstagramFilled style={iconStyle} />
              <GooglePlusCircleFilled style={iconStyle} />
            </InfoWrapper>
          </InfoContainter>
        </MainSection>
        <HR borderWidth="0.5px" color="#E5F1FE" />
        <BottomSection>
        <Nav>
            <StyledNavLink to="/">ABOUT US</StyledNavLink>
            <StyledNavLink to="/">CONTACT US</StyledNavLink>
            <StyledNavLink to="/">HELP</StyledNavLink>
            <StyledNavLink to="/">PRIVACY POLICY</StyledNavLink>
            <StyledNavLink to="/report">REPORT PROBLEM</StyledNavLink>
        </Nav>
        <Info color="#848997">Copyright © 2023 • Tudor Inc.</Info>
      </BottomSection>
      </Wrapper>
    </Container>
  );
}

const Container = styled.footer`
  display: flex;
  background-color: white;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: 100%;
  font-weight: 200;
  font-size: 14px;
`;

const Wrapper = styled.div`
  width: 85%;
`;

const HR = styled.hr`
    width: 100%;
    border-width: ${(props) => {
        return props.borderWidth;
    }};
    color: ${(props) => {
        return props.color;
    }};
`;
const InfoContainter = styled.div`
  display: grid;
  row-gap: 15px;
`;

const MainSection = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 80px 15px;
  width: 100%;
`;

const AppLogo = styled.h1`
  color: #eb7b42;
  font-size: 40px;
  margin-left: 60px;
`;

const InfoWrapper = styled.div`
  display: flex;
  justify-content: flex-start;
  align-items: center;
  gap: 20px;
  grid-column: ${(props) => {
    return props.gridColumn;
  }};
  grid-row: ${(props) => {
    return props.gridRow;
  }};
  margin-top: ${(props) => {
    return props.marginTop;
  }};
  color: ${(props) => {
    return props.color;
    }};
`;

const Address = styled.address`
    font-style: normal;
`;

const Info = styled.p`
    color: ${(props) => {
        return props.color;
    }};
`;

const Topic = styled.h3`
  font-weight: 400;
  color: #848997;
`;

const iconStyle = {
  fontSize: "20px",
  color: "#EB7B42",
};

const BottomSection = styled.div`
    width: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 25px 0px 50px 0px;
`;

const Nav = styled.nav`
    display: flex;
    gap: 20px;
`;

const StyledNavLink = styled(NavLink)`
    color: #0A142F;
    text-decoration: none;
`;


