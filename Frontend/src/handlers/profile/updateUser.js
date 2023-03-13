import api from '../apiHandler';


export const updateStudent = (user)=>{
    console.log("user: ", user)
    return api.put(`/api/v1/student`, user)
}

export const updateTutor = (user)=>{
    return api.put(`/api/v1/tutor`, user)
}
