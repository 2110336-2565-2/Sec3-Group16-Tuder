import FormT from './FormStyle.js';
import { useState } from 'react';
import signUpHandler from '../handlers/signUpHandler.js';
import signupContent from "../datas/SignUp.role.js";
import styled from 'styled-components';

import { useNavigate } from 'react-router-dom';



export default function FormSignUp(){
    // Input State
    const [firstname, setFirstName] = useState('');
    const [lastname, setLastName] = useState('');
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const  [confirmpassword, setConfirmPassword] = useState('');
    const [address, setAddress] = useState('');
    const [contactnumber, setContactNumber] = useState('');
    const [gender, setGender] = useState('');
    const [birthdate, setBirthDate] = useState('');
    const [role, setRole] = useState('');


    const valueSetter = {"First Name":setFirstName, "Last Name":setLastName,
     "Username":setUsername, "Email": setEmail, "Password": setPassword,
      "Confirm Password": setConfirmPassword, "Address" : setAddress,
       "Contact Number" : setContactNumber, "Gender":setGender,
        "Birth Date":setBirthDate, "Role":setRole}

    const  setStatus = useState('waiting')[1];
 
    const navigate = useNavigate()



    async function submitHandler(e){
        e.preventDefault();
        let birthdateISO = (new Date(birthdate)).toISOString(); 
        const signUpData = {
            username : username,
            password: password,
            email: email,
            confirm_password: confirmpassword,
            firstname: firstname,
            lastname: lastname,
            address: address,
            phone:  contactnumber,
            birthdate: birthdateISO,
            gender: gender,
            role: role
         }
         setStatus('submitting')
         try{
            await signUpHandler(signUpData, navigate);
            setStatus('success')
        } catch (error){

            // Handle by do sth

            console.log(error);
            setStatus('error');
        }
    }

    
    const signupContents = signupContent.contents;
    const signupContentElement = signupContents.map((content, index) => {
        
        return (
        
            <FormT.Content key={index}>
                { 
                content.map((element, elementindex) => {
                    const pholder = 'Enter your ' + element;
                    const name = element.replace(' ', '').toLowerCase()
                    let type = ''
                    let boxsize = ''
                    let value = ''
                    let onChange = (e) => valueSetter[element](e.target.value);

                    if(element === 'Username'){
                        type = 'text'
                        boxsize = '315px'
                        value = username
                    } else if(element === 'Email'){
                        type = 'email'
                        boxsize = '315px'
                        value = email
                    }else if(element === 'Password') {
                        type = 'password'
                        boxsize = '140px'
                        value = password
                    }
                    else if (element === 'Confirm Password'){
                        type = 'password'
                        boxsize = '140px'
                        value = confirmpassword
                    }else if(element === 'Birth Date'){
                        type = 'date'
                        boxsize = '140px'
                        value = birthdate
                    }else if(element === 'Contact Number'){
                        type = 'number'
                        boxsize = '315px'
                        value = contactnumber
                    }else if(element === 'First Name'){
                        type = 'text'
                        boxsize = '140px'
                        value = firstname
                        
                    }else if (element === 'Last Name'){
                        type = 'text'
                        boxsize = '140px'
                        value = lastname
                        
                    }else if(element === 'Gender'){
                        type = 'text'
                        boxsize = '140px'
                        value = gender
                    }
                    else if (element === 'Address'){
                        type = 'text'
                        boxsize = '315px'
                        value = address
                    }
                    
                    if(element === 'Birth Date'){
                        return (
                            <FormT.Component key={elementindex}>
                                <FormT.Label>{element} :</FormT.Label>
                                <FormT.DateInput BoxSize={boxsize} name={name} type={type} placeholder={pholder} value={value} onChange={onChange} required/>
                            </FormT.Component>
                        )
                    }else if(element === 'Gender'){
                        return(
                            <FormT.Component key={elementindex}>
                                <FormT.Label>{element} :</FormT.Label>
                                <FormT.Select BoxSize='170px' name={name} value={value} onChange={onChange}  required >
                                    <FormT.Option value="" disabled hidden>Please select</FormT.Option>
                                    <FormT.Option value='male'>male</FormT.Option>
                                    <FormT.Option value='female'>female</FormT.Option>
                                    <FormT.Option value='-'>I don't like hee</FormT.Option>
                                </FormT.Select>
                            </FormT.Component>
                        )
                    }else{
                        return (
                            <FormT.Component key={elementindex}>
                                <FormT.Label>{element} :</FormT.Label>
                                <FormT.TextInput BoxSize={boxsize} name={name} type={type} placeholder={pholder} value={value} onChange={onChange} required/>
                            </FormT.Component>
                        )
                    }
                })}
            </FormT.Content>
        )
    });
    return(
        <SignUpForm onSubmit={submitHandler}>
            <FormT.Div FormW='500px'>
                <FormT.Header>create account</FormT.Header>
                <FormT.Content>Sign in and start managing your candidates!</FormT.Content>
                {signupContentElement}
                <FormT.Content>
                    <FormT.Component>
                        <FormT.Label>As :</FormT.Label>
                        <FormT.Select BoxSize='343px' name='role' value={role} onChange={(e) => setRole(e.target.value)}  required >
                            <FormT.Option value="" disabled hidden>Please select</FormT.Option>
                            <FormT.Option value='student'>Student</FormT.Option>
                            <FormT.Option value='tutor'>Tutor</FormT.Option>
                        </FormT.Select>
                    </FormT.Component>
                </FormT.Content>
                <FormT.Content>
                    <FormT.Button type='submit'>Sign Up</FormT.Button>
                </FormT.Content>
                <FormT.Content>
                    <FormT.ContentSmall>
                        Already have an account ?
                    </FormT.ContentSmall>
                    <FormT.Link to='/Signin' underline='none'>Sign In</FormT.Link>
                </FormT.Content>
            </FormT.Div>
        </SignUpForm>
        
    )
              
}

const SignUpForm = styled.form`
    display: flex;
    justify-content: center;
`;

