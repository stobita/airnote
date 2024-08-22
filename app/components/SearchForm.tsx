import React, { useState } from "react";
import styled from "styled-components";

interface Props {
  onSubmit: (word: string) => void;
}

export const SearchForm = (props: Props) => {
  const [word, setWord] = useState("");
  const handleOnSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    props.onSubmit(word);
  };

  const handleOnChangeWord = (e: React.ChangeEvent<HTMLInputElement>) => {
    setWord(e.target.value);
  };

  return (
    <Wrapper>
      <form onSubmit={handleOnSubmit}>
        <Input
          type="text"
          name="word"
          placeholder="Search word"
          onChange={handleOnChangeWord}
        ></Input>
      </form>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  display: flex;
  align-items: center;
  flex-basis: 50%;
`;

const Input = styled.input`
  background: ${props => props.theme.bg};
  color: ${props => props.theme.solid};
  border: 1px solid ${props => props.theme.border};
  padding: 8px;
  font-size: 1.2rem;
  width: 100%;
`;
