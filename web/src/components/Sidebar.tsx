import React from "react";
import styled from "styled-components";
import { Tag } from "../model/link";

interface Props {
  tags: Tag[];
  onClickTitle: () => void;
  onClickTag: (id: number) => void;
}

export const Sidebar = (props: Props) => {
  const handleOnClickTag = (e: React.MouseEvent<HTMLLIElement>) => {
    const id = e.currentTarget.dataset.id;
    if (id) {
      props.onClickTag(Number(id));
    }
  };
  return (
    <Wrapper>
      <Title onClick={props.onClickTitle}>AirNote</Title>
      <ul>
        {props.tags.map(tag => (
          <Item key={tag.id} data-id={tag.id} onClick={handleOnClickTag}>
            {tag.text}
          </Item>
        ))}
      </ul>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  background: ${props => props.theme.bg};
  color: ${props => props.theme.text};
  border-right: 1px solid ${props => props.theme.border};
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  height: 100%;
  padding: 24px;
  box-sizing: border-box;
`;

const Title = styled.h2`
  cursor: pointer;
`;

const Item = styled.li`
  cursor: pointer;
  font-size: 1.1rem;
  margin: 8px 0;
`;
