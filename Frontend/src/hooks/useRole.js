import { useState } from 'react';
import getRole from '../utils/jwtGet';


export default function useRole(){
    const [role, setRole] = useState(getRole());

    function handleRole(){
        setRole(getRole());
    }

    return [role, handleRole];

}