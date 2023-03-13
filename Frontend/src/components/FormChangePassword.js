import React, { useState } from 'react';
import { useNavigate } from "react-router-dom";
import checkPasswordHandler from '../handlers/checkPasswordHandler.js';

import FormT from './FormStyle.js';
import styled from 'styled-components';

export default function FormChangePassword(){
    const [password, setPassword] = useState('');

    const navigate = useNavigate();

    async function submitHandler(e){
        e.preventDefault();
        const checkPasswordData = {
            password,
        }
        try{
            await checkPasswordHandler(checkPasswordData, navigate);
        } catch (error){
            console.log(error);
        }
    }

    return(
        <form onSubmit = {submitHandler} action='' method='post'>
            <FormFormat FormW='700px' >
                <FormT.Header>Enter your password</FormT.Header>
                
                <FormT.Content>
                    <FormT.TextInput BoxSize='500px' name='email' type='text' placeholder='Enter your Password' value={password} onChange={(e)=>setPassword(e.target.value)}/>
                </FormT.Content>
    
                <FormT.Content>
                    <FormT.Button type='submit' onClick={()=>navigate("/enter-new-password")}>Send</FormT.Button>
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