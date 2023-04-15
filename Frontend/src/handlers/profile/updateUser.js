import api from '../apiHandler';


export const updateStudent = (user)=>{
    return api.put(`/api/v1/student`, user)
}

export const updateTutor = (user)=>{
    return api.put(`/api/v1/tutor`, user)
}
