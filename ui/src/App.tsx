import React from 'react';
import logo from './logo.svg';
import './App.css';
import {HashRouter as Router, Route, Switch} from "react-router-dom";
import {PasswdInfo} from "./pc/PasswdInfo";

function App() {
  return (
      <Router>
        <Switch>
        {/*  <Route component={} path={"/login"}/>*/}
          <Route component={PasswdInfo} path={"/"}/>
        </Switch>
      </Router>
  );
}

export default App;
