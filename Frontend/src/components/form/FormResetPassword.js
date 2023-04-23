import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom";
import { toast } from 'react-hot-toast';

import FormT from './FormStyle.js';
import styled from 'styled-components';
import jwtDecode from 'jwt-decode';
import { resetPasswordHandler } from '../../handlers/resetPasswordHandler.js';

export default function FormResetPassword(){

    const navigate = useNavigate();
    
    const [password, setPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');
    const token = window.location.href.split("/").pop();
    // if token is not valid, navigate to 404 page
    useEffect(() => {
        try{
            let decode = jwtDecode(token);
            if (decode.exp < Date.now() / 1000) {
                navigate("/404");
            }
        } catch (error){
            navigate("/404");
        }
    }, [token, navigate])


    async function submitHandler(e){
        e.preventDefault();
        const data = {
            token: token,
            password: password,
            confirm_password: confirmPassword,
        }
        
        await resetPasswordHandler(data).then((res) => {
            if (res.data.success) {
                toast.success('Reset Password Successfully');
                navigate('/sign-in');
            }
        }).catch((error) => {
            toast.error("Something went wrong");
        })
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
                    <FormT.TextInput BoxSize='500px' name='email' type='password' placeholder='' value={password} onChange={(e)=>setPassword(e.target.value)}/>
                </TextInput>
                <TextInput>
                    <FormT.Content>Confirm Password</FormT.Content>
                    <FormT.TextInput BoxSize='500px' name='email' type='password' placeholder='' value={confirmPassword} onChange={(e)=>setConfirmPassword(e.target.value)}/>
                </TextInput>
                <FormT.Content>
                    <FormT.Button type='submit'>Confirm</FormT.Button>
                </FormT.Content>
                <FormT.Content>
                    <CancelButton type='cancel' onClick={()=>navigate("/sign-in")}>Sign In</CancelButton>
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
