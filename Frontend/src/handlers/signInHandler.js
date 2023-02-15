import api from './apiHandler';
import {setRole} from '../features/role/roleSlice'

export async function signIn(signInData){
    return  await api.post('/api/v1/login', signInData)
}

export default async function signInAction(data, navigate){
    return async (dispatch) => {

        await signIn(data).then((response) => {
        
            let res = response.data;
            
            // if login success, MOO will set token
            if (res.success === true) {

                let token = res.data.token;
                localStorage.setItem('jwtToken', token); 
                console.log(res);

                // redux state
                let role = res.data.role
              
                dispatch(
                    setRole(role)
                )

                navigate('/')
                
            }

        }).catch(function(error){
            // if internal error occurs, MOO will return error message
            if (error.response) {
                let res = error.response.data;
                console.log(res.message);
            }
        });

    }
}
