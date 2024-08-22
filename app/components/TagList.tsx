import React, { useContext } from "react";
import { DataContext } from "../context/dataContext";
import tagsRepository from "../api/tagsRepository";
import styled from "styled-components";

export const TagList = () => {
  const { tags, setLinks } = useContext(DataContext);
  const onClickTag = (id: number) => {
    tagsRepository.getLinks(id).then(links => {
      setLinks(links);
    });
  };
  const handleOnClickTag = (e: React.MouseEvent<HTMLLIElement>) => {
    const id = e.currentTarget.dataset.id;
    if (id) {
      onClickTag(Number(id));
    }
  };
  return (
    <ul>
      {tags.map(tag => (
        <Item key={tag.id} data-id={tag.id} onClick={handleOnClickTag}>
          {tag.text}
        </Item>
      ))}
    </ul>
  );
};

const Item = styled.li`
  cursor: pointer;
  font-size: 1.1rem;
  margin: 8px 0;
`;
