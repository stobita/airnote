import React from "react";
import { Link } from "../model/link";
import styled from "styled-components";
import colors from "../colors";
import { TagBadge } from "./TagBadge";

interface Props {
  item: Link;
  onClick: (l: Link) => void;
  onClickTag: (id: number) => void;
}

export const LinkItem = (props: Props) => {
  const {
    item,
    handleOnClick,
    displayDescription,
    handleOnClickTag
  } = useLinkItem(props.item, false, props.onClick, props.onClickTag);

  return (
    <Wrapper onClick={handleOnClick}>
      <Bottom>
        <Title>{item.title && item.title}</Title>
        <Description>{displayDescription}</Description>
        <Tags>
          {item.tags.map(v => (
            <TagBadge data-id={v.id} onClick={handleOnClickTag} key={v.id}>
              {v.text}
            </TagBadge>
          ))}
        </Tags>
      </Bottom>
    </Wrapper>
  );
};

export const useLinkItem = (
  item: Link,
  isBlock: boolean,
  onClick: (l: Link) => void,
  onClickTag: (id: number) => void
) => {
  const handleOnClick = () => {
    onClick(item);
  };

  const handleOnClickTag = (e: React.MouseEvent<HTMLElement>) => {
    e.stopPropagation();
    const id = e.currentTarget.dataset.id;
    if (id) {
      onClickTag(Number(id));
    }
  };

  const { description } = item;
  const displayDescriptionLength = isBlock ? 30 : 60;

  const displayDescription =
    description.length > displayDescriptionLength - 1
      ? `${description.slice(0, displayDescriptionLength)}...`
      : description.length > 0
      ? description
      : "no description";
  return { item, handleOnClick, displayDescription, handleOnClickTag };
};

const Wrapper = styled.div`
background: ${props => props.theme.bg}
  border: 1px solid ${colors.borderGray};
  border-radius: 4px;
  box-sizing: border-box;
  margin: 8px;
  cursor: pointer;
  overflow: hidden;
  width: 100%;
`;

const Description = styled.p`
  word-break: break-word;
`;

const Bottom = styled.div`
  padding: 8px;
`;

const Tags = styled.div`
  display: flex;
`;

const Title = styled.p`
  font-weight: bold;
`;
