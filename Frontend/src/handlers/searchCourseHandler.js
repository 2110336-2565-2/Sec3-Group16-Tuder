import api from './apiHandler';


export const  searchCourseHandler = (searchData)=>{
    return api.post('/api/v1/coursesearch', 
        searchData
    )

}

export const fetchCourseHandler = () => {
    return api.get('api/v1/courses')
}