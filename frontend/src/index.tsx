import * as React from "react";
import * as ReactDOM from "react-dom";

import 'semantic-ui-css/semantic.min.css'

import { Container } from 'semantic-ui-react'
import { MyMenu } from "./components/MyMenu";
import { Route, Switch } from "react-router";
import { BrowserRouter } from "react-router-dom";
import { ViewItems } from "./components/ViewItems";
import { App } from "./components/App";


ReactDOM.render(
    <BrowserRouter>
        <App />
    </BrowserRouter >
    ,
    document.getElementById("root")
);