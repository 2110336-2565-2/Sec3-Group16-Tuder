import FormT from './FormStyle.js';

export default function FormSignUp(){
    return(
        <form>
            <FormT.Div>
                <FormT.Header>create account</FormT.Header>
                <FormT.Content>Sign in and start managing your candidates!</FormT.Content>
                <FormT.Content>
                    <FormT.TextInput name='username' type='text' placeholder='Enter your Username'/>
                </FormT.Content>
                <FormT.Content>
                    <FormT.TextInput name='email' type='email' placeholder='Enter your E-mail'/>
                </FormT.Content>
                <FormT.Content>
                    <FormT.TextInput name='phonenumber' type='tel' pattern="[0-9]{3}-[0-9]{3}-[0-9]{4}" placeholder='Enter your Phone Number'/>
                </FormT.Content>
                <FormT.Content>
                    <FormT.TextInput name='password' type='password' placeholder='Enter your Password'/>
                </FormT.Content>
                <FormT.Content>
                    <FormT.TextInput name='repassword' type='password' placeholder='Retry Enter your Password'/>
                </FormT.Content>
                <FormT.Content>
                    <FormT.Button type='submit'>Sign Up</FormT.Button>
                </FormT.Content>
                <FormT.Content>
                    <FormT.ContentSmall>
                        Already have an account ?
                    </FormT.ContentSmall>
                    <FormT.Link to='/SignIn' underline='none'>Login</FormT.Link>
                </FormT.Content>
            </FormT.Div>
        </form>
        
    )
}
