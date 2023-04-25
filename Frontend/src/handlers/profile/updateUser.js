import api from '../apiHandler';


export const updateStudent = (user)=>{
    return api.put(`/api/v1/student`, user,{
        headers:
        {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('jwtToken')
        }
    })
}

export const updateTutor = (user)=>{
    return api.put(`/api/v1/tutor`, user,{
        headers:
        {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('jwtToken')
        }
    })
    
}
