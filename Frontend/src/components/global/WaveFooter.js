import React from "react";
import styled from "styled-components";

export default function WaveFooter({ backgroundColor }) {
  return (
    <Container>
      <Image
        src="/images/waveFooter.svg"
        alt="wave"
        backgroundColor={backgroundColor}
      />
    </Container>
  );
}

const Container = styled.div`
  width: 100%;
  position: relative;
`;

const Image = styled.img`
  margin-top: 30px;
  width: 100%;
  position: absolute;
  background-color: ${(props) => props.backgroundColor || "white"};
`;