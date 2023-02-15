import FormT from './FormStyle.js';
import { useState } from 'react';
import loginHandler from '../controllers/signInController.js';

export default function FormSignIn(){
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    
    return(
        <form onSubmit={loginHandler}>
            <FormT.Div FormW='350px'>
                <FormT.Header>Sign In</FormT.Header>
                <FormT.Content>Sign in and start managing your candidates!</FormT.Content>
                <FormT.Content>
                    <FormT.TextInput BoxSize='200px' name='username' type='text' placeholder='Username'/>
                </FormT.Content>
                <FormT.Content>
                    <FormT.TextInput BoxSize='200px' name='password' type='password' placeholder='Password'/>
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

