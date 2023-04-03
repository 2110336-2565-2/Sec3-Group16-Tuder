import React from 'react'
import styled from 'styled-components'
import Footer from "../components/global/Footer.js";

export default function Home(){
    return (
        <>
            <h1>This is homepage</h1>
            <Footer />
        </>
    )
}


const Header = styled.div`
    display: flex;
`