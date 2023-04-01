import api from './apiHandler';

export default async function signInHandler(reportData){

        await api.post('/api/v1/createissuereport', reportData).then((response) => {
            
            let res = response.data;
            // if login success, MOO will set token
            if (res.success === true) {
                console.log("Create issue report successfuly")
            }
        
        }).catch(function(error){
            // if internal error occurs, MOO will return error message
            if (error) {
                let res = error.response.data;
                throw new Error(res.message);
            }
        });
        
}