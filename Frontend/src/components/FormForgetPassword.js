import React from 'react';
import FormT from './FormStyle.js';
import styled from 'styled-components';

export default function FormSignIn(){

    return(
        <form action='' method='post'>
            <FormT.Div FormW='450px'>
                <FormT.Header>Reset Password</FormT.Header>
                <FormT.Content>Sign in and start managing your candidates!</FormT.Content>
                <FormT.Content>
                    <FormT.TextInput BoxSize='200px' name='username' type='text' placeholder='Enter your Email'/>
                </FormT.Content>
    
                <FormT.Content>
                    <FormT.Button type='submit'>Send</FormT.Button>
                </FormT.Content>
                <FormT.Content>
                    <ContentSpace />
                    <FormT.Link to='/sign-in' underline='none'>back to login</FormT.Link>
                </FormT.Content>
            </FormT.Div>
        </form>
    )
}

const ContentSpace = styled.div`
    display: flex;
    font-size: 11px;
    padding: 0px 10px;
    align-items: center;
    margin-right: 110px;
`