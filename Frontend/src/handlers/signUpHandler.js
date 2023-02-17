import api from './apiHandler';

export async function signUp(signUpData){
    return  await api.post('/api/v1/signUp', signUpData)
}


async function signUpAction(data, navigate) {

    
    
    // Input validation
    //************************ */
    //

    
    // Post to api
    return async (dispatch) => {

        await api.post('/api/v1/register', data)
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
                console.log(res.message);
                console.error(res.error)
            }
        });
    }
};

export default signUpAction;
