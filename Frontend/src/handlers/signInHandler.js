import api from './apiHandler';
import setRole from '../features/role/roleSlice'

export function signIn(signInData){
    return  api.post('/api/v1/login', signInData)
}

export default function signInAction(data){
    return (dispatch) => {
        
        signIn(data).then((response) => {
        
            let res = response.data;
            
            // if login success, MOO will set token
            if (res.success === true) {

                let token = res.data.token;
                localStorage.setItem('jwtToken', token); 
                console.log(res);

                // redux state
                let role = res.data.role
                console.log("user role: "+ role)
                dispatch(
                    setRole(role)
                )
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
