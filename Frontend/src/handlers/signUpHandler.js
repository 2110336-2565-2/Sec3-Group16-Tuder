import api from './apiHandler';

export default async function signUpHandler(signUpData , navigate){
    api.post('/api/v1/signUp', signUpData)
    .then(function(response){
            
            let res = response.data;
            
            // do some response handling
            if( res.success === true){
                
                navigate('/sign-in')
            }
            

        }).catch(function(error){
            // if internal error occurs, MOO will return error message
            if (error.response) {
                let res = error.response.data;
                throw Error( res.message)
            }
        });
    };



