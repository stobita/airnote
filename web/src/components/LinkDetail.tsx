import React, { useCallback, useState } from "react";
import { Link, Tag } from "../model/link";
import { ReactComponent as DefaultImage } from "../assets/default.svg";
import styled from "styled-components";
import { DeleteConfirmation } from "./DeleteConfirmation";
import linksRepository from "../api/linksRepository";
import { LinkDetailBody } from "./LinkDetailBody";

interface Props {
  item: Link;
  tags: Tag[];
  afterUpdate: () => void;
  afterDelete: () => void;
  onClickTag: (id: number) => void;
}

export const LinkDetail = (props: Props) => {
  const [isEdit, setIsEdit] = useState();
  const [isDelete, setIsDelete] = useState();
  const handleDeleteLink = useCallback(() => {
    linksRepository.deleteLink(props.item.id).then(() => {
      props.afterDelete();
    });
  }, [props]);
  const handleDeleteCancel = useCallback(() => {
    setIsDelete(false);
  }, []);
  const onClickEdit = () => {
    setIsEdit(true);
  };
  const onClickDelete = () => {
    setIsDelete(true);
  };

  return (
    <Wrapper>
      <DefaultImage />
      <UnderWrapper>
        {isDelete && (
          <DeleteConfirmation
            onSubmit={handleDeleteLink}
            onCancel={handleDeleteCancel}
          />
        )}
        <LinkDetailBody
          item={props.item}
          isEdit={isEdit}
          isDelete={isDelete}
          onClickEdit={onClickEdit}
          onClickDelete={onClickDelete}
          onClickTag={props.onClickTag}
        />
      </UnderWrapper>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  color: ${props => props.theme.text};
`;

const UnderWrapper = styled.div`
  padding: 8px 0;
`;
