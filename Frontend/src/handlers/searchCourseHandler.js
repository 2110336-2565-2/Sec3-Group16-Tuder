import api from './apiHandler';


export const  searchCourseHandler = (searchData)=>{
    return api.post('/api/v1/coursesearch', 
        searchData
    )

}