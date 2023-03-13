import api from '../apiHandler';


export const getStudentByUsername = (username)=>{
    return api.get(`/api/v1/student/${username}`)
}

export const getTutorByUsername = (username)=>{
    return api.get(`/api/v1/tutor/${username}`)
}
