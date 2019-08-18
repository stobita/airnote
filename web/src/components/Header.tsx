import React from "react";
import styled from "styled-components";
import { AddButton } from "./AddButton";
import colors from "../colors";

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
  display: flex;
  justify-content: flex-end;
  padding: 16px 24px;
  background: ${colors.mainGray};
`;
