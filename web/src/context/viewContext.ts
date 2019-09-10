import { createContext, Dispatch, SetStateAction } from "react";
import { ThemeType, themeNames } from "../theme";

type ViewContextProps = {
  themeName: ThemeType;
  setThemeName: Dispatch<SetStateAction<ThemeType>>;
};

const defaultProps = {
  themeName: themeNames[0],
  setThemeName: () => {}
};

export const ViewContext = createContext<ViewContextProps>(defaultProps);
