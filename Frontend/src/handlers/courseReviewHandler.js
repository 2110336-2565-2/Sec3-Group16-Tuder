import api from './apiHandler';


export const getCourseReviewHandler = (courseID)=>{
    return api.get('/api/v1/review/' + courseID 
    )

}