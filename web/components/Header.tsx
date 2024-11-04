import React from "react";
import styled from "styled-components";
import { AddButton } from "./AddButton";
import { SearchForm } from "./SearchForm";

interface Props {
  onClickAddButton: () => void;
  onSubmitWordSearch: (word: string) => void;
}

export const Header = (props: Props) => {
  return (
    <Wrapper>
      <SearchForm onSubmit={props.onSubmitWordSearch}></SearchForm>
      <AddButton onClick={props.onClickAddButton} />
    </Wrapper>
  );
};

const Wrapper = styled.div`
  position: fixed;
  width: calc(100% - 240px);
  background: ${props => props.theme.main};
  color: ${props => props.theme.solid};
  border-bottom: 1px solid ${props => props.theme.border};
  display: flex;
  justify-content: space-between;
  padding: 8px 24px;
  box-sizing: border-box;
`;
