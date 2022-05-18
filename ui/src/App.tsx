import React from 'react';
import logo from './logo.svg';
import './App.css';
import {BrowserRouter, Route, Routes} from "react-router-dom";
import {PasswdInfo} from "./pc/PasswdInfo";

function App() {
    return (
        <BrowserRouter>
            <Routes>
                {/*  <Route component={} path={"/login"}/>*/}
                <Route path={"/passwdInfo"} element={<PasswdInfo/>}></Route>
            </Routes>
        </BrowserRouter>

    );
}

export default App;
