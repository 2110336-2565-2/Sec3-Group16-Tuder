import FormT from './FormStyle.js';
import React, { useState } from 'react';
import toast from 'react-hot-toast';
import signInHandler from '../handlers/signInHandler.js';
import { useNavigate, useOutletContext } from 'react-router-dom';


export default function FormSignIn(props){
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [status, setStatus] = useState('waiting');
    const [role, handleRole] = useOutletContext();
    
    const navigate = useNavigate();


    async function submitHandler(e){
        e.preventDefault();
        const signInData = {
            username,
            password,
        }
        setStatus('submitting')
        try{
            await signInHandler(signInData, navigate);

            // update role in state
            handleRole.handleRole();
            setStatus('success')
        } catch (error){

            // Handle by do sth
            toast.error(error.message);
            console.log(error);
            setStatus('error');
        }
    }

    return(
        <form onSubmit={submitHandler} >
            <FormT.Div FormW='350px'>
                <FormT.Header>Sign In</FormT.Header>
                <FormT.Content>Sign In and start managing your candidates!</FormT.Content>
                <FormT.Content>
                    <FormT.TextInput BoxSize='200px' name='username' type='text' placeholder='Username' value={username} onChange={(e) => setUsername(e.target.value)}/>
                </FormT.Content>
                <FormT.Content>
                    <FormT.TextInput  BoxSize='200px' name='password' type='password' placeholder='Password' value={password} onChange={(e) => setPassword(e.target.value)}/>
                </FormT.Content>
                <FormT.ContentInline>
                    <FormT.ContentSmall>
                        <FormT.Checkbox type='checkbox' />
                        <div>
                            Remember me
                        </div>
                    </FormT.ContentSmall>
                    <FormT.ContentSmall>
                        <FormT.Link to='/Forgetpassword' underline='none' color='red'>Forget password?</FormT.Link>
                    </FormT.ContentSmall>
                </FormT.ContentInline>
                <FormT.Content>
                    <FormT.Button type='submit'>Sign in</FormT.Button>
                </FormT.Content>
                <FormT.Content>
                    <FormT.ContentSmall>
                        Don't have an account ?
                    </FormT.ContentSmall>
                    <FormT.Link to='/Signup' underline='none'>Sign Up</FormT.Link>
                </FormT.Content>
            </FormT.Div>
        </form>
    )
}

