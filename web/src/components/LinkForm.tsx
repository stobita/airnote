import React, { useState, useCallback } from "react";
import styled from "styled-components";
import { LinkPayload } from "../api/linksRepository";
import colors from "../colors";
import { Button } from "./Button";
import { ButtonPair } from "./ButtonPair";
import { TagInput } from "./TagInput";
import { Input } from "./Input";
import { Textarea } from "./Textarea";
import { Tag } from "../model/link";

interface Props {
  initFormValue?: LinkPayload;
  tags: Tag[];
  onSubmit: (p: LinkPayload) => Promise<number>;
  onCancel?: () => void;
  afterSubmit: (id: number) => void;
}

export const LinkForm = (props: Props) => {
  const [formValue, setFormValue] = useState<LinkPayload>(
    props.initFormValue || {
      url: "",
      description: "",
      tags: []
    }
  );

  const [formError, setFormError] = useState("");

  const onChangeValue = useCallback(
    (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
      e.persist();
      setFormValue(prev => ({ ...prev, [e.target.name]: e.target.value }));
    },
    []
  );

  const onChangeTag = useCallback((items: string[]) => {
    setFormValue(prev => ({ ...prev, tags: items }));
  }, []);

  const onSubmit = useCallback(() => {
    if (!formValue.url) {
      return setFormError("url must be set");
    }
    props
      .onSubmit(formValue)
      .then(id => {
        props.afterSubmit(id);
      })
      .catch(e => {
        setFormError("unexpected error");
      });
    setFormValue({
      url: "",
      description: "",
      tags: []
    });
  }, [formValue, props]);

  const onClickRecommendedTag = (e: React.MouseEvent<HTMLElement>) => {
    const selectedText = e.currentTarget.dataset.text;
    if (selectedText && !formValue.tags.some(v => v === selectedText)) {
      const tags = [...formValue.tags, selectedText];
      setFormValue(prev => ({ ...prev, tags: tags }));
    }
  };

  return (
    <>
      {formError && <ErrorMessage>{formError}</ErrorMessage>}
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
          tags={props.tags}
        />
      </Field>
      <Field>
        <FieldTitle>Recommended:</FieldTitle>
        {props.tags.map(v => (
          <UsedTag
            key={v.id}
            data-text={v.text}
            onClick={onClickRecommendedTag}
          >
            #{v.text}
          </UsedTag>
        ))}
      </Field>
      <Field>
        {props.onCancel ? (
          <ButtonPair
            left={
              <Button primary onClick={onSubmit}>
                Save
              </Button>
            }
            right={<Button onClick={props.onCancel}>Cancel</Button>}
          />
        ) : (
          <Button primary onClick={onSubmit}>
            Save
          </Button>
        )}
      </Field>
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

const FieldTitle = styled.span`
  color: ${colors.mainWhite};
  margin-right: 8px;
  font-weight: bold;
`;

const UsedTag = styled.span`
  color: ${colors.mainWhite};
  margin-right: 8px;
  cursor: pointer;
`;
