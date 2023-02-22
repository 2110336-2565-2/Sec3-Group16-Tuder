import api from './apiHandler';

export default async function signUpHandler(signUpData , navigate){
    await api.post('/api/v1/signUp', signUpData)
    .then(function(response){
            
            let res = response.data;
            
            // do some response handling
            console.log(res);
            if( res.success === true){
                
                navigate('/signIn')
            }
            

        }).catch(function(error){
            // if internal error occurs, MOO will return error message
            if (error.response) {
                let res = error.response.data;
                throw Error( res.message)
            }
        });
    };



