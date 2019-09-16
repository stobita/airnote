import React, { useContext } from "react";
import { Tag } from "../model/link";
import styled from "styled-components";
import { TagBadge } from "./TagBadge";
import tagsRepository from "../api/tagsRepository";
import { DataContext } from "../context/dataContext";

interface Props {
  items: Tag[];
}
export const LinkItemTags = (props: Props) => {
  const { setLinks } = useContext(DataContext);
  const handleOnClickTag = (e: React.MouseEvent<HTMLElement>) => {
    e.stopPropagation();
    const id = e.currentTarget.dataset.id;
    if (id) {
      tagsRepository.getLinks(Number(id)).then(links => {
        setLinks(links);
      });
    }
  };
  return (
    <div>
      <Wrapper>
        {props.items.map(v => (
          <TagBadge data-id={v.id} onClick={handleOnClickTag} key={v.id}>
            {v.text}
          </TagBadge>
        ))}
      </Wrapper>
    </div>
  );
};

const Wrapper = styled.div`
  display: flex;
`;
