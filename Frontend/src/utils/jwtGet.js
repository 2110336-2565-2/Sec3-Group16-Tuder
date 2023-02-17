import { useState, useEffect } from 'react';
import jwt_decode from 'jwt-decode';

export default function getRole(){
    const token = localStorage.getItem('jwtToken');
    if(token){
        const decoded = jwt_decode(token);
        return decoded.role;
    }else{
        return 'guest';
    }
}