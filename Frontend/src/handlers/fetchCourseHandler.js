import api from './apiHandler';

export const fetchCourseHandler = () => {
    return api.get('/courses')
}


export const fetchCoursebyJSONHandler = (searchData) => {
    return api.post('/course', searchData).then((response) => {
        
        let res = response.data;

        // if login success, MOO will set token

        if (res.success === true) {
            
            // change state and put jwt token on a local storage
            
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