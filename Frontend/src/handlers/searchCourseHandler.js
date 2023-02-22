import api from '../api';


export default async function searchCourseHandler(searchData, navigate){
    await api.post('/api/v1/course', searchData).then((response) => {
        
        let res = response.data;

        // if login success, MOO will set token

        if (res.success === true) {
            
            // change state and put jwt token on a local storage
            
            navigate('/')
            
        }

    }).catch(function(error){
        // if internal error occurs, MOO will return error message
        if (error.response) {
            let res = error.response.data;
            // console.log(res.message);
            throw new Error(res.message);
        }
    });

}