import React from "react";
import colors from "../colors";
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
  background: ${colors.mainWhite};
  border: 1px solid ${colors.borderGray};
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
  background: ${colors.mainGray};
  border-radius: 4px;
  &:nth-of-type(1) {
    transform: rotate(90deg);
  }
`;
