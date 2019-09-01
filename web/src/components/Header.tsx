import React from "react";
import styled from "styled-components";
import { AddButton } from "./AddButton";

interface Props {
  onClickAddButton: () => void;
}

export const Header = (props: Props) => {
  return (
    <Wrapper>
      <AddButton onClick={props.onClickAddButton} />
    </Wrapper>
  );
};

const Wrapper = styled.div`
  background: ${props => props.theme.bg};
  color: ${props => props.theme.text};
  border-bottom: 1px solid ${props => props.theme.border};
  display: flex;
  justify-content: flex-end;
  padding: 16px 24px;
`;
