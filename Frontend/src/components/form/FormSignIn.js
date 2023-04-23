import FormT from './FormStyle.js';
import React, { useState } from 'react';
import signInHandler from '../../handlers/signInHandler.js';
import { useNavigate, useOutletContext } from 'react-router-dom';
import { useDataContext } from '../../App.js';
import { toast } from 'react-hot-toast';


export default function FormSignIn(props){
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const  [status, setStatus] = useState('waiting');
    const  handleRole = useDataContext()[1];
    const navigate = useNavigate();

    async function submitHandler(e){
        e.preventDefault();
        if (status === 'submitting') return;
        setStatus('submitting')
        const loadingToast = toast.loading("Signing in...");
        const signInData = {
            username,
            password,
        }
        try{
            await signInHandler(signInData, navigate)

            // update role in state
            handleRole.handleRole();
            toast.success('Signed in')
            toast.dismiss(loadingToast);
            setStatus('success')
        } catch (error){
            // Handle by do sth
            toast.dismiss(loadingToast);
            toast.error("Invalid Username or Password");
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
                        <FormT.Link to='/forget-password' underline='none' color='red'>Forget password?</FormT.Link>
                    </FormT.ContentSmall>
                </FormT.ContentInline>
                <FormT.Content>
                    <FormT.Button type='submit'>Sign in</FormT.Button>
                </FormT.Content>
                <FormT.Content>
                    <FormT.ContentSmall>
                        Don't have an account ?
                    </FormT.ContentSmall>
                    <FormT.Link to='/sign-up' underline='none'>Sign Up</FormT.Link>
                </FormT.Content>
            </FormT.Div>
        </form>
    )
}

