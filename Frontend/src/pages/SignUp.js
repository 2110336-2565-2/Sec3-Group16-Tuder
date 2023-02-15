import styled from "styled-components"
import FormSignUp from "../components/FormSignUp.js";

export default function SignUp(){

    return(
        <ContainerWithHeight margintop='100px'>
            <SeperateSection>
                <ItemGrid justify='center' columngrid='1 / 3'>
                    <ImageTudor width='600px' src="/images/signin.png" alt="seprate70" />
                </ItemGrid>
                <ItemGrid columngrid='4'>
                    <FormSignUp />
                </ItemGrid>
            </SeperateSection>
        </ContainerWithHeight>
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