import { createContext, Dispatch, SetStateAction } from "react";
import { ThemeType, themeNames } from "../theme";

type ViewContextProps = {
  themeName: ThemeType;
  setThemeName: Dispatch<SetStateAction<ThemeType>>;

  slideOpen: boolean;
  setSlideOpen: Dispatch<SetStateAction<boolean>>;
  slideTargetLinkId: number;
  setSlideTargetLinkId: Dispatch<SetStateAction<number>>;
};

const defaultProps = {
  themeName: themeNames[0],
  setThemeName: () => {},

  slideOpen: false,
  setSlideOpen: () => {},

  slideTargetLinkId: 0,
  setSlideTargetLinkId: () => {}
};

export const ViewContext = createContext<ViewContextProps>(defaultProps);
