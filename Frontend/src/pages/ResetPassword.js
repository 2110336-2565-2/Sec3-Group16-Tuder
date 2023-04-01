import React from "react";
import styled from "styled-components"
import FormResetPassword from '../components/form/FormResetPassword.js'
import WaveFooter from "../components/global/WaveFooter.js";


export default function EnterNewPassword(){

    return(
            <Container>
                <ContainerWithHeight margintop='100px'>
                    <FormResetPassword />
                </ContainerWithHeight>
                <WaveFooterWrapper>
                    <WaveFooter />
                </WaveFooterWrapper>
            </Container>          
    )
}

const ContainerWithHeight = styled.div`
    display: flex;
    flex-direction: column
    padding: 0px 30px;
    margin-top: ${(props) => {
        return props.margintop
    }};
    justify-content: center;
`;

const WaveFooterWrapper = styled.div`
    position: absolute;
    bottom: 0;
    width: 100%;
`;

const Container = styled.div`
    position: relative;
    min-height: 100vh;
    width: 100%;
`;
