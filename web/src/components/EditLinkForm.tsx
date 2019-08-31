import React, { useCallback } from "react";
import { LinkForm } from "./LinkForm";
import linksRepository, { LinkPayload } from "../api/linksRepository";
import { Link, Tag } from "../model/link";

interface Props {
  afterSubmit: () => void;
  onCancel: () => void;
  target: Link;
  tags: Tag[];
}

export const EditLinkForm = (props: Props) => {
  const onSubmit = useCallback(
    (input: LinkPayload) => {
      return linksRepository
        .updateLink(props.target.id, {
          url: input.url,
          description: input.description,
          tags: input.tags
        })
        .then(res => {
          return res.id;
        })
        .catch(e => {
          return Promise.reject(e);
        });
    },
    [props.target.id]
  );

  const formInitValue: LinkPayload = {
    url: props.target.url,
    description: props.target.description,
    tags: props.target.tags.map(v => v.text)
  };

  return (
    <LinkForm
      initFormValue={formInitValue}
      onSubmit={onSubmit}
      afterSubmit={props.afterSubmit}
      onCancel={props.onCancel}
      tags={props.tags}
    />
  );
};
