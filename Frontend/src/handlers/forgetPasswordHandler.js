import api from './apiHandler';

export  const forgetPasswordHandler = (Data) => {
    return api.post('/api/v1/forget-password', Data)
}

