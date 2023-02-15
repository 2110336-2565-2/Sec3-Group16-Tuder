import FormT from './FormStyle.js';
import { useState } from 'react';
import signUpHandler from '../handlers/signUpHandler.js';

export default function FormSignUp(){
    const [firstname, setFirstname] = useState('');
    const [lastname, setLastname] = useState('');
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const  [confirmpassword, setConfirmPassword] = useState('');
    const [address, setAddress] = useState('');
    const [contactnumber, setContactNumber] = useState('');
    const [gender, setGender] = useState('');
    const [birthdate, setBirthdate] = useState('');
    const [school, setSchool] = useState('');


    return(
        <form onSubmit={signUpHandler}>
            <FormT.Div FormW='500px'>
                <FormT.Header>create account</FormT.Header>
                <FormT.Content>Sign in and start managing your candidates!</FormT.Content>
                <FormT.Content>
                    <FormT.Component>
                        <FormT.Label>First Name :</FormT.Label>
                        <FormT.TextInput BoxSize='140x' name='firstname' type='text' placeholder='Enter your first name' value={firstname} onChange={(e) => setFirstname(e.target.value)}/>
                    </FormT.Component>
                    <FormT.Space></FormT.Space>
                    <FormT.Component>
                        <FormT.Label>Last Name :</FormT.Label>
                        <FormT.TextInput BoxSize='140px' name='lastname' type='text' placeholder='Enter your last name' value={lastname} onChange={(e) => setLastname(e.target.value)}/>
                    </FormT.Component>
                </FormT.Content>
                <FormT.Content>
                    <FormT.Component>
                        <FormT.Label>Username :</FormT.Label>
                        <FormT.TextInput BoxSize='315px' name='username' type='text' placeholder='Enter your username' value={username} onChange={(e) => setUsername(e.target.value)}/>
                    </FormT.Component>
                </FormT.Content>
                <FormT.Content>
                    <FormT.Component>
                        <FormT.Label>Email :</FormT.Label>
                        <FormT.TextInput BoxSize='315px' name='email' type='email' placeholder='Enter your email' value={email} onChange={(e) => setEmail(e.target.value)}/>
                    </FormT.Component>
                </FormT.Content>
                <FormT.Content>
                    <FormT.Component>
                        <FormT.Label>Password :</FormT.Label>
                        <FormT.TextInput BoxSize='140x' name='password' type='password' placeholder='Enter your password' value={password} onChange={(e) => setPassword(e.target.value)}/>
                    </FormT.Component>
                    <FormT.Space></FormT.Space>
                    <FormT.Component>
                        <FormT.Label>Confirm Password :</FormT.Label>
                        <FormT.TextInput BoxSize='140px' name='confirmpassword' type='password' placeholder='Enter your password again' value={confirmpassword} onChange={(e) => setConfirmPassword(e.target.value)}/>
                    </FormT.Component>
                </FormT.Content>
                <FormT.Content>
                    <FormT.Component>
                        <FormT.Label>Address :</FormT.Label>
                        <FormT.TextInput BoxSize='315px' name='address' type='text' placeholder='Enter your address' value={address} onChange={(e) => setAddress(e.target.value)}/>
                    </FormT.Component>
                </FormT.Content>
                <FormT.Content>
                <FormT.Component>
                        <FormT.Label>Contact Number :</FormT.Label>
                        <FormT.TextInput BoxSize='315px' name='contactnumber' type='number' placeholder='Enter your contact number' value={contactnumber} onChange={(e) => setContactNumber(e.target.value)}/>
                    </FormT.Component>
                </FormT.Content>
                <FormT.Content>
                    <FormT.Component>
                        <FormT.Label>Gender :</FormT.Label>
                        <FormT.TextInput BoxSize='140x' name='gender' type='text' placeholder='Enter your gender'  value={gender} onChange={(e) => setGender(e.target.value)}/>
                    </FormT.Component>
                    <FormT.Space></FormT.Space>
                    <FormT.Component>
                        <FormT.Label>Birth Date :</FormT.Label>
                        <FormT.TextInput BoxSize='140px' name='birthdate' type='date' placeholder='Enter your birthdate' value={birthdate} onChange={(e) => setBirthdate(e.target.value)}/>
                    </FormT.Component>
                </FormT.Content>
                <FormT.Content>
                    <FormT.Component>
                        <FormT.Label>School :</FormT.Label>
                        <FormT.TextInput BoxSize='315px' name='school' type='text' placeholder='Enter your school' value={school} onChange={(e) => setSchool(e.target.value)}/>
                    </FormT.Component>
                </FormT.Content>
                <FormT.Content>
                    <FormT.Button type='submit'>Sign in</FormT.Button>
                </FormT.Content>
                <FormT.Content>
                    <FormT.ContentSmall>
                        Already have an account ?
                    </FormT.ContentSmall>
                    <FormT.Link to='/Signin' underline='none'>Sign In</FormT.Link>
                </FormT.Content>
            </FormT.Div>
        </form>
        
    )
}
