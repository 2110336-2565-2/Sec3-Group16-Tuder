import api from './apiHandler';

export const resetPasswordHandler = (Data) => {
    return api.post('/api/v1/reset-password', Data)
}

