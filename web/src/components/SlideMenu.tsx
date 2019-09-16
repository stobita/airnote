import React, { ReactNode, useContext } from "react";
import styled from "styled-components";
import { Transition } from "react-transition-group";
import { TransitionStatus } from "react-transition-group/Transition";
import { CloseButton } from "./CloseButton";
import { ViewContext } from "../context/viewContext";
import colors from "../colors";

interface Props {
  children: ReactNode;
}

export const SlideMenu = (props: Props) => {
  const { slideOpen, setSlideOpen } = useContext(ViewContext);
  const handleOnClickClose = () => {
    setSlideOpen(false);
  };
  return (
    <>
      <Transition
        in={slideOpen}
        timeout={{ appear: 0, exit: 500 }}
        unmountOnExit
      >
        {state => (
          <Inner state={state}>
            <Head>
              <CloseButton onClick={handleOnClickClose} />
            </Head>
            {props.children}
          </Inner>
        )}
      </Transition>
      {slideOpen && <SlideMask onClick={handleOnClickClose}></SlideMask>}
    </>
  );
};

const SlideMask = styled.div`
  width: 100%;
  height: 100%;
  background: ${colors.thinGray};
  opacity: 0.5;
  position: absolute;
`;
const Head = styled.div`
  display: flex;
  justify-content: flex-end;
  margin-bottom: 8px;
`;

const slideWidth = 500;

const Inner = styled.div<{ state: TransitionStatus }>`
  z-index: 1;
  background: ${props => props.theme.bg};
  color: ${props => props.theme.text};
  border-left: 1px solid ${props => props.theme.border};
  padding: 8px 24px;
  box-sizing: border-box;
  position: fixed;
  overflow-y: auto;
  right: -${slideWidth}px;
  transition: 0.5s;
  width: ${slideWidth}px;
  transform: translateX(
    ${({ state }) => (state === "entered" ? -slideWidth : 0)}px
  );
  min-height: 100vh;
  height: 100%;
`;
