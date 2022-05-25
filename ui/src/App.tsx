import React from 'react';
import logo from './logo.svg';
import './App.css';
import {BrowserRouter, HashRouter, Link, Route, Routes} from "react-router-dom";
import {PasswdInfo} from "./pc/PasswdInfo";
import {PasswdInfo as MPasswdInfo} from "./mobile/PasswdInfo";
import RouterUtil from "./common/utils/RouterUtil";
import {Login} from "./pc/Login";

class App extends React.Component<any, any>{

    constructor(props: any) {
        super(props);
        this.state = {}
    }

    componentDidMount() {
        if (document.location.pathname === "/") {
            if ((navigator.userAgent.match(/(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i))) {
                //手机
                RouterUtil.push("/m/passwdInfo")
            } else {
                //电脑
                RouterUtil.push("/login")
            }
        }
    }

    render() {
        return (
            <BrowserRouter>
                <Routes>
                    {/*  <Route component={} path={"/login"}/>*/}
                    <Route path={"/passwdInfo"} element={<PasswdInfo/>}></Route>
                    <Route path={"/login"} element={<Login/>}></Route>
                    <Route path={"/m/passwdInfo"} element={<MPasswdInfo/>}></Route>
                </Routes>
            </BrowserRouter>
        );
    }

}

export default App;
