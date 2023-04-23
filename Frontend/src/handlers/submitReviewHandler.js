import api from './apiHandler';
import { toast } from 'react-hot-toast';

export default async function submitReviewHandler(submitData){
    await api.post('/api/v1/review', 
        submitData,
        {headers:{'Content-Type': 'application/json','Authorization': 'Bearer ' + localStorage.getItem('jwtToken')}} 
    )
    .then(function(response){
        let res = response.data;
        if( res.success === true){
            // do some response handling
            console.log(res);
        }
    }
    ).catch(function(error){
        // if internal error occurs, MOO will return error message
        if (error.response) {
            let res = error.response.data;
            console.log(res.message);
            throw Error( res.message)
        }
    });
};

export const fetchCourseHandler = (id) => {
    return api.get('api/v1/course/'+id)
}
export const fetchTutorHandler = (id) => {
    return api.get('api/v1/tutorID/'+ id)
}
