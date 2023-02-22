import api from './apiHandler';


export const  searchCourseHandler = (searchData)=>{
    return api.get('/api/v1/course', searchData)

}