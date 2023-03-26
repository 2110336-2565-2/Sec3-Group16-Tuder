import styled from "styled-components";

const Button = styled.button`
  background-color: ${(props) =>
    props.variance === "cancel" ? "#FFFFFF" : "#FF7008"};
  color: ${(props) => (props.variance === "cancel" ? "#FF7008" : "#FFFFFF")};
  font-size: 13px;
  width: 100px;
  padding: 10px;
  border-radius: 10px;
  border: 2px solid #ff7008;

  cursor: pointer;
`;

export default Button;
