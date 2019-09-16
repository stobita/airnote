import React, { useState, useCallback, useContext, useEffect } from "react";
import styled from "styled-components";
import linksRepository, { LinkPayload } from "../api/linksRepository";
import colors from "../colors";
import { Button } from "./Button";
import { ButtonPair } from "./ButtonPair";
import { TagInput } from "./TagInput";
import { Input } from "./Input";
import { Textarea } from "./Textarea";
import { DataContext } from "../context/dataContext";
import { ViewContext } from "../context/viewContext";
import { DeleteConfirmation } from "./DeleteConfirmation";

export const LinkForm = () => {
  const { links, tags, setLinks } = useContext(DataContext);
  const {
    slideTargetLinkId,
    setSlideTargetLinkId,
    slideOpen,
    setSlideOpen
  } = useContext(ViewContext);

  useEffect(() => {
    if (slideOpen === false) {
      setSlideTargetLinkId(0);
    }
  }, [slideOpen]);

  const target = links.find(v => v.id === slideTargetLinkId);
  const isEdit = target;
  const initFormValue: LinkPayload = {
    url: target ? target.url : "",
    description: target ? target.description : "",
    tags: target ? target.tags.map(v => v.text) : []
  };

  const [isDelete, setIsDelete] = useState(false);
  const [formValue, setFormValue] = useState<LinkPayload>(
    initFormValue || {
      url: "",
      description: "",
      tags: []
    }
  );

  const handleAfterSubmit = async (id: number) => {
    const links = await linksRepository.getAllLinks();
    setLinks(links);
    const original = await linksRepository.getLinkOriginal(id);
    const newlinks = links.map(i =>
      i.id === id ? { ...i, title: original.title } : i
    );
    setLinks(newlinks);
  };

  const [formError, setFormError] = useState("");

  const onChangeValue = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    e.persist();
    setFormValue(prev => ({ ...prev, [e.target.name]: e.target.value }));
  };

  const onChangeTag = useCallback((items: string[]) => {
    setFormValue(prev => ({ ...prev, tags: items }));
  }, []);

  const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (!formValue.url) {
      return setFormError("url must be set");
    }
    const action = isEdit
      ? linksRepository.updateLink(slideTargetLinkId, formValue)
      : linksRepository.createLink(formValue);
    const res = await action.catch(err => {
      setFormError("unexpected error");
      return Promise.reject(err);
    });

    handleAfterSubmit(res.id);
    setFormValue({
      url: "",
      description: "",
      tags: []
    });
  };

  const onCancel = () => {
    setSlideTargetLinkId(0);
    setSlideOpen(false);
  };

  const handleOnDeleteLink = async () => {
    await linksRepository.deleteLink(slideTargetLinkId);
    const links = await linksRepository.getAllLinks();
    setLinks(links);
    setSlideTargetLinkId(0);
    setSlideOpen(false);
  };

  const handleOnClickDelete = () => {
    setIsDelete(true);
  };

  return (
    <>
      {formError && <ErrorMessage>{formError}</ErrorMessage>}
      <form onSubmit={onSubmit}>
        <Field>
          <Input
            name="url"
            type="text"
            placeholder="URL"
            value={formValue.url}
            onChange={onChangeValue}
          />
        </Field>
        <Field>
          <Textarea
            name="description"
            placeholder="Description"
            value={formValue.description}
            onChange={onChangeValue}
          />
        </Field>
        <Field>
          <TagInput
            name="tag"
            placeholder="Tag"
            value={formValue.tags ? formValue.tags : []}
            onChange={onChangeTag}
            tags={tags}
          />
        </Field>
        <Field>
          <ButtonPair
            left={<Button primary>{isEdit ? "Save" : "Add"}</Button>}
            right={
              <Button type="button" onClick={onCancel}>
                Cancel
              </Button>
            }
          />
        </Field>
        <Field>
          {isEdit && (
            <Button danger type="button" onClick={handleOnClickDelete}>
              Delete
            </Button>
          )}
        </Field>
        {isDelete && (
          <DeleteConfirmation
            onSubmit={handleOnDeleteLink}
            onCancel={onCancel}
          />
        )}
      </form>
    </>
  );
};

const Field = styled.div`
  display: flex;
  margin-bottom: 8px;
`;

const ErrorMessage = styled.div`
  padding: 16px 0;
  color: ${colors.danger};
`;
