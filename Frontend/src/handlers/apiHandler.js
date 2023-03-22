import axios from 'axios';

const api = axios.create({
    baseURL: process.env.REACT_APP_API_URL,
    headers: {
        post: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + localStorage.getItem('jwtToken')
        },
        get: {
        'Authorization': 'Bearer ' + localStorage.getItem('jwtToken')
        },
        put: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer ' + localStorage.getItem('jwtToken')
          },
        delete: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer ' + localStorage.getItem('jwtToken')
          },

    }
  });

export    default api;
