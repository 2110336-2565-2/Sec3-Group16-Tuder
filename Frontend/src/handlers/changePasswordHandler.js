import api from './apiHandler';

export default async function changePasswordHandler(Data, navigate){
    await api.post('/api/v1/changepassword', Data).then((response) => {
        
        let res = response.data;

        if (res.success === true) {
            navigate('/profile')
        }

    }).catch(function(error){
        // if internal error occurs, MOO will return error message
        if (error.response) {
            let res = error.response.data;
            console.log(res);
            throw new Error(res.message);
        }
    });

}
