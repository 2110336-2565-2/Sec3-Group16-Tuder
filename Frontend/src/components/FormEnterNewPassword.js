import React, { useState } from 'react';
import { useNavigate } from "react-router-dom";
import changePasswordHandler from '../handlers/changePasswordHandler.js';
import useUsername from '../hooks/useUsername.js';

import FormT from './FormStyle.js';
import styled from 'styled-components';

export default function FormEnterNewPassword(){
    const [username, handleUsername] = useUsername();
    const [password, setPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');

    const navigate = useNavigate();

    async function submitHandler(e){
        e.preventDefault();
        handleUsername();
        const changePasswordData = {
            username,
            password,
            confirmPassword,
        }
        try{
            await changePasswordHandler(changePasswordData, navigate);
        } catch (error){
            console.log(error);
        }
    }


    return(
        <form onSubmit={submitHandler} action='' method='post'>
            <FormFormat FormW='700px' >
                <Header>
                    <FormT.Header>Enter New password</FormT.Header>
                    <FormT.Content>Password must be at least 8 characters long.</FormT.Content>
                    <FormT.Content>Password must contain at least one upper case.</FormT.Content>
                    <FormT.Content>One lower case letter.</FormT.Content>
                    <FormT.Content>Password must contain at least one number or special character</FormT.Content>
                </Header>
                <TextInput>
                    <FormT.Content>New Password</FormT.Content>
                    <FormT.TextInput BoxSize='500px' name='email' type='text' placeholder='' value={password} onChange={(e)=>setPassword(e.target.value)}/>
                </TextInput>
                <TextInput>
                    <FormT.Content>Confirm Password</FormT.Content>
                    <FormT.TextInput BoxSize='500px' name='email' type='text' placeholder='' value={confirmPassword} onChange={(e)=>setConfirmPassword(e.target.value)}/>
                </TextInput>
                <FormT.Content>
                    <FormT.Button type='submit'>Send</FormT.Button>
                </FormT.Content>
                <FormT.Content>
                    <CancelButton type='cancel' onClick={()=>navigate("/profile")}>Cancel</CancelButton>
                </FormT.Content>
            </FormFormat>
        </form>
        
    )
}

const FormFormat = styled(FormT.Div)`
    display: flex;
    flex-direction: column;
    padding: 150px 10px 150px 10px;
    gap: 10px;
    align-items: center;
`

const CancelButton = styled(FormT.Button)`
    background-color: White;
    border: 2px solid Orange;
    color: Orange;
    &:hover {
        background-color: Orange;
        color: White;
    }

`

const TextInput = styled(FormT.Content)`
    display: flex;
    flex-direction: column;
`

const Header = styled(FormT.Content)`
    display: left;
    flex-direction: column;
    & > ${FormT.Content}::before {
        content: "•";
        margin-right: 10px;

      }
`

