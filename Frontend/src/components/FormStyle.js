import styled from "styled-components";
import { Link } from "react-router-dom";

// styled-components

const FormT = {
    Div: styled.div`
        width: 350px;
        padding: 10px;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    `,
    Header: styled.div`
        display: flex;
        padding: 10px;
        font-size: 50px;
    `,
    Content: styled.div`
        display: flex;
        padding: 10px;
        font-size: 13px;
    `,
    ContentSmall: styled.div`
        display: flex;
        font-size: 11px;
        padding: 0px 10px;
        align-items: center;
    `,
    ContentInline: styled.div`
        display: flex;
        padding: 10px;
        font-size: 13px;
        justify-content: center;
    `,
    TextInput: styled.input`
        width: 200px;
        padding: 10px 13px;
        border: 1px solid black;
        border-radius: 6px;
        background-color: white;
        color: black;
        &::placeholder{
            color:black;
        }
    `,
    Checkbox: styled.input`
        margin-right: 5px;
        accent-color: #EB7B42;
    `,
    Button: styled.button`
        width: 225px;
        padding: 10px 13px;
        border: 0px solid;
        border-radius: 6px;
        color:white;
        background-color: #EB7B42;
        box-shadow: rgba(50, 50, 93, 0.25) 0px 6px 12px -2px, rgba(0, 0, 0, 0.3) 0px 3px 7px -3px;  
        &:hover{
            background-color: rgb(240, 123, 36);
        }
    `,
    Link: styled(Link)`
        padding: 0px 10px;
        text-decoration: ${(props)=>{
            return props.underline
        }};
        color: ${(props) => {
            return props.color
        }};
    `
}

export default FormT;