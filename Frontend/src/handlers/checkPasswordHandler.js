import api from './apiHandler';

export default async function checkPasswordHandler(Data, navigate){
    await api.post('/api/v1/changepassword-check', Data, {
        headers: {
          "Content-Type": "application/json",
          Authorization: "Bearer " + localStorage.getItem("jwtToken"),
        }
      })
    
    .then((response) => {
        
        let res = response.data;

        if (res.success === true) {
            navigate('/enter-new-password')
        }

    }).catch(function(error){
        // if internal error occurs, MOO will return error message
        if (error.response) {
            let res = error.response.data;
            throw new Error(res.message);
        }
    });

}