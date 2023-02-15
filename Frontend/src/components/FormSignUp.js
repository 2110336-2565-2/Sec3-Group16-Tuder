import FormT from './FormStyle.js';
import signupContent from "../datas/SignUp.role.js";
import styled from 'styled-components';

export default function FormSignUp(props){
    const {role} = props

    let signupRole = null
    if(role === 'student'){
        signupRole = signupContent[0]
    }else{
        signupRole = signupContent[1]
    }
    const signupContents = signupRole.contents;
  
    const signupContentElement = signupContents.map((content, index) => {
        return (
            <FormT.Content key={index}>
                {content.map((element, elementindex) => {
                    const pholder = 'Enter your ' + element;
                    const name = element.replace(' ', '').toLowerCase()
                    let type = ''
                    let boxsize = ''
                    if(element === 'Email'){
                        type = 'email'
                        boxsize = '315px'
                    }else if(element === 'Password' | element === 'Confirm Password'){
                        type = 'password'
                        boxsize = '140px'
                    }else if(element === 'Birth Date'){
                        type = 'date'
                        boxsize = '140px'
                    }else if(element === 'Contact Number'){
                        type = 'number'
                        boxsize = '315px'
                    }else if(element === 'First Name' | element === 'Last Name'){
                        type = 'text'
                        boxsize = '140px'
                    }else if(element === 'Gender'){
                        type = 'text'
                        boxsize = '140px'
                    }
                    else{
                        type = 'text'
                        boxsize = '315px'
                    }
                    if(element === 'Birth Date'){
                        return (
                            <FormT.Component key={elementindex}>
                                <FormT.Label>{element} :</FormT.Label>
                                <FormT.DateInput BoxSize={boxsize} name={name} type={type} placeholder={pholder} />
                            </FormT.Component>
                        )
                    // need to solve this problem today
                    }else if(element === 'Gender'){
                        return(
                            <FormT.Component key={elementindex}>
                                <FormT.Label>{element} :</FormT.Label>
                                <FormT.TextInput BoxSize={boxsize} name={name} type={type} placeholder={pholder} />
                            </FormT.Component>
                        )
                    }else{
                        return (
                            <FormT.Component key={elementindex}>
                                <FormT.Label>{element} :</FormT.Label>
                                <FormT.TextInput BoxSize={boxsize} name={name} type={type} placeholder={pholder} />
                            </FormT.Component>
                        )
                    }
                })}
            </FormT.Content>
        )
    });
    return(
        <SignUpForm action='/signupcontroller' method='post'>
            <FormT.Div FormW='500px'>
                <FormT.Header>create account</FormT.Header>
                <FormT.Content>Sign in and start managing your candidates!</FormT.Content>
                {signupContentElement}
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