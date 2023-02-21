import { Fragment } from "react";
import styled from "styled-components";

export default function Footer(){


    return (
        <Fragment>
            <FooterSection>
                <FooterContent>
                <ImageFooter width='600px' src="/images/SignInFooter.png" alt="seprate70" />
                </FooterContent>
            </FooterSection>
        </Fragment>
    )
}


// styled-components for Navbar components
const FooterSection = styled.nav`
    margin-top: 3%;
`;

const FooterContent = styled.div`
    align-items: center;
    position:relative;
`;

const ImageFooter = styled.img`
    width: 100%;
    position: absolute;
`;