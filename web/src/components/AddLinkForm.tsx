import React, { useCallback } from "react";
import { LinkForm } from "./LinkForm";
import linksRepository, { LinkPayload } from "../api/linksRepository";
import { Tag } from "../model/link";

interface Props {
  tags: Tag[];
  afterSubmit: (id: number) => void;
}

export const AddLinkForm = (props: Props) => {
  const onSubmit = useCallback((input: LinkPayload) => {
    return linksRepository
      .createLink({
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
  }, []);

  return (
    <LinkForm
      onSubmit={onSubmit}
      afterSubmit={props.afterSubmit}
      tags={props.tags}
    />
  );
};
