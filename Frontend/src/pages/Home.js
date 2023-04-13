import React from 'react'
import styled from 'styled-components'
import Footer from "../components/global/Footer.js";

export default function Home(){
    return (
        <Container>
            <FlexHome>
                <FlexHeader>Learn desired course</FlexHeader>
                <FlexHeader>From your desired</FlexHeader>
                <FlexHeader>tutor</FlexHeader>
                <FlexContent>
                Are you looking to enhance your knowledge, acquire new skills, or boost your career prospects? 
                Look no further! Our study course is designed to provide you with the knowledge, resources, and support you need to achieve your goals.
                </FlexContent>
                <FlexButtonRow>
                    <FlexButtonContent>
                        <StartButton>Start Learning</StartButton>
                    </FlexButtonContent>
                    <FlexButtonContent>
                        <StartButton></StartButton>
                    </FlexButtonContent>
                </FlexButtonRow>
            </FlexHome>
        </Container>
    )
}

const Container = styled.div`
    margin-top: 4%;
    margin-left: 7%;
    margin-right: 7%;
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
