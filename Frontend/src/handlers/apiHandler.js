import axios from 'axios';
import { useEffect } from 'react';

const api = axios.create({
    baseURL: process.env.REACT_APP_API_URL 
  });

export    default api;


// const updateRole = (trigger,data) => {
//   useEffect(() => {
//     if 



//   })

// }
