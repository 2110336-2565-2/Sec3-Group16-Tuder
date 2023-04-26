import React from 'react'
import styled from 'styled-components'
import { Link } from 'react-router-dom';
export default function Home(){
    return (
        <>
            <HomeContainer>
                <HomeHeader>
                    <HomeTitle>Welcome to Tuder</HomeTitle>
                    <HomeSubtitle>The ultimate platform for online learning</HomeSubtitle>
                </HomeHeader>
                <Container>
                    <HomeDescription>
                    <HomeFeature>
                        <HomeFeatureTitle>Learn at Your Own Place</HomeFeatureTitle>
                        <HomeFeatureDescription>Tuder provides a flexible learning experience, allowing you to learn at your own place, anywhere.</HomeFeatureDescription>
                    </HomeFeature>
                    <HomeFeature>
                        <HomeFeatureTitle>Access to High-Quality Content</HomeFeatureTitle>
                        <HomeFeatureDescription>Our platform offers a vast library of courses, carefully curated by experts in various fields.</HomeFeatureDescription>
                    </HomeFeature>
                    <HomeFeature>
                        <HomeFeatureTitle>Interactive Learning Tools</HomeFeatureTitle>
                        <HomeFeatureDescription>Engage in interactive learning with discussions, designed to enhance your understanding and easy to use.</HomeFeatureDescription>
                    </HomeFeature>
                    </HomeDescription>
                </Container>
            </HomeContainer>
        </>
    )
}

const Container = styled.div`
    margin-top: 7%;
    margin-left: 4%;
    margin-right: 4%;
    margin-bottom: 4%;
    border-radius: 10px;
    padding: 3%;
`

const HomeContainer = styled.div`
    margin: 0 auto;
`

const HomeHeader = styled.div`
    text-align: center;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    margin-bottom: 60px;
    height: 500px;
    background-image: linear-gradient(rgba(0, 0, 0, 0.60), rgba(0, 0, 0, 0.60)), url("/images/backgroundtudor.jpg");
    background-repeat: no-repeat;
    background-attachment: fixed;
    background-position: bottom;
    background-size: cover;
    
`

const HomeTitle = styled.div`
    font-size: 45px;
    font-weight: bold;
    margin: 0;
    color: #FF702F;
`
const HomeSubtitle = styled.div`
    font-size: 18px;
    color: #EEEEEE;
    margin-top: 10px;
`

const HomeDescription = styled.div`
    display: flex;
    justify-content: center;
    gap: 40px;
`

const HomeFeature = styled.div`
    flex-basis: 33.33%;
    text-align: center;
`

const HomeFeatureTitle = styled.h2`
    font-size: 24px;
    font-weight: bold;
    margin-top: 20px;
`

const HomeFeatureDescription = styled.div`
    font-size: 16px;
    color: #666;
    margin-top: 10px;
`
const GapLink = styled.div`
    padding: 3px;
`
const FlexHome = styled.div`
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
`

const FlexHeader = styled.div`
    display: flex;
    font-size: 40px;
`

const FlexContent = styled.div`
    display: flex;
`

const FlexButtonRow = styled.div`
    display: flex;
    flex-direction: row;
`
const FlexButtonContent = styled.div`
    display: flex;
`

const StartButton = styled.button`
  width: 100px;
  height: 50px;
  border: 2px solid #FF7008;
  border-radius: 5px;

  font-family: 'Lexend Deca';
  font-style: normal;
  font-weight: 500;
  font-size: 16px;
  line-height: 20px;
  text-align: center;
  color: #FFFFFF;
  background-color: #FF7008;

  &:hover {
    background-color: #FF8009;
    border: 2px solid #FF8008;
    color: #FFFFFF;
  }

  cursor: pointer;
`;
