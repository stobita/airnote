import React, { useCallback, useState } from "react";
import { Link } from "../model/link";
import { ReactComponent as DefaultImage } from "../assets/default.svg";
import { ReactComponent as EditIconImage } from "../assets/edit.svg";
import { ReactComponent as DeleteIconImage } from "../assets/delete.svg";
import styled, { css } from "styled-components";
import colors from "../colors";
import { repositoryFactory } from "../api/repositoryFactory";
import { EditLinkForm } from "./EditLinkForm";
import { DeleteConfirmation } from "./DeleteConfirmation";
import { TagBadge } from "./TagBadge";

const linkRepository = repositoryFactory.get("links");

interface Props {
  item: Link;
  afterUpdate: () => void;
  afterDelete: () => void;
}

export const LinkDetail = (props: Props) => {
  const [isEdit, setIsEdit] = useState();
  const [isDelete, setIsDelete] = useState();
  const handleDeleteLink = useCallback(() => {
    linkRepository.deleteLink(props.item.id).then(() => {
      props.afterDelete();
    });
  }, [props]);
  const onClickEdit = useCallback(() => {
    setIsEdit(true);
  }, []);
  const onClickDelete = useCallback(() => {
    setIsDelete(true);
  }, []);

  const handleEditCancel = useCallback(() => {
    setIsEdit(false);
  }, []);
  const handleDeleteCancel = useCallback(() => {
    setIsDelete(false);
  }, []);

  const Container = () => {
    return isEdit ? (
      <EditLinkForm
        target={props.item}
        afterSubmit={props.afterUpdate}
        onCancel={handleEditCancel}
      />
    ) : (
      <TextPreview />
    );
  };

  const TextPreview = () => {
    return (
      <>
        <Row>
          <Anchor
            href={props.item.url}
            target="_blank"
            rel="noopener noreferrer"
          >
            {props.item.url}
          </Anchor>
          {!isDelete && !isEdit && (
            <Operator>
              <EditIcon onClick={onClickEdit} />
              <DeleteIcon onClick={onClickDelete} />
            </Operator>
          )}
        </Row>
        <p>{props.item.description}</p>
        <div>
          {props.item.tags.map(v => (
            <TagBadge key={v.id}>{v.text}</TagBadge>
          ))}
        </div>
      </>
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
  color: ${colors.mainWhite};
`;

const UnderWrapper = styled.div`
  padding: 8px 0;
`;

const Row = styled.div`
  display: flex;
  align-items: center;
  padding: 16px 0;
`;

const IconBase = css`
  fill: ${colors.mainWhite};
  height: 32px;
  padding: 8px;
`;

const Operator = styled.div`
  flex: 2;
  display: flex;
  justify-content: flex-end;
`;

const EditIcon = styled(EditIconImage)`
  ${IconBase}
`;

const DeleteIcon = styled(DeleteIconImage)`
  ${IconBase}
`;

const Anchor = styled.a`
  color: ${colors.mainWhite};
  flex: 5;
`;
