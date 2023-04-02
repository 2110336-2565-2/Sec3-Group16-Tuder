import jwt_decode from 'jwt-decode';
import useRole from '../hooks/useRole';
import useUsername from '../hooks/useUsername';

export function getRole(){
    const token = localStorage.getItem('jwtToken');
    if(token){
        const decoded = jwt_decode(token);
        return decoded.role;
    }else{
        return 'guest';
    }
}

export function getUsername(){
    const token = localStorage.getItem('jwtToken');
    if(token){
        const decoded = jwt_decode(token);
        return decoded.username;
    }else{
        return '';
    }
}

export function getUserID(){
    const token = localStorage.getItem('jwtToken');
    if(token){
        const decoded = jwt_decode(token);
        return decoded.userid;
    }else{
        return '';
    }
}

export function verify(){
    const token = localStorage.getItem('jwtToken');
    if(token){
        const decoded = jwt_decode(token);
        const currentTime = Date.now() / 1000;
        if(decoded.exp < currentTime){
            localStorage.removeItem('jwtToken');
            window.location.href = '/';
            return false;
        }
        return true;
    }
}

export function getUserId(){
    const token = localStorage.getItem('jwtToken');
    if(token){
        const decoded = jwt_decode(token);
        return decoded.userid;
    }else{
        return '';
    }
}
