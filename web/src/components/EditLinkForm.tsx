import React, { useCallback } from "react";
import { repositoryFactory } from "../api/repositoryFactory";
import { LinkForm } from "./LinkForm";
import { LinkPayload } from "../api/linksRepository";
import { Link } from "../model/link";

const linkRepository = repositoryFactory.get("links");

interface Props {
  afterSubmit: () => void;
  onCancel: () => void;
  target: Link;
}

export const EditLinkForm = (props: Props) => {
  const onSubmit = useCallback(
    (input: LinkPayload) => {
      return linkRepository
        .updateLink(props.target.id, {
          url: input.url,
          description: input.description,
          tags: input.tags
        })
        .then(() => {
          return;
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
    />
  );
};
