import React, { useCallback } from "react";
import { Link } from "../model/link";
import { ReactComponent as DefaultImage } from "../assets/default.svg";
import styled from "styled-components";
import colors from "../colors";
import { TagBadge } from "./TagBadge";

interface Props {
  item: Link;
  onClick: (l: Link) => void;
  onClickTag: (id: number) => void;
}

export const LinkItem = (props: Props) => {
  const handleOnClick = useCallback(() => {
    props.onClick(props.item);
  }, [props]);

  const handleOnClickTag = (e: React.MouseEvent<HTMLElement>) => {
    e.stopPropagation();
    const id = e.currentTarget.dataset.id;
    if (id) {
      props.onClickTag(Number(id));
    }
  };

  const { description } = props.item;

  const displayDescription =
    description.length > 30
      ? `${description.slice(0, 30)}...`
      : description.length > 0
      ? description
      : "no description";

  return (
    <Wrapper onClick={handleOnClick}>
      <Default />
      <Bottom>
        <Title>{props.item.title && props.item.title}</Title>
        <Description>{displayDescription}</Description>
        <Tags>
          {props.item.tags.map(v => (
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
