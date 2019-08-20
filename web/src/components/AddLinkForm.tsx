import React, { useCallback } from "react";
import { repositoryFactory } from "../api/repositoryFactory";
import { LinkForm } from "./LinkForm";
import { LinkPayload } from "../api/linksRepository";

const linkRepository = repositoryFactory.get("links");

interface Props {
  afterSubmit: () => void;
}

export const AddLinkForm = (props: Props) => {
  const onSubmit = useCallback((input: LinkPayload) => {
    return linkRepository
      .createLink({
        url: input.url,
        description: input.description
      })
      .then(() => {
        return;
      })
      .catch(e => {
        return Promise.reject(e);
      });
  }, []);

  return <LinkForm onSubmit={onSubmit} afterSubmit={props.afterSubmit} />;
};
