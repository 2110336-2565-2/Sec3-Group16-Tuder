import React from "react";
import styled from "styled-components"
import FormForgetPassword from '../components/form/FormForgetPassword.js'
import { IsGuest } from "../components/IsAuth.js";

export default function ForgetPassword(){

    return(
        <IsGuest>
            <ContainerWithHeight margintop='150px'>
                <FormForgetPassword />
            </ContainerWithHeight>
        </IsGuest>
    )
}

const ContainerWithHeight = styled.div`
    display: flex;
    padding: 0px 30px;
    margin-top: ${(props) => {
        return props.margintop
    }};
    justify-content: center;
`;
