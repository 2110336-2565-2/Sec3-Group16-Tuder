import React, {useState, useEffect} from 'react';
import FormT from './FormStyle.js';
import styled from 'styled-components';
import { toast } from 'react-hot-toast';
import {forgetPasswordHandler} from '../../handlers/forgetPasswordHandler.js';
import { useNavigate } from 'react-router-dom';

export default function FormSignIn(){

    const [email, setEmail] = React.useState('');
    const [isLoading, setIsLoading] = useState(false);
    const navigate = useNavigate();



    const handleSubmit = async(e) => {
        e.preventDefault();
        toast.promise(
            forgetPasswordHandler({email}),
            {
              loading: 'Sending email',
              success: 'Email has been sent!',
              error: (error) => error.response.data.message,
            },
            {
              // Optional options object for toast configuration
            }
          );
        
    }



    return(
            <FormT.Div FormW='400px'>
                <FormT.Header>Forget Password</FormT.Header>
                <FormT.Content>Sign in and start managing your candidates!</FormT.Content>
                <FormT.Content>
                    <FormT.TextInput BoxSize='200px' value={email} onChange={(e)=> setEmail(e.target.value)} type='email' placeholder='Enter your Email'/>
                </FormT.Content>
    
                <FormT.Content>
                    <FormT.Button onClick={handleSubmit}>Send</FormT.Button>
                </FormT.Content>
                <FormT.Content>
                    <ContentSpace />
                    <FormT.Link to='/sign-in' underline='none'>back to login</FormT.Link>
                </FormT.Content>

                {isLoading && toast.loading('Submitting data...')}
            </FormT.Div>
        
    )
}

const ContentSpace = styled.div`
    display: flex;
    font-size: 11px;
    padding: 0px 10px;
    align-items: center;
    margin-right: 110px;
`