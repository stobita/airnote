import React, { ReactNode, useState, useEffect } from "react";
import { ViewContext } from "../context/viewContext";
import { ThemeProvider } from "styled-components";
import theme, { ThemeType } from "../theme";
import { localStorageRepository } from "../localStorageService";

interface Props {
  children: ReactNode;
}

export const ViewContextProvider = (props: Props) => {
  const currentThemeName = localStorageRepository.getThemeName();
  const [themeName, setThemeName] = useState<ThemeType>(currentThemeName);

  const [slideOpen, setSlideOpen] = useState(false);
  const [slideTargetLinkId, setSlideTargetLinkId] = useState(0);

  useEffect(() => {
    localStorageRepository.setThemeName(themeName);
  }, [themeName]);

  useEffect(() => {
    if (slideTargetLinkId !== 0) {
      setSlideOpen(true);
    }
  }, [slideTargetLinkId]);

  return (
    <ThemeProvider theme={theme[themeName]}>
      <ViewContext.Provider
        value={{
          themeName,
          setThemeName,
          slideOpen,
          setSlideOpen,
          slideTargetLinkId,
          setSlideTargetLinkId
        }}
      >
        {props.children}
      </ViewContext.Provider>
    </ThemeProvider>
  );
};
