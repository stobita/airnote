import React from "react";
import styled from "styled-components";

interface Props {
  onClick: () => void;
}

export const CloseButton = (props: Props) => {
  return (
    <Button onClick={props.onClick}>
      <ButtonInner />
      <ButtonInner />
    </Button>
  );
};
const Button = styled.button`
  height: 64px;
  width: 64px;
  border-radius: 32px;
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
`;

const ButtonInner = styled.span`
  background: ${props => props.theme.text};
  position: absolute;
  width: 50%;
  height: 2px;
  border-radius: 4px;
  &:nth-of-type(1) {
    transform: rotate(45deg);
  }
  &:nth-of-type(2) {
    transform: rotate(135deg);
  }
`;
