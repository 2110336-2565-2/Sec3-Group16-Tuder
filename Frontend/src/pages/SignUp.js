import React, { Fragment } from "react";
import styled from "styled-components"
import WaveFooter from "../components/global/WaveFooter.js";
import { IsGuest } from "../components/IsAuth.js";
import FormSignUp from "../components/form/FormSignUp.js";

export default function SignUp(){

    return(
        <Fragment>
            <IsGuest>
            <ContainerWithHeight margintop='0px'>
                <SeperateSection>
                    <ItemGrid justify='center' columngrid='1 / 2'>
                        <BlackDiv>
                            <SignUpInfo>
                            Getting <br/>
                            Started With <br/>
                            Matching system
                            </SignUpInfo>
                            <Image src="/images/signupDecorator.svg" alt="decorator"/>
                        </BlackDiv>               
                    </ItemGrid>
                    <ItemGrid columngrid='3 / 4'>
                        <FormSignUp />
                    </ItemGrid>
                </SeperateSection>
            </ContainerWithHeight>
            <WaveFooter />
            </IsGuest>
        </Fragment>
    )
}


// styled-components
const BlackDiv = styled.div`
    position: relative;
    width: 100%;
    height: 110%;
    padding: 60px 100px 0px 100px;
    background: #45424A;

`;

const SignUpInfo = styled.p`
    display: grid;
    text-align: left;
    color: white;
    font-size: 50px;
`;

const Image = styled.img`
    position: absolute;
    height: 65%;
    bottom: 1%;
    left: 30%;
`;

const ContainerWithHeight = styled.div`
    // padding: 0px 30px;
    margin-top: ${(props) => {
        return props.margintop
    }};

`;

const SeperateSection = styled.div`
    display: grid;
    width: 100%;
`;

const ItemGrid = styled.div`
    width: 100%;
    grid-column: ${(props) => {
        return props.columngrid
    }};
    justify-self: ${(props) => {
        return props.justify
    }}
`;