import React from "react";
import { ReactComponent as EditIconImage } from "../assets/edit.svg";
import { ReactComponent as DeleteIconImage } from "../assets/delete.svg";
import { TagBadge } from "./TagBadge";
import styled, { css } from "styled-components";
import { Link } from "../model/link";

interface Props {
  item: Link;
  isEdit: boolean;
  isDelete: boolean;
  onClickEdit: () => void;
  onClickDelete: () => void;
  onClickTag: (id: number) => void;
}

export const LinkDetailBody = (props: Props) => {
  const handleOnClickTag = (e: React.MouseEvent<HTMLElement>) => {
    const id = e.currentTarget.dataset.id;
    if (id) {
      props.onClickTag(Number(id));
    }
  };

  const displayUrl =
    props.item.url.length > 48
      ? `${props.item.url.slice(0, 36)}...`
      : props.item.url;

  return (
    <>
      <Row>
        <Anchor href={props.item.url} target="_blank" rel="noopener noreferrer">
          {displayUrl}
        </Anchor>
        {!props.isDelete && !props.isEdit && (
          <Operator>
            <EditIcon onClick={props.onClickEdit} />
            <DeleteIcon onClick={props.onClickDelete} />
          </Operator>
        )}
      </Row>
      <Title>{props.item.title}</Title>
      <Tags>
        {props.item.tags.map(v => (
          <TagBadge data-id={v.id} onClick={handleOnClickTag} key={v.id}>
            {v.text}
          </TagBadge>
        ))}
      </Tags>
      <Description>{props.item.description}</Description>
    </>
  );
};

const Row = styled.div`
  display: flex;
  align-items: center;
  padding: 16px 0;
`;

const Operator = styled.div`
  flex: 2;
  display: flex;
  justify-content: flex-end;
`;

const IconBase = css`
  fill: ${props => props.theme.text};
  height: 32px;
  padding: 8px;
`;

const EditIcon = styled(EditIconImage)`
  ${IconBase}
`;

const DeleteIcon = styled(DeleteIconImage)`
  ${IconBase}
`;

const Anchor = styled.a`
  color: ${props => props.theme.text};
  flex: 5;
`;

const Description = styled.p`
  white-space: pre-wrap;
`;

const Title = styled.p`
  margin: 0;
  font-size: 1.2rem;
  font-weight: bold;
  padding-bottom: 24px;
`;

const Tags = styled.div`
  padding: 8px;
`;
