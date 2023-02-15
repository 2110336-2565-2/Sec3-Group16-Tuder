import FormT from './FormStyle.js';
import { useState } from 'react';

import signInController from "../controllers/signInController"


export default function FormSignIn(){
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    
const handleSubmit = (event) => {
    event.preventDefault(); // Prevent the default form submit behavior
    signInController(username, password); // Call signInController with username and password
}

    return(
        <form onSubmit={handleSubmit}>
            <FormT.Div FormW='350px'>
                <FormT.Header>Sign In</FormT.Header>
                <FormT.Content>Sign in and start managing your candidates!</FormT.Content>
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

