import styled from "styled-components"
import FormForgetPassword from '../components/FormForgetPassword.js'

export default function ForgetPassword(){

    return(
        <ContainerWithHeight margintop='100px'>
            <FormForgetPassword />
        </ContainerWithHeight>
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
