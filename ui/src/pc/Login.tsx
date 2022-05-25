import React from "react";
import css from "./style/login.module.css"
import {Button, Col, Input, Modal, Row} from "antd";
import BlankRow from "./component/BlankRow";
import {Notification} from "../common/component/Notification";
import StringUtil from "../common/utils/StringUtil";
import ObjectUtil from "../common/utils/ObjectUtil";
import HttpClient from "../common/component/HttpClient";
import HttpURL from "../common/env/HttpURL";

const STATUS = {
    "LOGIN": 0,
    "REGISTER": 1
}

class RegForm {
    username = ""
    password = ""
    rePassword = ""
}

class LoginForm {
    username = ""
    password = ""
}

class State {
    regVisible = false
    regForm = new RegForm()
    loginForm = new LoginForm()
}

class Props {

}

export class Login extends React.Component<Props, State> {

    constructor(props: Props) {
        super(props);
        this.state = new State()
    }

    private ui_register = {
        doOpen: () => {
            this.setState({
                regVisible: true
            })
        },
        doClear: () => {
            this.setState({
                regForm: {
                    username: "",
                    password: "",
                    rePassword: ""
                }
            })
        },
        doClose: () => {
            this.setState({
                regVisible: false
            })
            this.ui_register.doClear()
        },
        doValidate: (onBlank: () => void, onPasswdNotEqual: () => void) => {
            const validateFields = ["username", "password", "rePassword"]
            let key: keyof RegForm
            for (key in this.state.regForm) {
                if (validateFields.indexOf(key) > -1 && StringUtil.isBlank(this.state.regForm[key])) {
                    onBlank()
                    return false
                }
            }
            if (this.state.regForm.password !== this.state.regForm.rePassword) {
                onPasswdNotEqual()
                return false
            }
            return true
        },
        doOk: () => {
            const pass = this.ui_register.doValidate(() => {
                Notification.error("请填写完整信息！")
            }, () => {
                Notification.error("两次密码不一致！")
            })
            if (pass) {
                HttpClient.post(HttpURL.POST_REGISTER, {
                    username: this.state.regForm.username,
                    password: this.state.regForm.password
                }, resp => {
                    Notification.success("注册成功！请重新登录")
                    this.ui_register.doClose()
                }, err => {
                    Notification.error("注册失败！")
                })
            }
        },
        render: () => {
            return <div>
                <Modal visible={this.state.regVisible} onCancel={this.ui_register.doClose} onOk={this.ui_register.doOk}>
                    <div className={css.regModal}>
                        <Row>
                            <Col span={4}>账号</Col>
                            <Col span={20}>
                                <Input value={this.state.regForm.username} onChange={(val) => {
                                    this.setState({regForm: ObjectUtil.getNewProperty(this.state, "regForm.username", val.target.value)})
                                }}></Input>
                            </Col>
                        </Row>
                        <BlankRow/>
                        <Row>
                            <Col span={4}>密码</Col>
                            <Col span={20}>
                                <Input.Password value={this.state.regForm.password} onChange={(val) => {
                                    this.setState({regForm: ObjectUtil.getNewProperty(this.state, "regForm.password", val.target.value)})
                                }}></Input.Password>
                            </Col>
                        </Row>
                        <BlankRow/>
                        <Row>
                            <Col span={4}>确认密码</Col>
                            <Col span={20}>
                                <Input.Password value={this.state.regForm.rePassword} onChange={(val) => {
                                    this.setState({regForm: ObjectUtil.getNewProperty(this.state, "regForm.rePassword", val.target.value)})
                                }}></Input.Password>
                            </Col>
                        </Row>
                    </div>
                </Modal>
            </div>
        }
    }

    private ui_body = {
        doLogin: () => {
            HttpClient.post(HttpURL.POST_LOGIN, {
                username: this.state.loginForm.username,
                password: this.state.loginForm.password
            }, resp => {
                // todo
            })
        },
        render: () => {
            return <div>
                <Row>
                    <Col span={4}>账号</Col>
                    <Col span={20}>
                        <Input></Input>
                    </Col>
                </Row>
                <BlankRow/>
                <Row>
                    <Col span={4}>密码</Col>
                    <Col span={20}>
                        <Input.Password></Input.Password>
                    </Col>
                </Row>
                <BlankRow/>
                <Row>
                    <Col span={4}></Col>
                    <Col span={10} >
                        <Button onClick={this.ui_register.doOpen}>注册</Button>
                    </Col>
                    <Col span={10} style={{textAlign: "right"}}>
                        <Button type={"primary"}>登录</Button>
                    </Col>
                </Row>
            </div>
        }
    }

    render() {
        return <div className={css.loginBody}>
            {this.ui_body.render()}
            {this.ui_register.render()}
        </div>;
    }
}