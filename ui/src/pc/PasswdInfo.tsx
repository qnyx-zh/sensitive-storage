import React from 'react'
import {Button, Form, Input, Modal, Row, Table} from "antd";
import BlankElement from "./component/BlankElement";
import BlankRow from "./component/BlankRow";
import HttpClient from "../common/component/HttpClient";
import Search from "antd/es/input/Search";
import {CopyOutlined, LockOutlined, UnlockOutlined} from '@ant-design/icons';
import css from './style/passwd-info.module.css';
import ObjectUtil from "../common/utils/ObjectUtil";
import Resp from "../common/model/HttpModel";

class State {
    tableData: Array<{ key: string, label: string }> = []
    dialogVisible: boolean = false
    formData = {
        username: "",
        password: "",
        topic: "",
        description: ""
    }
    unlockedPasswd: { [key: string]: boolean } = {}
}

class Props {
}

export class PasswdInfo extends React.Component<Props, State> {

    private ui_table = {
        columns: [
            {
                title: '标题',
                dataIndex: 'topic',
                key: 'topic',
                render: (col: any, row: any, index: any) => {
                    return <div>
                        <label>{row.topic}</label>
                        <span className={css.passwdCopy} title={"复制密码"}>
                             <CopyOutlined/>
                        </span>
                    </div>
                }
            },
            {
                title: '用户名',
                dataIndex: 'username',
                key: 'username',
            },
            {
                title: '密码',
                dataIndex: 'password',
                key: 'password',
                render: (col: any, row: any, index: any) => {
                    return <div>
                        <label>{this.state.unlockedPasswd[row.id] ? row['password'] : '*******'}</label>
                        <span className={css.passwdLock} onClick={() => {
                            this.setState({
                                    unlockedPasswd: {
                                        ...this.state.unlockedPasswd,
                                        ["" + row.id]: !this.state.unlockedPasswd[row.id]
                                    }
                                }
                            )
                        }
                        }>
                            {this.state.unlockedPasswd[row.id] ? <UnlockOutlined/> : <LockOutlined/>}
                        </span>
                    </div>
                }
            },
            {
                title: '描述',
                dataIndex: 'description',
                key: 'description',
            },
            {
                title: '操作',
                dataIndex: 'ope',
                key: 'ope',
                render: (col: any, row: any, index: any) => {
                    return <div>
                        <Button onClick={() => {
                            this.ui_dialogs.openEditForm(row.id)
                        }}>编辑</Button>
                        <BlankElement/>
                        <Button onClick={() => {
                            this.ui_table.doDelete(row.id)
                        }}>删除</Button>
                    </div>;
                }
            }
        ],
        doDelete: (id: string) => {
            // todo
        },
        loadData: () => {
            const url = "http://127.0.0.1:4523/mock/991824/passwdInfos"
            HttpClient.get(url, resp => {
                this.setState({
                    tableData: resp.data.passwds
                })
            })
        },
        doSearch: (text: string) => {
            const url = "http://127.0.0.1:4523/mock/991824/search?q=" + text
            HttpClient.get(url, resp => {
                this.setState({
                    tableData: resp.data.passwds
                })
            })
        }
    }

    protected ui_dialogs = {
        layout: {
            labelCol: {span: 5},
            wrapperCol: {span: 19},
        },
        doClose: () => {
            this.setState({
                dialogVisible: false
            })
        },
        doOk: () => {
            const url = ""
            HttpClient.post(url, this.state.formData, (resp: Resp) => {
                this.ui_dialogs.doClose();
                this.ui_table.loadData();
            });
        },
        doClear: () => {
            this.setState({
                formData: {
                    username: "",
                    password: "",
                    topic: "",
                    description: ""
                }
            })
        },
        openEditForm: (id?: string) => {
            this.ui_dialogs.doClear()
            const showFn = () => {
                this.setState({
                        dialogVisible: true
                    }
                );
            }
            if (id) {
                const url = ""
                HttpClient.get(url, resp => {
                    this.ui_table.loadData()
                    this.ui_dialogs.doClose()
                })
            } else {
                showFn();
            }
        },
        render: () => {
            return (
                <Modal title={"编辑密码信息"}
                       onCancel={this.ui_dialogs.doClose}
                       onOk={this.ui_dialogs.doOk}
                       visible={this.state.dialogVisible}>
                    <Form {...this.ui_dialogs.layout}>
                        <Form.Item label={"标题"}>
                            <Input value={this.state.formData.topic} onChange={(val) => {
                                this.setState({formData: ObjectUtil.getNewProperty(this.state, "formData.topic", val.target.value)})
                            }}/>
                        </Form.Item>
                        <Form.Item label={"用户名"}>
                            <Input value={this.state.formData.username} onChange={(val) => {
                                this.setState({formData: ObjectUtil.getNewProperty(this.state, "formData.username", val.target.value)})
                            }}/>
                        </Form.Item>
                        <Form.Item label={"密码"}>
                            <Input value={this.state.formData.password} onChange={(val) => {
                                this.setState({formData: ObjectUtil.getNewProperty(this.state, "formData.password", val.target.value)})
                            }}/>
                        </Form.Item>
                        <Form.Item label={"描述"}>
                            <Input.TextArea value={this.state.formData.description} onChange={(val) => {
                                this.setState({formData: ObjectUtil.getNewProperty(this.state, "formData.description", val.target.value)})
                            }}/>
                        </Form.Item>
                    </Form>
                </Modal>
            )
        }
    }

    constructor(props: Props) {
        super(props);
        this.state = new State()
    }


    componentDidMount() {
        this.ui_table.loadData()
    }

    render() {
        return (<div>
            <Row>
                <Button type={"primary"} onClick={() => {
                    this.ui_dialogs.openEditForm()
                }}>添加</Button>
                <BlankElement/>
                <Search placeholder="请输入标题" onSearch={this.ui_table.doSearch} style={{width: 200}}/>
                <BlankElement/>
            </Row>
            <BlankRow/>
            <Table
                rowKey={"id"}
                className="components-table-demo-nested"
                columns={this.ui_table.columns}
                dataSource={this.state.tableData}/>
            {this.ui_dialogs.render()}
        </div>);
    }

}