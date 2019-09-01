import React, { useCallback, useState } from "react";
import { Link, Tag } from "../model/link";
import { ReactComponent as DefaultImage } from "../assets/default.svg";
import styled, { css } from "styled-components";
import { EditLinkForm } from "./EditLinkForm";
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
  const handleEditCancel = useCallback(() => {
    setIsEdit(false);
  }, []);
  const handleDeleteCancel = useCallback(() => {
    setIsDelete(false);
  }, []);
  const onClickEdit = () => {
    setIsEdit(true);
  };
  const onClickDelete = () => {
    setIsDelete(true);
  };

  const Container = () => {
    return isEdit ? (
      <EditLinkForm
        target={props.item}
        afterSubmit={props.afterUpdate}
        onCancel={handleEditCancel}
        tags={props.tags}
      />
    ) : (
      <LinkDetailBody
        item={props.item}
        isEdit={isEdit}
        isDelete={isDelete}
        onClickEdit={onClickEdit}
        onClickDelete={onClickDelete}
        onClickTag={props.onClickTag}
      />
    );
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
        <Container />
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
