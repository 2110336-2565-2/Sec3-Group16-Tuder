import FormT from './FormStyle.js';
import { useState } from 'react';
import signInAction from '../handlers/signInHandler.js';
import { useDispatch } from 'react-redux';
import { useSelector } from 'react-redux';


export default function FormSignIn(){
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    
    const dispatch = useDispatch();

    const role = useSelector(state => state.role);

    function submitHandler(e){
        e.preventDefault();
        const signInData = {
            username,
            password,
        }
        dispatch(signInAction(signInData));
        console.log(role)
    }

    return(
        <form onSubmit={submitHandler}>
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

