import React, { useCallback, useState } from "react";
import styled, { css } from "styled-components";
import colors from "../colors";
import { repositoryFactory } from "../api/repositoryFactory";
import { CreateLinkPayload } from "../api/linksRepository";

const linkRepository = repositoryFactory.get("links");

interface Props {
  afterPost: () => void;
}

export const AddLinkForm = (props: Props) => {
  const [formValue, setFormValue] = useState<CreateLinkPayload>({
    url: "",
    description: ""
  });

  const [formError, setFormError] = useState("");

  const onChangeValue = useCallback(
    (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
      e.persist();
      setFormValue(prev => ({ ...prev, [e.target.name]: e.target.value }));
    },
    []
  );

  const onSubmit = useCallback(() => {
    if (!formValue.url) {
      return setFormError("url must be set");
    }
    linkRepository
      .createLink({
        url: formValue.url,
        description: formValue.description
      })
      .then(res => {
        setFormValue({
          url: "",
          description: ""
        });
        setFormError("");
        props.afterPost();
      })
      .catch(err => {
        setFormError("unexpected error");
      });
  }, [formValue, props]);

  return (
    <Wrapper>
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
        <SubmitButton onClick={onSubmit}>Save</SubmitButton>
      </Field>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  padding: 16px;
`;

const Field = styled.div``;

const FieldItemBase = css`
  font-size: 16px;
  width: 100%;
  padding: 8px;
  box-sizing: border-box;
  border-radius: 4px;
  margin-bottom: 8px;
`;

const Input = styled.input`
  ${FieldItemBase}
  height: 32px;
  background: ${colors.mainWhite};
`;

const Textarea = styled.textarea`
  ${FieldItemBase}
  background: ${colors.mainWhite};
  resize: none;
  height: 128px
`;

const SubmitButton = styled.button`
  ${FieldItemBase}
  background: ${colors.primary};
  color: ${colors.mainWhite};
  font-weight: bold;
`;

const ErrorMessage = styled.div`
  padding: 16px 0;
  color: ${colors.danger};
`;
