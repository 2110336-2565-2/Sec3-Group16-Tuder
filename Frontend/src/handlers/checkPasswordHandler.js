import api from './apiHandler';

export default async function checkPasswordHandler(Data, navigate){
    await api.post('/api/v1/changepassword-check', Data).then((response) => {
        
        let res = response.data;

        if (res.success === true) {
            
            console.log(res.success); 
            navigate('/enter-new-password')
            
        }

    }).catch(function(error){
        // if internal error occurs, MOO will return error message
        if (error.response) {
            let res = error.response.data;
            console.log("error Wrong password");
            throw new Error(res.message);
        }
    });

}