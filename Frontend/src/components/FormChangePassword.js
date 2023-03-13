import React, { useState } from 'react';
import { useNavigate } from "react-router-dom";
import { toast } from 'react-hot-toast';
import checkPasswordHandler from '../handlers/checkPasswordHandler.js';
import useUsername from '../hooks/useUsername.js';

import FormT from './FormStyle.js';
import styled from 'styled-components';

export default function FormChangePassword(){
    const [username, handleUsername] = useUsername();
    const [password, setPassword] = useState('');

    const navigate = useNavigate();

    async function submitHandler(e){
        e.preventDefault();
        handleUsername();
        
        const checkPasswordData = {
            username,
            password,
        }
        try{
            await checkPasswordHandler(checkPasswordData, navigate);
            toast.success('Correct Password');
        } catch (error){
            console.log(error);
            toast.error('Wrong Password');
        }
    }

    return(
        <form onSubmit = {submitHandler} action='' method='post'>
            <FormFormat FormW='700px' >
                <FormT.Header>Enter your password</FormT.Header>
                
                <FormT.Content>
                    <FormT.TextInput BoxSize='500px' name='email' type='password' placeholder='Enter your Password' value={password} onChange={(e)=>setPassword(e.target.value)}/>
                </FormT.Content>
    
                <FormT.Content>
                    <FormT.Button type='submit'>Send</FormT.Button>
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
    border: 1px solid black;
    border-radius: 10px;
`