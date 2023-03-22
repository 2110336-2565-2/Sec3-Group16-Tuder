import api from './apiHandler';


// export const  getCancellingClass = (searchData)=>{
//     return api.post('/api/v1/c', 
//         searchData
//     )

// }

export const fetchCancellingClassHandler = () => {
    return api.get('api/v1/cancelling-classes')
}

export const submitApprovmentHandler = (classId) => {
    return api.post('api/v1/audit-class-cancellation', {
        classId: classId,
        approve: true
    })
}

export const submitRejectionHandler = (classId) => {
    return api.post('api/v1/audit-class-cancellation', {
        classId: classId,
        approve: false
    })
}