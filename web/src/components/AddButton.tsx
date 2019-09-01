import React from "react";
import styled from "styled-components";

interface Props {
  onClick: () => void;
}

export const AddButton = (props: Props) => {
  return (
    <Button onClick={props.onClick}>
      <ButtonInner />
      <ButtonInner />
    </Button>
  );
};

const Button = styled.button`
  border: 1px solid ${props => props.theme.border};
  height: 64px;
  width: 64px;
  border-radius: 32px;
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
`;

const ButtonInner = styled.span`
  position: absolute;
  width: 50%;
  height: 2px;
  background: ${props => props.theme.text};
  border-radius: 4px;
  &:nth-of-type(1) {
    transform: rotate(90deg);
  }
`;
