import styled from "styled-components"
import FormSignUp from "../components/FormSignUp.js";

export default function SignUp(props){
    const {role} = props
    return(
        <ContainerWithHeight margintop='auto'>
            <FlexSection>
                <FormSignUp role={role} />
            </FlexSection>
        </ContainerWithHeight>
    )
}


// styled-components

const ContainerWithHeight = styled.div`
    padding: 0px 30px;
    margin-top: ${(props) => {
        return props.margintop
    }};
`;

const FlexSection = styled.div`
    display: flex;
    justify-content: flex-end;
`;