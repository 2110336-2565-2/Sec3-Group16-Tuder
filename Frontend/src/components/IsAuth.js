import React from "react";
import useRole  from "../hooks/useRole";
import {Navigate}  from "react-router-dom";

export function IsTutor({children}) {
    const [role, handleRole] = useRole();

    if(role === "tutor"){
        return children;
    }else{
        return <Navigate to="/"/>
    }
}

export function IsStudent({children}) {
    const [role, handleRole] = useRole();

    if(role === "student"){
        return children;
    }else{
        return <Navigate to="/"/>
    }
}

export function IsAdmin({children}) {
    const [role, handleRole] = useRole();

    if(role === "admin"){
        return children;
    }else{
        return <Navigate to="/"/>
    }
}

export function IsGuest({children}) {
    const [role, handleRole] = useRole();

    if(role === "guest"){
        return children;
    }else{
        return <Navigate to="/"/>
    }
}

export function IsUser({children}) {
    const [role, handleRole] = useRole();

    if(role === "student" || role === "tutor"){
        return children;
    }else{
        return <Navigate to="/"/>
    }
}

export function AuthVerify({children}) {
    
}





