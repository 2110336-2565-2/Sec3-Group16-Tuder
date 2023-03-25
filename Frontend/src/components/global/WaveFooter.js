import React from "react";
import styled from "styled-components";

export default function WaveFooter({ backgroundColor, bottom }) {
  return (
    <Container backgroundColor={backgroundColor}>
      <Image
        src="/images/waveFooter.svg"
        alt="wave"
        backgroundColor={backgroundColor}
        bottom={bottom}
      />
    </Container>
  );
}

const Container = styled.div`
  width: 100%;
  background-color: ${(props) => props.backgroundColor};
`;

const Image = styled.img`
  width: 100%;
  position: relative;
  bottom: -4px;
  background-color: ${(props) => props.backgroundColor};
`;
