import React, { useState, useCallback } from "react";
import styled from "styled-components";
import { LinkPayload } from "../api/linksRepository";
import colors from "../colors";
import { Button } from "./Button";
import { ButtonPair } from "./ButtonPair";
import { Link } from "../model/link";
import { TagInput } from "./TagInput";
import { Input } from "./Input";
import { Textarea } from "./Textarea";

interface Props {
  initFormValue?: Link;
  onSubmit: (p: LinkPayload) => Promise<void>;
  onCancel?: () => void;
  afterSubmit: () => void;
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
      .then(() => {
        props.afterSubmit();
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
        />
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
