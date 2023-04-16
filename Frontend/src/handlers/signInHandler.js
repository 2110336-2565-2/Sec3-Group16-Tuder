import api from './apiHandler';

export default async function signInHandler(signInData, navigate){


        await api.post('/api/v1/login', signInData).then((response) => {
            
            let res = response.data;
            // if login success, MOO will set token
            if (res.success === true) {
                
                
                
                // change state and put jwt token on a local storage
            let token = res.data.token;
            localStorage.setItem('jwtToken', token); 
            window.dispatchEvent(new Event("storage"));
            
            navigate('/')
            
            }
        
        }).catch(function(error){
            // if internal error occurs, MOO will return error message
            if (error) {
                let res = error.response.data;
                throw new Error(res.message);
            }
        });
        
}
