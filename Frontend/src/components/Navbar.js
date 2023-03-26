import { Outlet, Link } from 'react-router-dom';
import navbarContent from '../datas/Navbar.role.js';
import styled from 'styled-components';
import {signOutAction} from '../handlers/signOutHandler';
import {useNavigate} from 'react-router-dom';
import React, { Fragment, useEffect } from 'react';
import useRole from '../hooks/useRole';
import { useDataContext } from '../App.js';
import { toast } from 'react-hot-toast';



export default function Navbar(){
    // choose Navbar contents array from role 
    
    var [role, handleRole] = useDataContext();
    role = role.role;
    handleRole = handleRole.handleRole;
    const navigate = useNavigate();

    window.addEventListener('storage', () => {
        handleRole();
    })

    

    
    
    function signOutHandler(e) {
        e.preventDefault();
        signOutAction(navigate);
        handleRole();
        toast.success('Signed Out', {style: {     
            color: '#713200',
          },
          iconTheme: {
            primary: '#DC143C',
            
          },
        });
    }

    let navbarRole = null
    if(role === 'admin'){
        navbarRole = navbarContent[3]
    }else if(role === 'student'){
        navbarRole = navbarContent[2]
    }else if(role === 'tutor'){
        navbarRole = navbarContent[1]
    }else{
        navbarRole = navbarContent[0]
    }
    
    // raw data contents JSON
    const contents = navbarRole.content;
    // change to component for use in JSX  --> Generate NavItem for each content
    const contentElement = contents.map((content, index) => {
        if(content.title === 'Home'){
            return <NavbarItem key="home"><TuderLinkNav to={content.link} key={index}>{content.title}</TuderLinkNav></NavbarItem>
        }else if(content.title === 'Sign Up'){
            return <NavbarItem key="signUp"><TuderButton type='tudor-button' to={content.link} key={index}>{content.title}</TuderButton></NavbarItem>
        }else if(content.title === 'Sign In'){
            return <NavbarItem key="signIn"><TuderLinkNav to={content.link} key={index}>{content.title}</TuderLinkNav></NavbarItem>
        }else if(content.title === 'Sign Out'){
            return <NavbarItem key="signOut"><TuderButton type='red-button' onClick={signOutHandler} key={index}>{content.title}</TuderButton></NavbarItem>
        }else{
            const urlLink = (content.link)?content.link:("/" + content.title.toLowerCase().replace(/ /g, "-"));
            return <NavbarItem key={content.title}><TuderLinkNav to= {urlLink} key={index}>{content.title}</TuderLinkNav></NavbarItem>
        }
    });

    return (
        <Fragment>
            <NavbarSection>
                <NavbarHeader>Tuder</NavbarHeader>
                <NavbarItems>
                    {contentElement}
                </NavbarItems>
            </NavbarSection>   
            <Outlet context={[{role},{handleRole}]} />   
        </Fragment>
    )
}


// styled-components for Navbar components
const NavbarSection = styled.nav`
    height: 70px;
    width: 100%;
    display: flex;
    padding: 10px 30px;
    box-shadow: 0px 2.98px 7.45px rgba(0, 0, 0, 0.1);
    justify-content: space-between;
    align-items: center;
`;

const NavbarHeader = styled.div`
    color:rgb(227, 105, 12); 
    font-size: 30px;
    font-weight: 700  ;
`;

const NavbarItems = styled.ul`
    list-style: none;
    display: flex;
`;

const NavbarItem = styled.li`
    padding-left: 20px;
    padding-right: 20px;
    font-weight: 350;
    font-size: 15px;
`;

const TuderButton = styled(Link)`
    text-decoration: none;
    background-color: ${(props) => {
        if(props.type === 'red-button'){
            return 'rgb(255, 0, 0)'
        }else if(props.type === 'tudor-button'){
            return '#EB7B42'
        }
        return '#EB7B42'
    }};
    padding: 8px 13px;
    border: 0px solid;
    border-radius: 6px;
    color: white;
    &:hover{
        background-color: ${(props) => {
        if(props.type === 'red-button'){
            return 'rgb(255, 68, 68)'
        }else if(props.type === 'tudor-button'){
            return 'rgb(240, 123, 36)'
        }
        return 'rgb(240, 123, 36)'
    }};
    }
`;

const TuderLinkNav = styled(Link)`
    text-decoration: none;
    color: black;
    &:hover{
        color: #EB7B42;
    }
`;