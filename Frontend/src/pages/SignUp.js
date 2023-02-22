import { Fragment } from "react";
import styled from "styled-components"
import FormSignUp from "../components/FormSignUp.js";
import Footer from "../components/SignUpFooter.js";

export default function SignUp(){

    return(
        <Fragment>

            <ContainerWithHeight margintop='0px'>
                <SeperateSection>
                    <ItemGrid justify='center' columngrid='1 / 2'>
                        <BlackDiv>
                            <SignUpInfo>
                            Getting <br/>
                            Started With <br/>
                            Matching system
                            </SignUpInfo>
                            
                        </BlackDiv>               
                    </ItemGrid>
                    <ItemGrid columngrid='3 / 4'>
                        <FormSignUp />
                    </ItemGrid>
                </SeperateSection>
            </ContainerWithHeight>
            <Footer/>
        </Fragment>
    )
}


// styled-components
const BlackDiv = styled.div`
position: relative;
width: 800px;
height: 950px;

background: #45424A;

`;

const SignUpInfo = styled.div`
    display: grid;
    text-align: left;
    color: white;
    padding-top: 60px;
    padding-left: 100px;
    font-size: 50px;
`;

const ImageTudor = styled.img`
    max-width: ${(props) => {
        return props.maxwidth
    }};
    margin-top: ${(props)=>{
        return props.mt
    }};
`;

const ContainerWithHeight = styled.div`
    // padding: 0px 30px;
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