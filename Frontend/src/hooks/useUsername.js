import { useState } from 'react';
import getUsername from '../utils/jwtGetUsername';


export default function useUsername(){
    const [username, setUsername] = useState(getUsername());

    function handleUsername(){
        setUsername(getUsername());
    }

    return [username, handleUsername];

}