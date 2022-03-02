import React from "react";
import ReactDOM from "react-dom";
import { ChakraProvider, extendTheme } from "@chakra-ui/react";
import { BrowserRouter } from "react-router-dom";
import Router from "./router";
import { RecoilRoot } from "recoil";
import NavBar from "./components/NavBar";

ReactDOM.render(
  <React.StrictMode>
    <RecoilRoot>
      <ChakraProvider>
        <BrowserRouter>
          <NavBar />
          <Router />
        </BrowserRouter>
      </ChakraProvider>
    </RecoilRoot>
  </React.StrictMode>,
  document.getElementById("root")
);
