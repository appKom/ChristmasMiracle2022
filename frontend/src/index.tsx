import React from "react";
import ReactDOM from "react-dom";
import { ChakraProvider, extendTheme } from "@chakra-ui/react";
import { BrowserRouter } from "react-router-dom";
import Router from "./router";
import { RecoilRoot } from "recoil";
import NavBar from "./components/NavBar";
import Footer from "./components/Footer";
import { Box } from "@chakra-ui/react";
const theme = extendTheme({
  styles: {
    global: {
      "html, body": {
        height: "100%",
        color: "#e9e5e1",
        webkitfontsmoothing: "antialiased",
        mozOsxFontSmoothing: "grayscale",
        background: "#303a43",
      },
    },
  },
});

ReactDOM.render(
  <React.StrictMode>
    <RecoilRoot>
      <ChakraProvider theme={theme}>
        <BrowserRouter>
          <Box minHeight="90vh">
            <NavBar />
            <Router />
          </Box>
          <Footer />
        </BrowserRouter>
      </ChakraProvider>
    </RecoilRoot>
  </React.StrictMode>,
  document.getElementById("root")
);
