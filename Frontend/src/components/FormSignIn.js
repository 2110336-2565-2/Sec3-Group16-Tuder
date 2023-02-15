import FormT from './FormStyle.js';
import { useState } from 'react';
import axios from 'axios';

const api = axios.create({
    baseURL: 'http://localhost:8080' // replace with the URL of your locally running server
  });

export default function FormSignIn(){
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    async function handleSubmit(event) {

        event.preventDefault();
       
            //console.log({username})
            const response = await api.post('/api/v1/login', { username:"hee", password:"brightHee"})
            .then(function(response){
                console.log(response)
                const token = response.data.data.token;
               
                localStorage.setItem('jwtToken', token); 
                const tokencheck = localStorage.getItem('jwtToken');
                console.log(tokencheck);
            // Redirect to the home page
            // window.location.href = '/';
            }).catch(function(error){
                console.error(error);
            });
           
           

        
    };
    
    return(
        <form >
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
                    <FormT.Button onClick= {handleSubmit} type='submit'>Sign in</FormT.Button>
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

