import FormT from './FormStyle.js';

export default function FormSignUp(){
    return(
        <form>
            <FormT.Div FormW='500px'>
                <FormT.Header>create account</FormT.Header>
                <FormT.Content>Sign in and start managing your candidates!</FormT.Content>
                <FormT.Content>
                    <FormT.Component>
                        <FormT.Label>First Name :</FormT.Label>
                        <FormT.TextInput BoxSize='140x' name='firstname' type='text' placeholder='Enter your first name'/>
                    </FormT.Component>
                    <FormT.Space></FormT.Space>
                    <FormT.Component>
                        <FormT.Label>Last Name :</FormT.Label>
                        <FormT.TextInput BoxSize='140px' name='lastname' type='text' placeholder='Enter your last name'/>
                    </FormT.Component>
                </FormT.Content>
                <FormT.Content>
                    <FormT.Component>
                        <FormT.Label>Username :</FormT.Label>
                        <FormT.TextInput BoxSize='315px' name='username' type='text' placeholder='Enter your username'/>
                    </FormT.Component>
                </FormT.Content>
                <FormT.Content>
                    <FormT.Component>
                        <FormT.Label>Email :</FormT.Label>
                        <FormT.TextInput BoxSize='315px' name='email' type='email' placeholder='Enter your email'/>
                    </FormT.Component>
                </FormT.Content>
                <FormT.Content>
                    <FormT.Component>
                        <FormT.Label>Password :</FormT.Label>
                        <FormT.TextInput BoxSize='140x' name='password' type='password' placeholder='Enter your password'/>
                    </FormT.Component>
                    <FormT.Space></FormT.Space>
                    <FormT.Component>
                        <FormT.Label>Confirm Password :</FormT.Label>
                        <FormT.TextInput BoxSize='140px' name='confirmpassword' type='password' placeholder='Enter your password again'/>
                    </FormT.Component>
                </FormT.Content>
                <FormT.Content>
                    <FormT.Component>
                        <FormT.Label>Address :</FormT.Label>
                        <FormT.TextInput BoxSize='315px' name='address' type='text' placeholder='Enter your address'/>
                    </FormT.Component>
                </FormT.Content>
                <FormT.Content>
                <FormT.Component>
                        <FormT.Label>Contact Number :</FormT.Label>
                        <FormT.TextInput BoxSize='315px' name='contactnumber' type='number' placeholder='Enter your contact number'/>
                    </FormT.Component>
                </FormT.Content>
                <FormT.Content>
                    <FormT.Component>
                        <FormT.Label>Gender :</FormT.Label>
                        <FormT.TextInput BoxSize='140x' name='password' type='password' placeholder='Enter your password'/>
                    </FormT.Component>
                    <FormT.Space></FormT.Space>
                    <FormT.Component>
                        <FormT.Label>Birth Date :</FormT.Label>
                        <FormT.TextInput BoxSize='140px' name='confirmpassword' type='password' placeholder='Enter your password again'/>
                    </FormT.Component>
                </FormT.Content>
                <FormT.Content>
                    <FormT.Component>
                        <FormT.Label>School :</FormT.Label>
                        <FormT.TextInput BoxSize='315px' name='contactnumber' type='number' placeholder='Enter your contact number'/>
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
