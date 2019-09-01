import React from "react";
import { Link } from "../model/link";
import { ReactComponent as DefaultImage } from "../assets/default.svg";
import styled from "styled-components";
import colors from "../colors";
import { TagBadge } from "./TagBadge";
import { useLinkItem } from "./LinkItem";

interface Props {
  item: Link;
  onClick: (l: Link) => void;
  onClickTag: (id: number) => void;
}

export const LinkItemPanel = (props: Props) => {
  const {
    item,
    handleOnClick,
    displayDescription,
    handleOnClickTag
  } = useLinkItem(props.item, true, props.onClick, props.onClickTag);

  return (
    <Wrapper onClick={handleOnClick}>
      <Default />
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

const Wrapper = styled.div`
  border: 1px solid ${colors.borderGray};
  border-radius: 4px;
  box-sizing: border-box;
  margin: 8px;
  cursor: pointer;
  overflow: hidden;
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

const Default = styled(DefaultImage)`
  height: auto;
`;
