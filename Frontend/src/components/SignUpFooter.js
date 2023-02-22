import { Fragment } from "react";
import styled from "styled-components";

export default function Footer(){


    return (
        <Fragment>
            <FooterSection>
                <FooterContent>
                <ImageFooter width='600px' src="/images/SignUpFooter.png" alt="seprate70" />
                </FooterContent>
            </FooterSection>
        </Fragment>
    )
}


// styled-components for Navbar components
const FooterSection = styled.nav`
    margin-top: -50%;
    pointer-events: none;
`;

const FooterContent = styled.div`
    align-items: center;
    position:relative;
`;

const ImageFooter = styled.img`
    width: 100%;
    position: absolute;
`;