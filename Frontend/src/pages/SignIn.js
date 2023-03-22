import React, { Fragment } from "react";
import styled from "styled-components"
import FormSignIn from "../components/form/FormSignIn.js";
import Footer from "../components/footer/SignInFooter.js";
import { IsGuest } from "../components/IsAuth.js";

export default function SignIn(){
    

    return(
        <Fragment>
            <IsGuest>

            <ContainerWithHeight margintop='100px'>
                <SeperateSection>
                    <ItemGrid justify='center' columngrid='1 / 3'>
                        <ImageTudor width='600px' src="/images/signin.png" alt="seprate70" />
                    </ItemGrid>
                    <ItemGrid columngrid='4'>
                        <FormSignIn />
                    </ItemGrid>
                </SeperateSection>
            </ContainerWithHeight>
            </IsGuest>
            <Footer/>
        </Fragment>
    )
}


// styled-components
const ImageTudor = styled.img`
    max-width: ${(props) => {
        return props.maxwidth
    }};
`;

const ContainerWithHeight = styled.div`
    padding: 0px 30px;
    margin-top: ${(props) => {
        return props.margintop
    }};
`;

const SeperateSection = styled.div`
    display: grid;
`;

const ItemGrid = styled.div`
    grid-column: ${(props) => {
        return props.columngrid
    }};
    justify-self: ${(props) => {
        return props.justify
    }}
`;