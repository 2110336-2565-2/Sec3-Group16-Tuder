import jwt_decode from 'jwt-decode';

export default function getUsername(){
    const token = localStorage.getItem('jwtToken');
    if(token){
        const decoded = jwt_decode(token);
        return decoded.username;
    }else{
        return 'guest';
    }
}